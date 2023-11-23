-- Creates the 'accounts' table with columns: id (auto-incrementing primary key), owner (string), balance (integer), currency (string), and created_at (timestamp with time zone)
CREATE TABLE accounts (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

-- Creates the 'entries' table with columns: id (auto-incrementing primary key), account_id (integer), amount (integer), and created_at (timestamp with time zone)
CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

-- Creates the 'transfers' table with columns: id (auto-incrementing primary key), from_account_id (integer), to_account_id (integer), amount (integer), and created_at (timestamp with time zone)
CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

-- Creates indexes for the 'accounts' table on the 'owner' column
CREATE INDEX ON "accounts" ("owner");

-- Creates indexes for the 'entries' table on the 'account_id' column
CREATE INDEX ON "entries" ("account_id");

-- Creates indexes for the 'transfers' table on the 'from_account_id', 'to_account_id', and combination of 'from_account_id' and 'to_account_id'
CREATE INDEX ON "transfers" ("from_account_id");
CREATE INDEX ON "transfers" ("to_account_id");
CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

-- Adds a comment on the 'amount' column in the 'entries' table describing it can hold negative or positive values
COMMENT ON COLUMN "entries"."amount" IS 'This can be a negative or positive value';

-- Adds a comment on the 'amount' column in the 'transfers' table specifying that it must be a positive value
COMMENT ON COLUMN "transfers"."amount" IS 'This must be a positive value';

-- Adds foreign key constraints on the 'account_id' column in the 'entries' table, and on the 'from_account_id' and 'to_account_id' columns in the 'transfers' table, referencing the 'id' column in the 'accounts' table
ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");
ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");
ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");
