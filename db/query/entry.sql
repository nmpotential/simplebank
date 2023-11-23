-- name: CreateEntry :one
-- Inserts a new entry into the 'entries' table with provided details
INSERT INTO entries (
  account_id, -- ID of the account associated with the entry
  amount      -- Amount of the entry
) VALUES (
  $1,         -- Placeholder for account_id
  $2          -- Placeholder for amount
) RETURNING *; -- Returns the newly created entry

-- name: GetEntry :one
-- Retrieves a single entry from 'entries' table based on ID
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
-- Retrieves multiple entries based on account_id, limited by a count and offset
SELECT * FROM entries
WHERE account_id = $1 -- Filters entries based on the account ID
ORDER BY id
LIMIT $2 -- Limits the number of results
OFFSET $3; -- Specifies the starting point for results
