CREATE TABLE IF NOT EXISTS users (
    id serial primary key,
    name text not null,
    email text not null unique,
    password_hash text not null,
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);

CREATE TABLE IF NOT EXISTS refresh_tokens (
    id uuid primary key default gen_random_uuid(),
    user_id integer not null references users(id),
    parent_id uuid not null,
    token_hash text not null unique,
    revoked bool not null default false,
    created_at timestamptz default now(),
    expires_at timestamptz not null,
    user_agent text,
    user_ip text
);

comment on column refresh_tokens.user_ip is 'user ip address if available';
comment on column refresh_tokens.user_agent is 'user device information if available';

CREATE INDEX idx_refresh_tokens_token ON refresh_tokens(token_hash);