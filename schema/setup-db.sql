CREATE TABLE "Event" (
  "ID"  varchar  NOT NULL UNIQUE,
  "name"  varchar  NOT NULL,
  "version"  varchar  NOT NULL,
  "release"  varchar NOT NULL,
  "platformID"  varchar NOT NULL,
  "package"  varchar  NOT NULL,
  "description"  varchar NOT NULL,
  "payload"  varchar NOT NULL,
  "event_receiver_id"  varchar NOT NULL,
  "success"  boolean  NOT NULL,
  CONSTRAINT Event_pk PRIMARY KEY ("ID", "name", "version", "release", "package", "platformID")
);

CREATE TABLE "Event_receiver" (
  "ID"  varchar  NOT NULL  UNIQUE,
  "name"  varchar  NOT NULL,
  "type"  varchar  NOT NULL,
  "version"  varchar  NOT NULL  UNIQUE,
  "description"  varchar  NOT NULL,
  "enabled"  boolean  NOT NULL,
  CONSTRAINT Event_receiver_pk PRIMARY KEY ("ID", "name", "type", "version", "enabled")
);

CREATE TABLE "Event_receiver_group" (
  "ID"  varchar  NOT NULL  UNIQUE,
  "name"  varchar NOT NULL,
  "type" varchar NOT NULL,
  "version"  varchar NOT NULL,
  "enabled"  boolean  NOT NULL,
  CONSTRAINT Event_receiver_group_pk PRIMARY KEY ("ID", "type", "name", "version", "enabled")
);

CREATE TABLE "Event_receiver_group_to_event_receiver" (
  "event_receiver_group"  varchar,
  "event_receiver"  varchar,
  UNIQUE ("event_receiver_group", "event_receiver")
);


ALTER TABLE "Event" ADD CONSTRAINT "Event_fk0" FOREIGN KEY ("event_receiver_id") REFERENCES "Event_receiver"("ID");

ALTER TABLE "Event_receiver_group_to_event_receiver" ADD CONSTRAINT "Event_receiver_group_to_event_receiver_fk0" FOREIGN KEY ("event_receiver_group") REFERENCES "Event_receiver_group"("ID");

ALTER TABLE "Event_receiver_group_to_event_receiver" ADD CONSTRAINT "Event_receiver_group_to_event_receiver_fk1" FOREIGN KEY ("event_receiver") REFERENCES "Event_receiver"("ID");