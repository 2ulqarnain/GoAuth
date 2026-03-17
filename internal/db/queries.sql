-- name: CreateUser :one
INSERT INTO users (name, email, password_hash)
VALUES ($1, $2, $3)
RETURNING id, name, email;

-- name: GetUserByEmail :one
select name, email, password_hash from users where email = $1;

-- name: InsertRefreshToken :one
INSERT INTO refresh_tokens (user_id, parent_id, token_hash, expires_at, user_agent, user_ip)
VALUES ($1,$2,$3,$4,$5,$6);

-- name: 