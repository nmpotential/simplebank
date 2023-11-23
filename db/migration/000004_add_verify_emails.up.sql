-- Creates a table named "verify_emails" to store email verification data
CREATE TABLE "verify_emails" (
  "id" bigserial PRIMARY KEY,                -- Unique identifier for email verification
  "username" varchar NOT NULL,               -- Username associated with the email
  "email" varchar NOT NULL,                  -- Email address
  "secret_code" varchar NOT NULL,            -- Secret code used for verification
  "is_used" bool NOT NULL DEFAULT false,     -- Indicates if the email is already used for verification
  "created_at" timestamptz NOT NULL DEFAULT (now()), -- Date and time of creation
  "expired_at" timestamptz NOT NULL DEFAULT (now() + interval '15 minutes') -- Expiration time for verification
);

-- Adds a foreign key constraint to ensure the "username" column in "verify_emails" references the "username" column in "users" table
ALTER TABLE "verify_emails" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

-- Adds a new column "is_email_verified" to the "users" table to track if the user's email is verified
ALTER TABLE "users" ADD COLUMN "is_email_verified" bool NOT NULL DEFAULT false;
