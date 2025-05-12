package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "realfin/testutil/keeper"
	"realfin/testutil/nullify"
	"realfin/x/oracle/keeper"
	"realfin/x/oracle/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNPrice(keeper keeper.Keeper, ctx context.Context, n int) []types.Price {
	items := make([]types.Price, n)
	for i := range items {
		items[i].Symbol = strconv.Itoa(i)

		keeper.SetPrice(ctx, items[i])
	}
	return items
}

func TestPriceGet(t *testing.T) {
	keeper, ctx := keepertest.OracleKeeper(t)
	items := createNPrice(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPrice(ctx,
			item.Symbol,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPriceRemove(t *testing.T) {
	keeper, ctx := keepertest.OracleKeeper(t)
	items := createNPrice(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePrice(ctx,
			item.Symbol,
		)
		_, found := keeper.GetPrice(ctx,
			item.Symbol,
		)
		require.False(t, found)
	}
}

func TestPriceGetAll(t *testing.T) {
	keeper, ctx := keepertest.OracleKeeper(t)
	items := createNPrice(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPrice(ctx)),
	)
}
