-- name: GetEntry :one
SELECT * FROM entry
WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entry
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateEntry :one
INSERT INTO entry (
  account_id, amount
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateEntry :exec
UPDATE entry
SET amount = $2
WHERE id = $1
RETURNING *;

-- name: DeleteEntry :exec
DELETE FROM entry WHERE id = $1;