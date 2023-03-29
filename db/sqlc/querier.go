// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
)

type Querier interface {
	AddInventoryAmount(ctx context.Context, arg AddInventoryAmountParams) (Inventory, error)
	AddProductAmount(ctx context.Context, arg AddProductAmountParams) (Product, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	CreateInventory(ctx context.Context, arg CreateInventoryParams) (Inventory, error)
	CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error)
	CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetEntry(ctx context.Context, id int64) (Entry, error)
	GetInventory(ctx context.Context, id int64) (Inventory, error)
	GetInventoryForUpdate(ctx context.Context, id int64) (Inventory, error)
	GetOrder(ctx context.Context, id int64) (Order, error)
	GetProduct(ctx context.Context, id int64) (Product, error)
	GetProductForUpdate(ctx context.Context, id int64) (Product, error)
	GetUser(ctx context.Context, username string) (User, error)
	GetUserForUpdate(ctx context.Context, username string) (User, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
	ListInventory(ctx context.Context, arg ListInventoryParams) ([]Inventory, error)
	ListOrder(ctx context.Context, arg ListOrderParams) ([]Order, error)
	ListProduct(ctx context.Context, arg ListProductParams) ([]Product, error)
	UpdateInventory(ctx context.Context, arg UpdateInventoryParams) (Inventory, error)
	UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error)
	UpdateUserHashedPassword(ctx context.Context, arg UpdateUserHashedPasswordParams) (User, error)
	UpdateUserRole(ctx context.Context, arg UpdateUserRoleParams) (User, error)
}

var _ Querier = (*Queries)(nil)