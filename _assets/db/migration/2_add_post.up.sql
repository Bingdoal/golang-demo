BEGIN;

CREATE TABLE IF NOT EXISTS "post" (
    "id" bigserial NOT NULL PRIMARY KEY,
    "content" text NOT NULL,
    "author_id" bigint NOT NULL,
    "creation_time" timestamp WITH time zone,
    "modification_time" timestamp WITH time zone
);

ALTER TABLE
    "post"
ADD
    CONSTRAINT "post_author" FOREIGN KEY ("author_id") REFERENCES "user" ("id");

COMMIT;