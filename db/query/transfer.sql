-- name: CreateTransfer :one
-- Inserts a new transfer record into the 'transfers' table with details of the transaction
INSERT INTO transfers (
  from_account_id,  -- ID of the account from which the transfer is initiated
  to_account_id,    -- ID of the account receiving the transfer
  amount            -- Amount being transferred
) VALUES (
  $1,               -- Placeholder for from_account_id
  $2,               -- Placeholder for to_account_id
  $3                -- Placeholder for amount
) RETURNING *;      -- Returns the newly created transfer record

-- name: GetTransfer :one
-- Retrieves a single transfer record from 'transfers' table based on its unique ID
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
-- Retrieves multiple transfer records from 'transfers' table based on account IDs with pagination
SELECT * FROM transfers
WHERE 
    from_account_id = $1 OR  -- Filters transfers by the sender's account ID
    to_account_id = $2       -- Filters transfers by the receiver's account ID
ORDER BY id                -- Orders the results by ID
LIMIT $3                   -- Limits the number of results returned
OFFSET $4;                 -- Specifies the starting point for results (pagination)
