package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "realfin/testutil/keeper"
	"realfin/testutil/nullify"
	"realfin/x/creditscore/keeper"
	"realfin/x/creditscore/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNRate(keeper keeper.Keeper, ctx context.Context, n int) []types.Rate {
	items := make([]types.Rate, n)
	for i := range items {
		items[i].Symbol = strconv.Itoa(i)

		keeper.SetRate(ctx, items[i])
	}
	return items
}

func TestRateGet(t *testing.T) {
	keeper, ctx := keepertest.CreditscoreKeeper(t)
	items := createNRate(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetRate(ctx,
			item.Symbol,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestRateRemove(t *testing.T) {
	keeper, ctx := keepertest.CreditscoreKeeper(t)
	items := createNRate(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRate(ctx,
			item.Symbol,
		)
		_, found := keeper.GetRate(ctx,
			item.Symbol,
		)
		require.False(t, found)
	}
}

func TestRateGetAll(t *testing.T) {
	keeper, ctx := keepertest.CreditscoreKeeper(t)
	items := createNRate(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRate(ctx)),
	)
}
