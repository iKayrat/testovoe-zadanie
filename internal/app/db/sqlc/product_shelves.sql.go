// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: product_shelves.sql

package db

import (
	"context"
	"database/sql"
)

const createPSHelves = `-- name: CreatePSHelves :one
INSERT INTO Product_Shelves (id, shelf_id, is_primary)
VALUES ($1,$2,$3) RETURNING id, shelf_id, is_primary
`

type CreatePSHelvesParams struct {
	ID        sql.NullInt32 `json:"id"`
	ShelfID   sql.NullInt32 `json:"shelf_id"`
	IsPrimary sql.NullBool  `json:"is_primary"`
}

func (q *Queries) CreatePSHelves(ctx context.Context, arg CreatePSHelvesParams) (ProductShelf, error) {
	row := q.db.QueryRowContext(ctx, createPSHelves, arg.ID, arg.ShelfID, arg.IsPrimary)
	var i ProductShelf
	err := row.Scan(&i.ID, &i.ShelfID, &i.IsPrimary)
	return i, err
}

const deletePShelves = `-- name: DeletePShelves :one
DELETE FROM Product_Shelves WHERE id = $1 RETURNING id
`

func (q *Queries) DeletePShelves(ctx context.Context, id sql.NullInt32) (sql.NullInt32, error) {
	row := q.db.QueryRowContext(ctx, deletePShelves, id)
	err := row.Scan(&id)
	return id, err
}

const getPShelves = `-- name: GetPShelves :many
SELECT id, shelf_id, is_primary FROM Product_Shelves
`

func (q *Queries) GetPShelves(ctx context.Context) ([]ProductShelf, error) {
	rows, err := q.db.QueryContext(ctx, getPShelves)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProductShelf
	for rows.Next() {
		var i ProductShelf
		if err := rows.Scan(&i.ID, &i.ShelfID, &i.IsPrimary); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPShelvesByID = `-- name: GetPShelvesByID :one
SELECT id, shelf_id, is_primary FROM Product_Shelves WHERE id = $1
`

func (q *Queries) GetPShelvesByID(ctx context.Context, id sql.NullInt32) (ProductShelf, error) {
	row := q.db.QueryRowContext(ctx, getPShelvesByID, id)
	var i ProductShelf
	err := row.Scan(&i.ID, &i.ShelfID, &i.IsPrimary)
	return i, err
}

const getPShelvesByIDs = `-- name: GetPShelvesByIDs :one
SELECT id, shelf_id, is_primary FROM Product_Shelves WHERE id = $1 AND shelf_id = $2
`

type GetPShelvesByIDsParams struct {
	ID      sql.NullInt32 `json:"id"`
	ShelfID sql.NullInt32 `json:"shelf_id"`
}

func (q *Queries) GetPShelvesByIDs(ctx context.Context, arg GetPShelvesByIDsParams) (ProductShelf, error) {
	row := q.db.QueryRowContext(ctx, getPShelvesByIDs, arg.ID, arg.ShelfID)
	var i ProductShelf
	err := row.Scan(&i.ID, &i.ShelfID, &i.IsPrimary)
	return i, err
}

const getPShelvesByShelfID = `-- name: GetPShelvesByShelfID :one
SELECT id, shelf_id, is_primary FROM Product_Shelves WHERE shelf_id = $1
`

func (q *Queries) GetPShelvesByShelfID(ctx context.Context, shelfID sql.NullInt32) (ProductShelf, error) {
	row := q.db.QueryRowContext(ctx, getPShelvesByShelfID, shelfID)
	var i ProductShelf
	err := row.Scan(&i.ID, &i.ShelfID, &i.IsPrimary)
	return i, err
}

const updatePShelves = `-- name: UpdatePShelves :one
UPDATE Product_Shelves
SET is_primary = false
WHERE id = $1 AND shelf_id = $2
RETURNING id, shelf_id, is_primary
`

type UpdatePShelvesParams struct {
	ID      sql.NullInt32 `json:"id"`
	ShelfID sql.NullInt32 `json:"shelf_id"`
}

func (q *Queries) UpdatePShelves(ctx context.Context, arg UpdatePShelvesParams) (ProductShelf, error) {
	row := q.db.QueryRowContext(ctx, updatePShelves, arg.ID, arg.ShelfID)
	var i ProductShelf
	err := row.Scan(&i.ID, &i.ShelfID, &i.IsPrimary)
	return i, err
}
