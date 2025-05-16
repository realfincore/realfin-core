package creditscore

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"realfin/x/creditscore/keeper"
	"realfin/x/creditscore/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the rate
	for _, elem := range genState.RateList {
		k.SetRate(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.RateList = k.GetAllRate(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
