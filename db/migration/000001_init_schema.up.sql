CREATE TABLE "accounts" (
    "id" bigserial PRIMARY KEY, 
    "owner" varchar NOT NULL, 
    "balance" bigint NOT NULL, 
    "currency" varchar NOT NULL,
    "create_at" timestamptz NOT NULL DEFAULT (now())
);



CREATE TABLE authors (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  bio  text
);
