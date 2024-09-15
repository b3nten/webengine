-- name: AddUser :one
INSERT INTO users (username, email) VALUES (?, ?) RETURNING *;

-- name: GetUserById :one
SELECT * FROM users WHERE id = ?;

-- name: GetUsersByUsername :many
SELECT * FROM users WHERE username = ?;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = ?;