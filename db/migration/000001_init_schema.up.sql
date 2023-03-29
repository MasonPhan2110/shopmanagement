CREATE TABLE "user" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NUll,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "role" varchar NOT NULL,
  "update_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "order" (
  "id" bigserial PRIMARY KEY,
  "buyer_id" bigserial NOT NULL,
  "product_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "unit" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "product" (
    "id" bigserial PRIMARY KEY,
    "type" varchar NOT NULL,
    "name" varchar UNIQUE NOT NULL,
    "amount" bigint NOT NULL,
    "unit" varchar NOT NULL,
    "update_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entry" (
    "id" bigserial PRIMARY KEY,
    "order_id" bigserial NOT NULL,
    "inventory_id" bigserial NOT NULL,
    "amount"  bigint NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "inventory" (
    "id" bigserial PRIMARY KEY,
    "type" varchar NOT NULL,
    "name" varchar UNIQUE NOT NULL,
    "amount" bigint NOT NULL,
    "update_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "buyer" (
    "id" bigserial PRIMARY KEY,
    "name" varchar UNIQUE NOT NULL,
    "address" varchar NOT NULL,
    "update_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "order" ("buyer_id");

CREATE INDEX ON "order" ("product_id");

CREATE INDEX ON "product" ("name");

ALTER TABLE "order" ADD FOREIGN KEY ("buyer_id") REFERENCES "buyer" ("id");

ALTER TABLE "order" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");

ALTER TABLE "entry" ADD FOREIGN KEY ("order_id") REFERENCES "order" ("id");

ALTER TABLE "entry" ADD FOREIGN KEY ("inventory_id") REFERENCES "inventory" ("id");
