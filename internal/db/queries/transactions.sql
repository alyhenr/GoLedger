-- name: CreateTransaction :one
INSERT INTO transactions (reference, created_at)
VALUES ($1, NOW())
RETURNING *;

-- name: GetTranscation :one
SELECT * FROM transactions
WHERE id = $1;

-- name: ListTransactions :many
SELECT * FROM transactions
ORDER BY id DESC
LIMIT $1 OFFSET $2;
