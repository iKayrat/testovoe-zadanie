// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: order.sql

package db

import (
	"context"
	"database/sql"
)

const createOrder = `-- name: CreateOrder :one
INSERT INTO Orders (
    order_number, 
    product_name, 
    quantity, 
    product_id, 
    additional_shelf
) VALUES (
      $1, $2, $3, $4, $5
) RETURNING order_id, order_number, product_name, quantity, product_id, additional_shelf, created_at, updated_at
`

type CreateOrderParams struct {
	OrderNumber     string         `json:"order_number"`
	ProductName     string         `json:"product_name"`
	Quantity        int32          `json:"quantity"`
	ProductID       sql.NullInt32  `json:"product_id"`
	AdditionalShelf sql.NullString `json:"additional_shelf"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, createOrder,
		arg.OrderNumber,
		arg.ProductName,
		arg.Quantity,
		arg.ProductID,
		arg.AdditionalShelf,
	)
	var i Order
	err := row.Scan(
		&i.OrderID,
		&i.OrderNumber,
		&i.ProductName,
		&i.Quantity,
		&i.ProductID,
		&i.AdditionalShelf,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteOrder = `-- name: DeleteOrder :one
DELETE FROM Orders WHERE order_id = $1
RETURNING order_id
`

func (q *Queries) DeleteOrder(ctx context.Context, orderID int32) (int32, error) {
	row := q.db.QueryRowContext(ctx, deleteOrder, orderID)
	var order_id int32
	err := row.Scan(&order_id)
	return order_id, err
}

const getOrderByGrouping = `-- name: GetOrderByGrouping :many
SELECT o.order_number,
       p.title,
       o.product_id,
       o.quantity,
       s.shelf_name AS primary_shelf,
       COALESCE(STRING_AGG(ps_false_shelves.shelf_name, ', '), '') AS additional_shelves
FROM Orders o
JOIN products p ON o.product_id = p.id
JOIN product_shelves ps ON p.id = ps.id AND is_primary = true
JOIN Shelves s ON ps.shelf_id = s.shelf_id
LEFT JOIN product_shelves ps_false ON p.id = ps_false.id AND ps_false.is_primary = false
LEFT JOIN Shelves ps_false_shelves ON ps_false.shelf_id = ps_false_shelves.shelf_id
WHERE o.order_number IN ($1)
GROUP BY o.order_number, p.title, o.product_id, o.quantity, s.shelf_name ORDER BY primary_shelf, product_id LIMIT 100
`

type GetOrderByGroupingRow struct {
	OrderNumber       string        `json:"order_number"`
	Title             string        `json:"title"`
	ProductID         sql.NullInt32 `json:"product_id"`
	Quantity          int32         `json:"quantity"`
	PrimaryShelf      string        `json:"primary_shelf"`
	AdditionalShelves interface{}   `json:"additional_shelves"`
}

func (q *Queries) GetOrderByGrouping(ctx context.Context, orderNumber string) ([]GetOrderByGroupingRow, error) {
	rows, err := q.db.QueryContext(ctx, getOrderByGrouping, orderNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetOrderByGroupingRow
	for rows.Next() {
		var i GetOrderByGroupingRow
		if err := rows.Scan(
			&i.OrderNumber,
			&i.Title,
			&i.ProductID,
			&i.Quantity,
			&i.PrimaryShelf,
			&i.AdditionalShelves,
		); err != nil {
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

const getOrderByNumber = `-- name: GetOrderByNumber :one
SELECT order_id, order_number, product_name, quantity, product_id, additional_shelf, created_at, updated_at FROM Orders WHERE order_number = $1
`

func (q *Queries) GetOrderByNumber(ctx context.Context, orderNumber string) (Order, error) {
	row := q.db.QueryRowContext(ctx, getOrderByNumber, orderNumber)
	var i Order
	err := row.Scan(
		&i.OrderID,
		&i.OrderNumber,
		&i.ProductName,
		&i.Quantity,
		&i.ProductID,
		&i.AdditionalShelf,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getOrderByProductID = `-- name: GetOrderByProductID :one
SELECT order_id, order_number, product_name, quantity, product_id, additional_shelf, created_at, updated_at FROM Orders WHERE product_id = $1
`

func (q *Queries) GetOrderByProductID(ctx context.Context, productID sql.NullInt32) (Order, error) {
	row := q.db.QueryRowContext(ctx, getOrderByProductID, productID)
	var i Order
	err := row.Scan(
		&i.OrderID,
		&i.OrderNumber,
		&i.ProductName,
		&i.Quantity,
		&i.ProductID,
		&i.AdditionalShelf,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getOrders = `-- name: GetOrders :many
SELECT order_id, order_number, product_name, quantity, product_id, additional_shelf, created_at, updated_at FROM Orders
`

func (q *Queries) GetOrders(ctx context.Context) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, getOrders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.OrderID,
			&i.OrderNumber,
			&i.ProductName,
			&i.Quantity,
			&i.ProductID,
			&i.AdditionalShelf,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
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

const updateOrder = `-- name: UpdateOrder :one
UPDATE Orders
SET quantity = $2, additional_shelf = $3
WHERE order_id = $1
RETURNING order_id, order_number, product_name, quantity, product_id, additional_shelf, created_at, updated_at
`

type UpdateOrderParams struct {
	OrderID         int32          `json:"order_id"`
	Quantity        int32          `json:"quantity"`
	AdditionalShelf sql.NullString `json:"additional_shelf"`
}

func (q *Queries) UpdateOrder(ctx context.Context, arg UpdateOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, updateOrder, arg.OrderID, arg.Quantity, arg.AdditionalShelf)
	var i Order
	err := row.Scan(
		&i.OrderID,
		&i.OrderNumber,
		&i.ProductName,
		&i.Quantity,
		&i.ProductID,
		&i.AdditionalShelf,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}