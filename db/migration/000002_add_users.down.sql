ALTER TABLE IF EXISTS "account" DROP CONSTRAINT IF EXISTS "name_currency_key";

ALTER TABLE IF EXISTS "account" DROP CONSTRAINT IF EXISTS "account_name_fkey";

DROP TABLE IF EXISTS "user";