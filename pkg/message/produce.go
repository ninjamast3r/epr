// Copyright © 2019, SAS Institute Inc., Cary, NC, USA.  All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package message

import (
	"fmt"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/sassoftware/event-provenance-registry/pkg/utils"
)

// Producer defines an inteface for producing events
type Producer interface {
	Async(string, interface{})
	Send(string, interface{}) error
	ConsumeSuccesses(*sync.WaitGroup)
	ConsumeErrors(*sync.WaitGroup)
	Close() error
}

type producer struct {
	sync  sarama.SyncProducer
	async sarama.AsyncProducer
}

// NewConfig creates a new sarama.Config object with TLS disabled.
func NewConfig(version string) (*sarama.Config, error) {
	kafkaver, err := sarama.ParseKafkaVersion(version)
	if err != nil {
		return nil, err
	}

	config := sarama.NewConfig()
	config.Version = kafkaver
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	config.Producer.Retry.Max = 10                   // Retry up to 10 times to produce the message
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewHashPartitioner
	config.Metadata.Timeout = 30 * time.Second
	config.ClientID = "server.service" + utils.NewULIDAsString()
	return config, nil
}

func newSASLConfig(user, password, version string) (*sarama.Config, error) {
	config, err := NewConfig(version)
	if err != nil {
		return nil, err
	}

	config.Net.TLS.Enable = true
	config.Net.SASL.Enable = true
	config.Net.SASL.User = user
	config.Net.SASL.Password = password
	return config, nil
}

// NewSCRAMConfig creates a new SASL SCRAM enabled sarama config for communicating over TLS
func NewSCRAMConfig(user, password, version string) (*sarama.Config, error) {
	config, err := newSASLConfig(user, password, version)
	if err != nil {
		return nil, err
	}
	config.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient {
		return &SCRAMClient{HashGeneratorFcn: SHA512}
	}
	config.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA512

	return config, nil
}

// NewPlainConfig creates a new SASL PLAINTEXT enabled sarama config for communicating over TLS
func NewPlainConfig(user, password, version string) (*sarama.Config, error) {
	config, err := newSASLConfig(user, password, version)
	if err != nil {
		return nil, err
	}
	config.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	return config, nil
}

// NewConfigEnv returns a new Sarama config with Kafka security options enabled or disabled depending on
// configured envirionment variables. If SASL_USERNAME and SASL_PASSWORD are set, a config with SASL
// and TLS enabled will be returned, else a config with the former disabled will be returned. If SASL is
// enabled SASL_MECHANISM must be set to SCRAM or PLAIN
// Supported SASL_MECHANISMS: SCRAM, PLAIN
// Future    SASL_MECHANISMS: OAUTH2
func NewConfigEnv(version string) (*sarama.Config, error) {
	saslAuth := GetSASLAuthentication()

	if !saslAuth.SASLEnabled() {
		logger.Info("no Kafka Authentication Enabled")
		return NewConfig(version)
	}

	// NONE case covered by the above .SASLEnabled() function
	switch saslAuth.Mechanism { //nolint:exhaustive
	case SCRAM:
		logger.Info("Kafka Authentication Mechanism: SASL SCRAM")
		return NewSCRAMConfig(saslAuth.Username, saslAuth.Password, version)
	case PLAIN:
		logger.Info("Kafka Authentication Mechanism: SASL PLAIN")
		return NewPlainConfig(saslAuth.Username, saslAuth.Password, version)
	case OAUTH2: // TODO: add support for OAUTH2
		logger.Info("Kafka Authentication Mechanism: SASL OAUTH2")
		return nil, fmt.Errorf("SASL_MECHANISM 'OAUTH2' not currently supported")
	default:
		return nil, fmt.Errorf("SASL_MECHANISM not supported")
	}
}

// NewProducer creates a producer instance
func NewProducer(brokers []string, config *sarama.Config) (Producer, error) {
	if len(brokers) == 0 {
		return &producer{
			sync:  nil,
			async: nil,
		}, nil
	}

	async, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	syn, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &producer{
		sync:  syn,
		async: async,
	}, nil
}

// Close closes down the Kafka producer
func (p *producer) Close() error {
	if p.sync == nil && p.async == nil {
		return nil
	}
	if err := p.sync.Close(); err != nil {
		return err
	}
	return p.async.Close()
}

// ConsumeSuccesses consumes and logs successful message sends.
func (p *producer) ConsumeSuccesses(wg *sync.WaitGroup) {
	if p.async == nil {
		logger.V(1).Info("kafka messaging disabled")
		return
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for suc := range p.async.Successes() {
			e, _ := suc.Value.Encode()
			logger.V(1).Info(fmt.Sprintf("message sent successfully: '%s'\n", string(e)))
		}
	}()
}

// ConsumeErrors consumes and logs messaging errors
func (p *producer) ConsumeErrors(wg *sync.WaitGroup) {
	if p.async == nil {
		logger.V(1).Info("kafka messaging disabled")
		return
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range p.async.Errors() {
			logger.V(1).Info(fmt.Sprintf("error sending message '%s'\n", err))
		}
	}()
}

// Async encodes an arbitrary struct and sends it on the given topic asynchronously.
func (p *producer) Async(topic string, value interface{}) {
	if p.async == nil {
		logger.V(1).Info("kafka messaging disabled")
		return
	}
	encodedMsg := &messageInfo{
		msgType: value,
		Topic:   topic,
	}

	p.async.Input() <- &sarama.ProducerMessage{
		Topic: topic,
		Value: encodedMsg,
	}
}

// Send encodes an arbitrary struct and sends it on the given topic synchronously.
func (p *producer) Send(topic string, value interface{}) error {
	if p.async == nil {
		logger.V(1).Info("kafka messaging disabled")
		return nil
	}
	encodedMsg := &messageInfo{
		msgType: value,
		Topic:   topic,
	}

	_, _, err := p.sync.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: encodedMsg,
	})
	if err != nil {
		return err
	}

	return nil
}
