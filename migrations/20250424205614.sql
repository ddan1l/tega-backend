-- Create "projects" table
CREATE TABLE "projects" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" text NULL,
  "slug" text NULL,
  "expires_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id")
);
-- Create index "idx_projects_deleted_at" to table: "projects"
CREATE INDEX "idx_projects_deleted_at" ON "projects" ("deleted_at");
-- Create "project_users" table
CREATE TABLE "project_users" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "user_id" bigint NULL,
  "project_id" bigint NULL,
  "expires_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_project_users_project" FOREIGN KEY ("project_id") REFERENCES "projects" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_project_users_user" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_project_users_deleted_at" to table: "project_users"
CREATE INDEX "idx_project_users_deleted_at" ON "project_users" ("deleted_at");
