-- name: ListProducts :many
SELECT 
  *
FROM
  products;

-- name: FindProductByID :one
SELECT 
  * 
FROM 
  products 
WHERE 
  id = $1;

-- name: CreateProduct :one
INSERT INTO products (
  name, price_in_center, quantity
) VALUES (
  $1, $2, $3
)
RETURNING *;