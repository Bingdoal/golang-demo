BEGIN;

CREATE TABLE IF NOT EXISTS "user" (
    "id" bigserial NOT NULL PRIMARY KEY,
    "name" character varying(255) NOT NULL UNIQUE,
    "password" text NOT NULL,
    "email" character varying(255) NOT NULL UNIQUE,
    "creation_time" timestamp WITH time zone,
    "modification_time" timestamp WITH time zone
);

COMMIT;