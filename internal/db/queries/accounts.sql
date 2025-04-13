-- name: CreateAccount :one
INSERT INTO accounts (name, created_at)
VALUES ($1, NOW())
RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id DESC
LIMIT $1 OFFSET $2;
