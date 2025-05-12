package oracle_test

import (
	"testing"

	keepertest "realfin/testutil/keeper"
	"realfin/testutil/nullify"
	oracle "realfin/x/oracle/module"
	"realfin/x/oracle/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		PriceList: []types.Price{
			{
				Symbol: "0",
			},
			{
				Symbol: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OracleKeeper(t)
	oracle.InitGenesis(ctx, k, genesisState)
	got := oracle.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PriceList, got.PriceList)
	// this line is used by starport scaffolding # genesis/test/assert
}
