CREATE TABLE "users" (
                         "id" bigserial PRIMARY KEY,
                         "login" varchar UNIQUE NOT NULL,
                         "password_hash" varchar NOT NULL,
                         "auth_token" varchar UNIQUE
);

CREATE TABLE "flats" (
                         "id" bigserial PRIMARY KEY,
                         "title" varchar,
                         "description" varchar,
                         "address" varchar NOT NULL,
                         "owner_id" bigint NOT NULL
);

CREATE TABLE "user_flats" (
                              "user_id" bigint,
                              "flat_id" bigint,
                              PRIMARY KEY ("user_id", "flat_id")
);

CREATE TABLE "groups" (
                          "id" bigserial PRIMARY KEY,
                          "user_id" bigint NOT NULL,
                          "title" varchar NOT NULL,
                          "hide" bool NOT NULL
);

CREATE TABLE "flat_groups" (
                               "flat_id" bigint NOT NULL,
                               "group_id" bigint NOT NULL,
                               PRIMARY KEY ("flat_id", "group_id")
);

CREATE TABLE "items" (
                         "id" bigserial PRIMARY KEY,
                         "group_id" bigint NOT NULL,
                         "flat_id" bigint NOT NULL,
                         "title" varchar NOT NULL,
                         "description" varchar,
                         "hide" bool NOT NULL,
                         "status" varchar NOT NULL
);

CREATE INDEX ON "users" ("auth_token");

CREATE INDEX ON "users" ("login");

CREATE UNIQUE INDEX ON "flats" ("address", "owner_id");

CREATE UNIQUE INDEX ON "groups" ("title", "user_id");

CREATE UNIQUE INDEX ON "items" ("flat_id", "group_id", "title");

ALTER TABLE "flats" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "user_flats" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "user_flats" ADD FOREIGN KEY ("flat_id") REFERENCES "flats" ("id") ON DELETE CASCADE;

ALTER TABLE "groups" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "flat_groups" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id") ON DELETE CASCADE;

ALTER TABLE "flat_groups" ADD FOREIGN KEY ("flat_id") REFERENCES "flats" ("id") ON DELETE CASCADE;

ALTER TABLE "items" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id") ON DELETE CASCADE;

ALTER TABLE "items" ADD FOREIGN KEY ("flat_id") REFERENCES "flats" ("id") ON DELETE CASCADE;
