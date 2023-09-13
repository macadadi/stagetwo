-- +goose Up

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "name" varchar,
  "created_at" timestamptz
);


-- +goose Down
DROP TABLE IF EXISTS users;