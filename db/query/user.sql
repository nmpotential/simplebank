-- name: CreateUser :one
-- Inserts a new user record into the 'users' table with provided details
INSERT INTO users (
  username,          -- User's username
  hashed_password,   -- Hashed password for security
  full_name,         -- User's full name
  email              -- User's email address
) VALUES (
  $1,                -- Placeholder for username
  $2,                -- Placeholder for hashed_password
  $3,                -- Placeholder for full_name
  $4                 -- Placeholder for email
) RETURNING *;       -- Returns the newly created user record

-- name: GetUser :one
-- Retrieves a single user record from 'users' table based on username
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: UpdateUser :one
-- Updates a user record with provided data if not null (COALESCE used to manage NULL values)
UPDATE users
SET
  hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),   -- Updates hashed_password if provided
  password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at), -- Updates password_changed_at if provided
  full_name = COALESCE(sqlc.narg(full_name), full_name),          -- Updates full_name if provided
  email = COALESCE(sqlc.narg(email), email),                      -- Updates email if provided
  is_email_verified = COALESCE(sqlc.narg(is_email_verified), is_email_verified) -- Updates is_email_verified if provided
WHERE
  username = sqlc.arg(username)     -- Identifies the user by username
RETURNING *;                        -- Returns the updated user record
