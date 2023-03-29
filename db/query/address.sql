-- name: CreateBuyerAddress :one
INSERT INTO "address" (
  buyer_id,
  address
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetAddress :one
SELECT * FROM "address"
WHERE buyer_id = $1 AND address = $2
LIMIT 1;

-- name: ListAddress :many
SELECT * FROM "address"
WHERE buyer_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;