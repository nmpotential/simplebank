-- Drops the table "verify_emails" and all dependent objects if it exists, ensuring a clean removal
DROP TABLE IF EXISTS "verify_emails" CASCADE;

-- Removes the column "is_email_verified" from the "users" table
ALTER TABLE "users" DROP COLUMN "is_email_verified";
