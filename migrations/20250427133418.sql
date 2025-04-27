-- Create "roles" table
CREATE TABLE "roles" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_roles_deleted_at" to table: "roles"
CREATE INDEX "idx_roles_deleted_at" ON "roles" ("deleted_at");
-- Modify "project_users" table
ALTER TABLE "project_users" ADD COLUMN "role_id" bigint NULL, ADD CONSTRAINT "fk_project_users_role" FOREIGN KEY ("role_id") REFERENCES "roles" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
