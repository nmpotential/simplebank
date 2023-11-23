-- Creates a table named "sessions" to store session-related information
CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,                    -- Unique identifier for a session
  "username" varchar NOT NULL,              -- Username associated with the session
  "refresh_token" varchar NOT NULL,         -- Refresh token for session management
  "user_agent" varchar NOT NULL,            -- User agent information
  "client_ip" varchar NOT NULL,              -- Client IP address
  "is_blocked" boolean NOT NULL DEFAULT false,  -- Indicates if the session is blocked
  "expires_at" timestamptz NOT NULL,        -- Expiration time for the session
  "created_at" timestamptz NOT NULL DEFAULT (now())  -- Creation time for the session
);

-- Adds a foreign key constraint to ensure the "username" column in "sessions" references the "username" column in "users" table
ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");
