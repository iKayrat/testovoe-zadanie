-- name: CreateShelf :one
INSERT INTO Shelves (shelf_name) VALUES ($1)RETURNING shelf_id;

-- name: GetShelfs :many
SELECT * FROM Shelves;

-- name: GetShelfByID :one
SELECT * FROM Shelves WHERE shelf_id = 1;

-- name: GetShelfByName :one
SELECT * FROM Shelves WHERE shelf_name = $1;

-- name: UpdateShelf :exec
UPDATE Shelves
SET shelf_name = $2
WHERE shelf_id = $1
RETURNING *;

-- name: DeleteShelf :one
DELETE FROM Shelves WHERE shelf_id = $1
RETURNING shelf_id;
