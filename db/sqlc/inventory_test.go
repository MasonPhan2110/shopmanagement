package db

import (
	"context"
	"example/shopmanagement/util"
	"testing"

	"github.com/test-go/testify/require"
)

func createRandomItem(t *testing.T) Inventory {
	arg := CreateInventoryParams{
		Type:   util.RandomString(6),
		Name:   util.RandomString(6),
		Amount: util.RandomInt(10, 100),
	}
	item, err := testQueries.CreateInventory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, item)

	require.Equal(t, arg.Type, item.Type)
	require.Equal(t, arg.Name, item.Name)
	require.Equal(t, arg.Amount, item.Amount)

	require.NotZero(t, item.ID)
	require.NotZero(t, item.CreatedAt)
	return item
}
