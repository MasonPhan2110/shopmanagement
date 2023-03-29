 -- name: CreateProduct :one
INSERT INTO "product" (
  type,
  name,
  amount
) VALUES (
  $1, $2, $3
) RETURNING *;


-- name: GetProduct :one
SELECT * FROM "product"
WHERE id = $1 LIMIT 1;

-- name: GetProductForUpdate :one
SELECT * FROM "product"
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListProduct :many
SELECT * FROM "product"
WHERE name = $1
ORDER BY id
LIMIT $2
OFFSET $3;


-- name: AddProductAmount :one
UPDATE "product"
SET amount = amount + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;

