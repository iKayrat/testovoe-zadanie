-- Users_table
CREATE TABLE IF NOT EXISTS "users" (
  id bigserial PRIMARY KEY,
  firstname varchar NOT NULL,
  lastname varchar NOT NULL,
  username varchar NOT NULL,
  email varchar NOT NULL,
  password varchar NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT (now())
);
