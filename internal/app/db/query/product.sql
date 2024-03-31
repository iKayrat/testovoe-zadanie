-- name: CreateProduct :one
INSERT INTO "products" (
	title,
	active,
	price,
	description
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetProduct :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

-- name: ImportProducts :exec
INSERT INTO products(
    id,
	title,
	active,
	price,
	description,
	created_at,
	updated_at
) VALUES(
    $1,$2,$3,$4,$5,$6,$7
);

-- name: ListProducts :many
SELECT * FROM products
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateProduct :one
UPDATE products
SET
	title = $2,
	active = $3,
	price = $4,
	description = $5,
	updated_at = $6
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :one
DELETE FROM products
WHERE id = $1
RETURNING id;