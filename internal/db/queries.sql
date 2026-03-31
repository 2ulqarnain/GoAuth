-- name: CreateUser :one
INSERT INTO users (name, email, password_hash)
VALUES ($1, $2, $3)
RETURNING id, name, email;

-- name: GetUserByEmail :one
select name, email, password_hash from users where email = $1;

-- name: InsertRefreshToken :one
INSERT INTO refresh_tokens (user_id, parent_id, token_hash, expires_at, user_agent, user_ip)
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING id, user_id, expires_at;

-- name: GetRefreshToken :one
SELECT id,user_id,token_hash,is_revoked,expires_at
FROM refresh_tokens
WHERE token_hash = $1;