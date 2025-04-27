-- Modify "project_users" table
ALTER TABLE "project_users" DROP COLUMN "expires_at";
-- Modify "projects" table
ALTER TABLE "projects" DROP COLUMN "expires_at";
