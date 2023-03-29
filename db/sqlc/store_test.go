package db

import (
	"context"
	"example/shopmanagement/util"
	"fmt"
	"testing"
)

func TestCreateOrder(t *testing.T) {
	store := NewStore(testDB)
	product := createRandomProduct(t)

	fmt.Println(">> Before: ", product.Amount)

	n := 5
	amount := int64(1)
	unit := "g"

	errs := make(chan error)
	results := make(chan CreateOrderTXResult)

	for i := 0; i < n; i++ {
		go func() {
			ctx := context.Background()
			result, err := store.CreateOrderTX(ctx, CreateOrderTXParams{
				BuyerName: util.RandomString(4),
				ProductID: product.ID,
				Amount:    amount,
				Unit:      unit,
				Address:   util.RandomString(6),
			})
			errs <- err
			results <- result
		}()
	}
}
