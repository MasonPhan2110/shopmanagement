package db

import (
	"context"
	"example/shopmanagement/util"
	"testing"

	"github.com/test-go/testify/require"
)

func createRandomProduct(t *testing.T) Product {
	arg := CreateProductParams{
		Type:   util.RandomString(6),
		Name:   util.RandomString(6),
		Amount: util.RandomInt(10, 100),
	}
	product, err := testQueries.CreateProduct(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, arg.Type, product.Type)
	require.Equal(t, arg.Name, product.Name)
	require.Equal(t, arg.Amount, product.Amount)

	require.NotZero(t, product.ID)
	require.NotZero(t, product.CreatedAt)
	return product
}
