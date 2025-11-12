-- name: CreateClient :one
INSERT INTO clients (username, balance, currency, email, password_hash, provider)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetClient :one
SELECT * FROM clients
WHERE username = $1;

-- name: UpdateClientBalance :one
UPDATE clients
SET balance = $2
WHERE username = $1
RETURNING *;

-- name: DeleteClient :exec
DELETE FROM clients
WHERE username = $1;

-- name: ListClients :many
SELECT * FROM clients
ORDER BY created_at DESC LIMIT $1 OFFSET $2;

 
