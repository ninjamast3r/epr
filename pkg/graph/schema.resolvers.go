package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.29

import (
	"context"
	"fmt"

	"gitlab.sas.com/async-event-infrastructure/server/pkg/models"
)

// CreateEvent is the resolver for the create_event field.
func (r *mutationResolver) CreateEvent(ctx context.Context, input models.EventInput) (*models.Event, error) {
	panic(fmt.Errorf("not implemented: CreateEvent - create_event"))
}

// UpdateEvent is the resolver for the update_event field.
func (r *mutationResolver) UpdateEvent(ctx context.Context, id string, input models.EventInput) (*models.Event, error) {
	panic(fmt.Errorf("not implemented: UpdateEvent - update_event"))
}

// DeleteEvent is the resolver for the delete_event field.
func (r *mutationResolver) DeleteEvent(ctx context.Context, id string) (*bool, error) {
	panic(fmt.Errorf("not implemented: DeleteEvent - delete_event"))
}

// CreateEventReceiver is the resolver for the create_event_receiver field.
func (r *mutationResolver) CreateEventReceiver(ctx context.Context, input models.EventReceiverInput) (*models.EventReceiver, error) {
	panic(fmt.Errorf("not implemented: CreateEventReceiver - create_event_receiver"))
}

// UpdateEventReceiver is the resolver for the update_event_receiver field.
func (r *mutationResolver) UpdateEventReceiver(ctx context.Context, id string, input models.EventReceiverInput) (*models.EventReceiver, error) {
	panic(fmt.Errorf("not implemented: UpdateEventReceiver - update_event_receiver"))
}

// DeleteEventReceiver is the resolver for the delete_event_receiver field.
func (r *mutationResolver) DeleteEventReceiver(ctx context.Context, id string) (*bool, error) {
	panic(fmt.Errorf("not implemented: DeleteEventReceiver - delete_event_receiver"))
}

// CreateEventReceiverGroup is the resolver for the create_event_receiver_group field.
func (r *mutationResolver) CreateEventReceiverGroup(ctx context.Context, input models.EventReceiverGroupInput) (*models.EventReceiverGroup, error) {
	panic(fmt.Errorf("not implemented: CreateEventReceiverGroup - create_event_receiver_group"))
}

// UpdateEventReceiverGroup is the resolver for the update_event_receiver_group field.
func (r *mutationResolver) UpdateEventReceiverGroup(ctx context.Context, id string, input models.EventReceiverGroupInput) (*models.EventReceiverGroup, error) {
	panic(fmt.Errorf("not implemented: UpdateEventReceiverGroup - update_event_receiver_group"))
}

// DeleteEventReceiverGroup is the resolver for the delete_event_receiver_group field.
func (r *mutationResolver) DeleteEventReceiverGroup(ctx context.Context, id string) (*bool, error) {
	panic(fmt.Errorf("not implemented: DeleteEventReceiverGroup - delete_event_receiver_group"))
}

// AddEventReceiverToEventReceiverGroup is the resolver for the add_event_receiver_to_event_receiver_group field.
func (r *mutationResolver) AddEventReceiverToEventReceiverGroup(ctx context.Context, eventReceiverGroup string, eventReceiver string) (*bool, error) {
	panic(fmt.Errorf("not implemented: AddEventReceiverToEventReceiverGroup - add_event_receiver_to_event_receiver_group"))
}

// RemoveEventReceiverFromEventReceiverGroup is the resolver for the remove_event_receiver_from_event_receiver_group field.
func (r *mutationResolver) RemoveEventReceiverFromEventReceiverGroup(ctx context.Context, eventReceiverGroup string, eventReceiver string) (*bool, error) {
	panic(fmt.Errorf("not implemented: RemoveEventReceiverFromEventReceiverGroup - remove_event_receiver_from_event_receiver_group"))
}

// Event is the resolver for the event field.
func (r *queryResolver) Event(ctx context.Context, id string) (*models.Event, error) {
	panic(fmt.Errorf("not implemented: Event - event"))
}

// EventReceiver is the resolver for the event_receiver field.
func (r *queryResolver) EventReceiver(ctx context.Context, id string) (*models.EventReceiver, error) {
	panic(fmt.Errorf("not implemented: EventReceiver - event_receiver"))
}

// EventReceiverGroup is the resolver for the event_receiver_group field.
func (r *queryResolver) EventReceiverGroup(ctx context.Context, id string) (*models.EventReceiverGroup, error) {
	panic(fmt.Errorf("not implemented: EventReceiverGroup - event_receiver_group"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
