#!/bin/bash
set -e
echo "Creating Database"

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGREST_DB" <<-EOSQL
  GRANT ALL PRIVILEGES ON DATABASE dev TO dev;

  CREATE TABLE hello_world_table (
    id serial PRIMARY KEY,
    description VARCHAR (50)
  );

  CREATE TABLE user_table (
    user_id serial PRIMARY KEY,
    email VARCHAR (50) UNIQUE NOT NULL,
    password VARCHAR (50) NOT NULL
  );

  -- USER TABLE
  INSERT INTO user_table (email, password) VALUES
  ('test1@gmail.com', 'password'),
  ('test2@gmail.com', 'password');

  -- HELLO WORLD TABLE
  INSERT INTO hello_world_table (description) VALUES
  ('Hello World');

EOSQL

echo "Complete Database"
