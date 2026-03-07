CREATE TABLE IF NOT EXISTS users (
    id serial primary key,
    name text not null,
    email text not null unique,
    password_hash text not null,
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);

CREATE TABLE IF NOT EXISTS refresh_tokens (
    id serial primary key,
    user_id int not null references users(id) on delete cascade,
    token_hash text not null,
    expires_at timestamptz default now()
);

CREATE INDEX idx_refresh_tokens_token ON refresh_tokens(token_hash);