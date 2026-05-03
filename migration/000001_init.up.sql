CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS accounts(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    email text UNIQUE NOT NULL,
    password text NOT NULL
);

CREATE TYPE priority AS ENUM(
    'low',
    'medium',
    'high'
);

CREATE TABLE IF NOT EXISTS todos(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    account_id uuid NOT NULL REFERENCES accounts(id),
    title text NOT NULL,
    content text NOT NULL,
    priority priority NOT NULL,
    is_done boolean NOT NULL DEFAULT false,
    created_at timestamptz NOT NULL DEFAULT now(),

    UNIQUE (account_id, title)
);
