package creditscore_test

import (
	"testing"

	keepertest "realfin/testutil/keeper"
	"realfin/testutil/nullify"
	creditscore "realfin/x/creditscore/module"
	"realfin/x/creditscore/types"

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

	k, ctx := keepertest.CreditscoreKeeper(t)
	creditscore.InitGenesis(ctx, k, genesisState)
	got := creditscore.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.RateList, got.RateList)
	// this line is used by starport scaffolding # genesis/test/assert
}
