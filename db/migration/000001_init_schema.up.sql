CREATE TABLE "account" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar(255) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE "entry" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE "transfer" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now()
);

ALTER TABLE "entry" ADD CONSTRAINT "entry_account_id_fkey" FOREIGN KEY ("account_id") REFERENCES "account" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "transfer" ADD CONSTRAINT "transfer_from_account_id_fkey" FOREIGN KEY ("from_account_id") REFERENCES "account" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
    ADD CONSTRAINT "transfer_to_account_id_fkey" FOREIGN KEY ("to_account_id") REFERENCES "account" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

CREATE INDEX "account_name_idx" ON "account" ("name");

CREATE INDEX "entry_account_id_idx" ON "entry" ("account_id");

CREATE INDEX "transfer_from_account_id_idx" ON "transfer" ("from_account_id");

CREATE INDEX "transfer_to_account_id_idx" ON "transfer" ("to_account_id");

CREATE INDEX "transfer_from_account_id_to_account_id_idx" ON "transfer" ("from_account_id", "to_account_id");

