-- Create "projects" table
CREATE TABLE "projects" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" text NULL,
  "slug" text NULL,
  "description" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_projects_deleted_at" to table: "projects"
CREATE INDEX "idx_projects_deleted_at" ON "projects" ("deleted_at");
-- Create "roles" table
CREATE TABLE "roles" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "title" text NULL,
  "slug" text NULL,
  "project_id" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_roles_project" FOREIGN KEY ("project_id") REFERENCES "projects" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_roles_deleted_at" to table: "roles"
CREATE INDEX "idx_roles_deleted_at" ON "roles" ("deleted_at");
-- Create "policies" table
CREATE TABLE "policies" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "slug" character varying(100) NOT NULL,
  "effect" character varying(20) NOT NULL DEFAULT 'allow',
  "role_id" bigint NOT NULL,
  "project_id" bigint NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_policies_project" FOREIGN KEY ("project_id") REFERENCES "projects" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_policies_role" FOREIGN KEY ("role_id") REFERENCES "roles" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "chk_policies_effect" CHECK ((effect)::text = ANY ((ARRAY['allow'::character varying, 'deny'::character varying])::text[]))
);
-- Create index "idx_policies_deleted_at" to table: "policies"
CREATE INDEX "idx_policies_deleted_at" ON "policies" ("deleted_at");
-- Create index "idx_policies_slug" to table: "policies"
CREATE UNIQUE INDEX "idx_policies_slug" ON "policies" ("slug");
-- Create "policy_actions" table
CREATE TABLE "policy_actions" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "policy_id" bigint NOT NULL,
  "action" character varying(50) NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_policies_actions" FOREIGN KEY ("policy_id") REFERENCES "policies" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_policy_actions_deleted_at" to table: "policy_actions"
CREATE INDEX "idx_policy_actions_deleted_at" ON "policy_actions" ("deleted_at");
-- Create index "idx_policy_actions_policy_id" to table: "policy_actions"
CREATE INDEX "idx_policy_actions_policy_id" ON "policy_actions" ("policy_id");
-- Create "policy_conditions" table
CREATE TABLE "policy_conditions" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "policy_id" bigint NOT NULL,
  "field" character varying(100) NOT NULL,
  "operator" character varying(20) NOT NULL,
  "value" character varying(255) NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_policies_conditions" FOREIGN KEY ("policy_id") REFERENCES "policies" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_policy_conditions_deleted_at" to table: "policy_conditions"
CREATE INDEX "idx_policy_conditions_deleted_at" ON "policy_conditions" ("deleted_at");
-- Create index "idx_policy_conditions_policy_id" to table: "policy_conditions"
CREATE INDEX "idx_policy_conditions_policy_id" ON "policy_conditions" ("policy_id");
-- Create "policy_resources" table
CREATE TABLE "policy_resources" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "policy_id" bigint NOT NULL,
  "resource" character varying(50) NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_policies_resources" FOREIGN KEY ("policy_id") REFERENCES "policies" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_policy_resources_deleted_at" to table: "policy_resources"
CREATE INDEX "idx_policy_resources_deleted_at" ON "policy_resources" ("deleted_at");
-- Create index "idx_policy_resources_policy_id" to table: "policy_resources"
CREATE INDEX "idx_policy_resources_policy_id" ON "policy_resources" ("policy_id");
-- Create "users" table
CREATE TABLE "users" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "full_name" text NULL,
  "email" text NULL,
  "password_hash" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_users_email" UNIQUE ("email")
);
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "users" ("deleted_at");
-- Create "project_users" table
CREATE TABLE "project_users" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "user_id" bigint NULL,
  "project_id" bigint NULL,
  "role_id" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_project_users_project" FOREIGN KEY ("project_id") REFERENCES "projects" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_project_users_role" FOREIGN KEY ("role_id") REFERENCES "roles" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_project_users_user" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_project_users_deleted_at" to table: "project_users"
CREATE INDEX "idx_project_users_deleted_at" ON "project_users" ("deleted_at");
-- Create "tokens" table
CREATE TABLE "tokens" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "token" text NULL,
  "user_id" bigint NULL,
  "expires_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_tokens_user" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_tokens_deleted_at" to table: "tokens"
CREATE INDEX "idx_tokens_deleted_at" ON "tokens" ("deleted_at");
