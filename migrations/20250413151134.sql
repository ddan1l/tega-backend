-- Modify "tokens" table
ALTER TABLE "tokens" ADD COLUMN "expiresd_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP;
