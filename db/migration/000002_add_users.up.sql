-- Creates a table named "users" to store user information
CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,                   -- Unique username as the primary key
  "hashed_password" varchar NOT NULL,               -- Hashed password for security
  "full_name" varchar NOT NULL,                     -- Full name of the user
  "email" varchar UNIQUE NOT NULL,                  -- Unique email address for the user
  "password_changed_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'),  -- Time of password change
  "created_at" timestamptz NOT NULL DEFAULT (now()) -- Time of user creation
);

-- Adds a foreign key constraint to ensure the "owner" column in "accounts" references the "username" column in "users" table
ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

-- Adds a unique constraint on the combination of "owner" and "currency" columns in the "accounts" table
ALTER TABLE "accounts" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner", "currency");
