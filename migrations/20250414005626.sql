-- Modify "users" table
ALTER TABLE "users" ADD CONSTRAINT "uni_users_email" UNIQUE ("email");
