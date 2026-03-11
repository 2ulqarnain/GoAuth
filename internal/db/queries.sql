-- name: CreateUser :one
INSERT INTO users (name, email, password_hash)
VALUES ($1, $2, $3)
RETURNING id, name, email;

-- name: GetUserByEmail :one
select name, email, password_hash from users where email = $1;