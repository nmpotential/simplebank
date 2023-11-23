-- name: CreateVerifyEmail :one
-- Inserts a new record into the 'verify_emails' table with username, email, and secret_code values
INSERT INTO verify_emails (
    username,     -- Username associated with the email verification
    email,        -- Email to be verified
    secret_code   -- Secret code for verification
) VALUES (
    $1,           -- Placeholder for username
    $2,           -- Placeholder for email
    $3            -- Placeholder for secret_code
) RETURNING *;    -- Returns the newly created verification email record

-- name: UpdateVerifyEmail :one
-- Updates a verification email record setting is_used to TRUE under specific conditions
UPDATE verify_emails
SET
    is_used = TRUE      -- Updates the is_used field to mark the verification as used
WHERE
    id = @id            -- Matches the ID of the record
    AND secret_code = @secret_code  -- Matches the secret_code provided
    AND is_used = FALSE -- Ensures the verification email hasn't been used
    AND expired_at > now() -- Checks if the verification hasn't expired
RETURNING *;            -- Returns the updated verification email record
