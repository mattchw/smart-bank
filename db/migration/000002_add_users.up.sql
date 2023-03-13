CREATE TABLE "user" (
  "username" varchar(255) PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'),  
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "account" ADD FOREIGN KEY ("name") REFERENCES "user" ("username");

-- CREATE UNIQUE INDEX ON "account" ("name", "currency");
ALTER TABLE "account" ADD CONSTRAINT "name_currency_key" UNIQUE ("name", "currency");