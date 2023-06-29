-- name: CreateUser :execresult
INSERT INTO users (id, name, created_at, updated_at, api_key)
VALUES (?, ?, ?, ?, UUID());

-- name: GetUserByID :one
SELECT * FROM users WHERE id = ?;

-- name: GetUserByAPIKey :one
SELECT * FROM users WHERE api_key = ?;