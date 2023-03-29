 -- name: CreateInventory :one
INSERT INTO "inventory" (
  type,
  name,
  amount
) VALUES (
  $1, $2, $3
) RETURNING *;


-- name: GetInventory :one
SELECT * FROM "inventory"
WHERE id = $1 LIMIT 1;

-- name: GetInventoryForUpdate :one
SELECT * FROM "inventory"
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListInventory :many
SELECT * FROM "inventory"
WHERE name = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateInventory :one
UPDATE "inventory"
SET amount = $2
WHERE id = $1
RETURNING *;

-- name: AddInventoryAmount :one
UPDATE "inventory"
SET amount = amount + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;

