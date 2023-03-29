package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	PrepareOrderTX(ctx context.Context, arg PrepareOrderParams) (PrepareOrderResult, error)
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// execTx excutes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

func (store *SQLStore) CreateOrderTX(ctx context.Context, arg CreateOrderTXParams) (CreateOrderTXResult, error) {
	var result CreateOrderTXResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Buyer, err = q.GetBuyer(ctx, arg.BuyerName)
		if err != nil {
			return err
		}

		result.Order, err = q.CreateOrder(ctx, CreateOrderParams{
			BuyerID:   result.Buyer.ID,
			ProductID: arg.ProductID,
			Amount:    arg.Amount,
			Unit:      arg.Unit,
		})
		if err != nil {
			return err
		}

		result.Product, err = addAmountProduct(ctx, q, result.Order.ProductID, result.Order.Amount, result.Order.Unit)
		if err != nil {
			return err
		}
		return nil
	})

	return result, err
}

func (store *SQLStore) PrepareOrderTX(ctx context.Context, arg PrepareOrderParams) (PrepareOrderResult, error) {
	var result PrepareOrderResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Order, err = q.GetOrder(ctx, arg.OrderID)

		if err != nil {
			return err
		}

		result.Items, err = addAmountInventory(ctx, q, arg.ItemIDs, arg.ItemAmount)
		if err != nil {
			return err
		}
		return nil
	})

	return result, err
}

func addAmountInventory(
	ctx context.Context,
	q *Queries,
	itemIDs []int64,
	itemAmount []int64,
) (items []Inventory, err error) {
	for i := 0; i < len(itemIDs); i++ {
		item, err := q.AddInventoryAmount(ctx, AddInventoryAmountParams{
			ID:     itemIDs[i],
			Amount: itemAmount[i],
		})
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return
}

func addAmountProduct(
	ctx context.Context,
	q *Queries,
	productID int64,
	amount int64,
	unit string,
) (product Product, err error) {
	amount = toGam(amount, unit)
	product, err = q.AddProductAmount(ctx, AddProductAmountParams{
		Amount: amount,
		ID:     productID,
	})
	return
}

func toGam(amount int64, unit string) int64 {
	switch unit {
	case "g":
		return amount
	case "kg":
		return amount * 1000
	default:
		return 0
	}
}

type CreateOrderTXParams struct {
	BuyerName string `json:"buyer_id"`
	ProductID int64  `json:"product_id"`
	Amount    int64  `json:"amount"`
	Unit      string `json:"unit"`
}
type CreateOrderTXResult struct {
	Order   Order   `json:"order"`
	Buyer   Buyer   `json:"buyer"`
	Product Product `json:"product"`
}

type PrepareOrderParams struct {
	OrderID    int64   `json:"order_id"`
	ItemIDs    []int64 `json:"list_item_id"`
	ItemAmount []int64 `json:"list_item_amount"`
}
type PrepareOrderResult struct {
	Order   Order       `json:"order"`
	Product Product     `json:"product"`
	Items   []Inventory `json:"items"`
}
