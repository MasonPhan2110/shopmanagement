-- name: CreateBuyer :one
INSERT INTO "buyer" (
  name,
  address
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetBuyer :one
SELECT * FROM "buyer"
WHERE name = $1 LIMIT 1;

-- name: GetProductForUpdate :one
SELECT * FROM "buyer"
WHERE name = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListBuyer :many
SELECT * FROM "buyer"
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateBuyer :one
UPDATE "buyer"
SET address = $2
WHERE id = $1
RETURNING *;

