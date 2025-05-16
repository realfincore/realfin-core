package realestate

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"realfin/testutil/sample"
	realestatesimulation "realfin/x/realestate/simulation"
	"realfin/x/realestate/types"
)

// avoid unused import issue
var (
	_ = realestatesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateRate = "op_weight_msg_rate"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateRate int = 100

	opWeightMsgUpdateRate = "op_weight_msg_rate"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateRate int = 100

	opWeightMsgDeleteRate = "op_weight_msg_rate"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteRate int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	realestateGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		RateList: []types.Rate{
			{
				Creator: sample.AccAddress(),
				Symbol:  "0",
			},
			{
				Creator: sample.AccAddress(),
				Symbol:  "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&realestateGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateRate int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateRate, &weightMsgCreateRate, nil,
		func(_ *rand.Rand) {
			weightMsgCreateRate = defaultWeightMsgCreateRate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateRate,
		realestatesimulation.SimulateMsgCreateRate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateRate int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateRate, &weightMsgUpdateRate, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateRate = defaultWeightMsgUpdateRate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateRate,
		realestatesimulation.SimulateMsgUpdateRate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteRate int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteRate, &weightMsgDeleteRate, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteRate = defaultWeightMsgDeleteRate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteRate,
		realestatesimulation.SimulateMsgDeleteRate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateRate,
			defaultWeightMsgCreateRate,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				realestatesimulation.SimulateMsgCreateRate(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateRate,
			defaultWeightMsgUpdateRate,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				realestatesimulation.SimulateMsgUpdateRate(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteRate,
			defaultWeightMsgDeleteRate,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				realestatesimulation.SimulateMsgDeleteRate(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
