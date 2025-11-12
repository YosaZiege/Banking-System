CREATE TABLE "clients" (
  "username" varchar PRIMARY KEY,
  "balance" bigint NOT NULL,
  "currency" bi NOT NULL,
  "email" varchar NOT NULL,
  "password_hash" varchar NOT NULL,
  "provider" varchar NOT NULL DEFAULT 'Google',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "client_username" varchar,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_client" varchar,
  "to_client" varchar,
  "amout" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "entries" ADD FOREIGN KEY ("client_username") REFERENCES "clients" ("username");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_client") REFERENCES "clients" ("username");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_client") REFERENCES "clients" ("username");










