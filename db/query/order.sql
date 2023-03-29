-- name: CreateOrder :one
INSERT INTO "order" (
  buyer_id,
  product_id,
  amount,
  unit
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetOrder :one
SELECT * FROM "order"
WHERE id = $1 LIMIT 1;

-- name: ListOrder :many
SELECT * FROM "order"
WHERE buyer_id = $1 OR product_id = $2
ORDER BY id
LIMIT $3
OFFSET $4;