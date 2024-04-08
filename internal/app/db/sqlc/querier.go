package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error)
	DeleteProduct(ctx context.Context, id int64) (int64, error)
	GetProduct(ctx context.Context, id int64) (Product, error)
	ImportProducts(ctx context.Context, arg ImportProductsParams) error
	ListProducts(ctx context.Context, arg ListProductsParams) ([]Product, error)
	UpdateProduct(ctx context.Context, arg UpdateProductParams) error

	CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error)
	DeleteOrder(ctx context.Context, orderID int32) (int32, error)
	GetOrderByGrouping(ctx context.Context, orderNumber string) ([]GetOrderByGroupingRow, error)

	GetOrderByNumber(ctx context.Context, orderNumber string) (Order, error)
	GetOrderByProductID(ctx context.Context, productID sql.NullInt32) (Order, error)
	GetOrders(ctx context.Context) ([]Order, error)
	UpdateOrder(ctx context.Context, arg UpdateOrderParams) error

	CreatePSHelves(ctx context.Context, arg CreatePSHelvesParams) (ProductShelf, error)
	DeletePShelves(ctx context.Context, id sql.NullInt32) (sql.NullInt32, error)
	GetPShelves(ctx context.Context) ([]ProductShelf, error)
	GetPShelvesByID(ctx context.Context, id sql.NullInt32) (ProductShelf, error)
	GetPShelvesByIDs(ctx context.Context, arg GetPShelvesByIDsParams) (ProductShelf, error)
	GetPShelvesByShelfID(ctx context.Context, shelfID sql.NullInt32) (ProductShelf, error)
	UpdatePShelves(ctx context.Context, arg UpdatePShelvesParams) error

	CreateShelf(ctx context.Context, shelfName string) (int32, error)
	DeleteShelf(ctx context.Context, shelfID int32) (int32, error)
	GetShelfByID(ctx context.Context) (Shelf, error)
	GetShelfByName(ctx context.Context, shelfName string) (Shelf, error)
	GetShelfs(ctx context.Context) ([]Shelf, error)
	UpdateShelf(ctx context.Context, arg UpdateShelfParams) error
}

// var _ Querier = (*Queries)(nil)
