-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4) RETURNING *;