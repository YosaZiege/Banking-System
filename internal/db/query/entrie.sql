-- name: CreateEntry :one
INSERT INTO entries (client_username, amout)
VALUES ($1, $2)
RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1;

-- name: ListEntriesByClient :many
SELECT * FROM entries
WHERE client_username = $1
ORDER BY created_at DESC;

-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = $1;

