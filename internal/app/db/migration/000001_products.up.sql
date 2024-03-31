-- Products_table
CREATE TABLE IF NOT EXISTS Products (
  id bigserial PRIMARY KEY,
  title varchar NOT NULL,
  active boolean,
  price float,
  description varchar,
  created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT (now())
);

