-- name: CreatePSHelves :one
INSERT INTO Product_Shelves (id, shelf_id, is_primary)
VALUES ($1,$2,$3) RETURNING *;

-- name: GetPShelves :many
SELECT * FROM Product_Shelves;

-- name: GetPShelvesByID :one
SELECT * FROM Product_Shelves WHERE id = $1;

-- name: GetPShelvesByShelfID :one
SELECT * FROM Product_Shelves WHERE shelf_id = $1;

-- name: GetPShelvesByIDs :one
SELECT * FROM Product_Shelves WHERE id = $1 AND shelf_id = $2;

-- name: UpdatePShelves :one
UPDATE Product_Shelves
SET is_primary = false
WHERE id = $1 AND shelf_id = $2
RETURNING *;

-- name: DeletePShelves :one
DELETE FROM Product_Shelves WHERE id = $1 RETURNING id;
