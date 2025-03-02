package realfin_test

import (
	"testing"

	keepertest "realfin/testutil/keeper"
	"realfin/testutil/nullify"
	realfin "realfin/x/realfin/module"
	"realfin/x/realfin/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RealfinKeeper(t)
	realfin.InitGenesis(ctx, k, genesisState)
	got := realfin.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
