package db

import (
	"context"
	"example/shopmanagement/util"
	"fmt"
	"testing"

	"github.com/test-go/testify/require"
)

func TestCreateOrder(t *testing.T) {
	store := NewStore(testDB)
	product := createRandomProduct(t)

	fmt.Println(">> Before: ", product.Amount)

	n := 5
	amount := int64(1)
	price := uint64(10000)
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
				Price:     price,
			})
			errs <- err
			results <- result
		}()
	}
	// existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check Buyer
		buyer := result.Buyer
		require.NotEmpty(t, buyer)
		require.NotEmpty(t, buyer.ID)
		require.NotEmpty(t, buyer.Name)
		require.NotZero(t, buyer.ID)
		require.NotZero(t, buyer.CreatedAt)

		// check Order
		order := result.Order
		require.NotEmpty(t, order)
		require.Equal(t, order.BuyerID, buyer.ID)
		require.Equal(t, order.Amount, amount)
		require.Equal(t, product.ID, order.ProductID)
		require.Equal(t, unit, order.Unit)
		require.NotZero(t, order.ID)
		require.NotZero(t, order.CreatedAt)

		// check Product
		product_result := result.Product
		require.NotEmpty(t, product_result)
		require.Equal(t, product.Amount-int64((i+1))*amount, product_result.Amount)
		require.Equal(t, product.Name, product_result.Name)
		require.Equal(t, product.Type, product_result.Type)
		require.Equal(t, product.ID, product_result.ID)
		require.Equal(t, product.CreatedAt, product_result.CreatedAt)

		// check price
		require.Equal(t, price, result.Price)
	}
}
