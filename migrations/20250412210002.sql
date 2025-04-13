-- Drop index "idx_users_deleted_at" from table: "users"
DROP INDEX "idx_users_deleted_at";
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "users" ("deleted_at");
