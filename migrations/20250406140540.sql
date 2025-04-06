-- Create "accesses" table
CREATE TABLE "accesses" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_accesses_deleted_at" to table: "accesses"
CREATE INDEX "idx_accesses_deleted_at" ON "accesses" ("deleted_at");
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
-- Create "role_accesses" table
CREATE TABLE "role_accesses" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "role_id" bigint NULL,
  "access_id" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_role_accesses_access" FOREIGN KEY ("access_id") REFERENCES "accesses" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_role_accesses_role" FOREIGN KEY ("role_id") REFERENCES "roles" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_role_accesses_deleted_at" to table: "role_accesses"
CREATE INDEX "idx_role_accesses_deleted_at" ON "role_accesses" ("deleted_at");
-- Create "users" table
CREATE TABLE "users" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" text NULL,
  "email" text NULL,
  "password_hash" text NULL,
  "last_name" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "users" ("deleted_at");
-- Create "user_roles" table
CREATE TABLE "user_roles" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "role_id" bigint NULL,
  "user_id" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_user_roles_role" FOREIGN KEY ("role_id") REFERENCES "roles" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_user_roles_user" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_user_roles_deleted_at" to table: "user_roles"
CREATE INDEX "idx_user_roles_deleted_at" ON "user_roles" ("deleted_at");
