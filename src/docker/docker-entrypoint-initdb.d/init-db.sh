#!/bin/bash
set -e
echo "Creating Database"

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGREST_DB" <<-EOSQL
  GRANT ALL PRIVILEGES ON DATABASE dev TO dev;

  CREATE TABLE user_table (
    user_id serial PRIMARY KEY,
    email VARCHAR (50) UNIQUE NOT NULL,
    password VARCHAR (50) NOT NULL,
    created_on TIMESTAMP NOT NULL,
    last_login TIMESTAMP
  );

EOSQL

echo "Complete Database"
