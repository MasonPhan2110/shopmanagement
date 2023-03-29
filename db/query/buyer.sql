-- name: CreateBuyer :one
INSERT INTO "buyer" (
  name
) VALUES (
  $1
)
RETURNING *;

-- name: GetBuyer :one
SELECT * FROM "buyer"
WHERE name = $1 LIMIT 1;

-- name: ListBuyer :many
SELECT * FROM "buyer"
ORDER BY id
LIMIT $1
OFFSET $2;


