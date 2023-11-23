-- Removes the specified unique constraint "owner_currency_key" from the "accounts" table if it exists
ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "owner_currency_key";

-- Removes the specified foreign key constraint "accounts_owner_fkey" from the "accounts" table if it exists
ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "accounts_owner_fkey";

-- Drops the table "users" from the database if it exists
DROP TABLE IF EXISTS "users";
