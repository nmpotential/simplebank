-- name: CreateAccount :one
-- Inserts a new account into the 'accounts' table with provided details
INSERT INTO accounts (
  owner,      -- Account owner's name
  balance,    -- Initial balance
  currency    -- Currency type
) VALUES (
  $1,         -- Placeholder for owner
  $2,         -- Placeholder for balance
  $3          -- Placeholder for currency
) RETURNING *; -- Returns the newly created account

-- name: GetAccount :one
-- Retrieves a single account from 'accounts' table based on ID
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
-- Retrieves a single account for update from 'accounts' table based on ID
SELECT * FROM accounts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListAccounts :many
-- Retrieves multiple accounts based on owner, limited by a count and offset
SELECT * FROM accounts
WHERE owner = $1
ORDER BY id
LIMIT $2 -- Limits the number of results
OFFSET $3; -- Specifies the starting point for results

-- name: UpdateAccount :one
-- Updates the balance of a single account based on ID
UPDATE accounts
SET balance = $2
WHERE id = $1
RETURNING *; -- Returns the updated account details

-- name: AddAccountBalance :one
-- Adds a specified amount to an account's balance based on ID
UPDATE accounts
SET balance = balance + sqlc.arg(amount) -- Increases the balance by the provided amount
WHERE id = sqlc.arg(id) -- Identifies the account by ID
RETURNING *; -- Returns the updated account details

-- name: DeleteAccount :exec
-- Deletes a specific account from the 'accounts' table based on ID
DELETE FROM accounts
WHERE id = $1; -- Specifies the ID of the account to be deleted
