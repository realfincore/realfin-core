package keeper_test

import (
	"context"
	"testing"

	keepertest "realfin/testutil/keeper"
	"realfin/testutil/nullify"
	"realfin/x/oracle/keeper"
	"realfin/x/oracle/types"

	"github.com/stretchr/testify/require"
)

func createNPrice(keeper keeper.Keeper, ctx context.Context, n int) []types.Price {
	items := make([]types.Price, n)
	for i := range items {
		items[i].Id = keeper.AppendPrice(ctx, items[i])
	}
	return items
}

func TestPriceGet(t *testing.T) {
	keeper, ctx := keepertest.OracleKeeper(t)
	items := createNPrice(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetPrice(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestPriceRemove(t *testing.T) {
	keeper, ctx := keepertest.OracleKeeper(t)
	items := createNPrice(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePrice(ctx, item.Id)
		_, found := keeper.GetPrice(ctx, item.Id)
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

func TestPriceCount(t *testing.T) {
	keeper, ctx := keepertest.OracleKeeper(t)
	items := createNPrice(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetPriceCount(ctx))
}
