package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "realfin/testutil/keeper"
	"realfin/x/creditscore/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.CreditscoreKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
