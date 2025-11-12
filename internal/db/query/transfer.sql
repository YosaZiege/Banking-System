-- name: CreateTransfer :one
INSERT INTO transfers (from_client, to_client, amout)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1;

-- name: ListTransfersByClient :many
SELECT * FROM transfers
WHERE from_client = $1 OR to_client = $1
ORDER BY created_at DESC;

-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = $1;

