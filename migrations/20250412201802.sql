-- Rename a column from "first_name" to "full_name"
ALTER TABLE "users" RENAME COLUMN "first_name" TO "full_name";
-- Modify "users" table
ALTER TABLE "users" DROP COLUMN "last_name";
