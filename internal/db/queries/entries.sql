-- name: CreateEntry :one
INSERT INTO entries (account_id, amount, created_at)
VALUES ($1, $2, NOW())
RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1;

-- name: ListEntries :many
SELECT * FROM entries
WHERE account_id = $1
ORDER BY id DESC
LIMIT $2 OFFSET $3;
