ALTER TABLE "items" DROP COLUMN IF EXISTS "group_id";
ALTER TABLE "items" DROP COLUMN IF EXISTS "flat_id";
ALTER TABLE "items" DROP COLUMN IF EXISTS "status";

ALTER TABLE "items" ADD COLUMN "user_id" bigint NOT NULL default 0;

CREATE UNIQUE INDEX ON "items" ("title", "user_id");

CREATE TABLE "flat_group_items" (
    "flat_id" bigint NOT NULL,
    "group_id" bigint NOT NULL,
    "item_id" bigint NOT NULL,
    "status" varchar NOT NULL
);

CREATE UNIQUE INDEX ON "flat_group_items" ("flat_id", "group_id", "item_id");

ALTER TABLE "flat_group_items" ADD FOREIGN KEY ("flat_id") REFERENCES "flats" ("id") ON DELETE CASCADE;

ALTER TABLE "flat_group_items" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id") ON DELETE CASCADE;

ALTER TABLE "flat_group_items" ADD FOREIGN KEY ("item_id") REFERENCES "items" ("id") ON DELETE CASCADE;
