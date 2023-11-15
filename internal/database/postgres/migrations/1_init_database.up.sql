CREATE TABLE "users" (
                         "id" bigserial PRIMARY KEY,
                         "auth_token" varchar UNIQUE NOT NULL
);

CREATE TABLE "flats" (
                         "id" bigserial PRIMARY KEY,
                         "title" varchar UNIQUE,
                         "description" varchar,
                         "address" varchar UNIQUE NOT NULL
);

CREATE TABLE "groups" (
                          "id" bigserial PRIMARY KEY,
                          "title" varchar UNIQUE NOT NULL,
                          "user_id" integer NOT NULL,
                          "hide" bool NOT NULL
);

CREATE TABLE "items" (
                         "id" bigserial PRIMARY KEY,
                         "title" varchar UNIQUE NOT NULL,
                         "description" varchar,
                         "hide" bool NOT NULL
);

CREATE TABLE "user_flats" (
                              "user_id" integer,
                              "flat_id" integer,
                              PRIMARY KEY ("user_id", "flat_id")
);

CREATE TABLE "flat_groups" (
                               "flat_id" integer,
                               "group_id" integer,
                               PRIMARY KEY ("flat_id", "group_id")
);

CREATE TABLE "flat_group_items" (
                                    "flat_id" integer,
                                    "group_id" integer,
                                    "item_id" integer,
                                    "status" varchar NOT NULL,
                                    PRIMARY KEY ("flat_id", "group_id", "item_id")
);

CREATE INDEX ON "users" ("auth_token");

ALTER TABLE "groups" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_flats" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_flats" ADD FOREIGN KEY ("flat_id") REFERENCES "flats" ("id");

ALTER TABLE "flat_groups" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");

ALTER TABLE "flat_groups" ADD FOREIGN KEY ("flat_id") REFERENCES "flats" ("id");

ALTER TABLE "flat_group_items" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");

ALTER TABLE "flat_group_items" ADD FOREIGN KEY ("item_id") REFERENCES "items" ("id");

ALTER TABLE "flat_group_items" ADD FOREIGN KEY ("flat_id") REFERENCES "flats" ("id");
