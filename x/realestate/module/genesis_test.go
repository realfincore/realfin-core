package realestate_test

import (
	"testing"

	keepertest "realfin/testutil/keeper"
	"realfin/testutil/nullify"
	realestate "realfin/x/realestate/module"
	"realfin/x/realestate/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		RateList: []types.Rate{
			{
				Symbol: "0",
			},
			{
				Symbol: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RealestateKeeper(t)
	realestate.InitGenesis(ctx, k, genesisState)
	got := realestate.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.RateList, got.RateList)
	// this line is used by starport scaffolding # genesis/test/assert
}
