ALTER TABLE "items" ADD COLUMN "group_id" bigint NOT NULL default 0;
ALTER TABLE "items" ADD COLUMN "flat_id" bigint NOT NULL default 0;
ALTER TABLE "items" ADD COLUMN "status" varchar NOT NULL default 'not-ok';

ALTER TABLE "items" DROP COLUMN IF EXISTS "user_id" CASCADE;

DROP TABLE IF EXISTS "flat_group_items" CASCADE;