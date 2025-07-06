-- Create "user" table
CREATE TABLE "public"."user" (
  "id" bigserial NOT NULL,
  "name" text NULL,
  "email" text NULL,
  "phone" text NULL,
  "password" text NULL,
  "status" bigint NULL,
  "updated_at" timestamptz NULL,
  "created_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_user_email" to table: "user"
CREATE UNIQUE INDEX "idx_user_email" ON "public"."user" ("email");
