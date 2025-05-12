package oracle

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"realfin/testutil/sample"
	oraclesimulation "realfin/x/oracle/simulation"
	"realfin/x/oracle/types"
)

// avoid unused import issue
var (
	_ = oraclesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreatePrice = "op_weight_msg_price"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePrice int = 100

	opWeightMsgUpdatePrice = "op_weight_msg_price"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdatePrice int = 100

	opWeightMsgDeletePrice = "op_weight_msg_price"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeletePrice int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	oracleGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PriceList: []types.Price{
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
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&oracleGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreatePrice int
	simState.AppParams.GetOrGenerate(opWeightMsgCreatePrice, &weightMsgCreatePrice, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePrice = defaultWeightMsgCreatePrice
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePrice,
		oraclesimulation.SimulateMsgCreatePrice(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdatePrice int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdatePrice, &weightMsgUpdatePrice, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePrice = defaultWeightMsgUpdatePrice
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePrice,
		oraclesimulation.SimulateMsgUpdatePrice(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeletePrice int
	simState.AppParams.GetOrGenerate(opWeightMsgDeletePrice, &weightMsgDeletePrice, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePrice = defaultWeightMsgDeletePrice
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeletePrice,
		oraclesimulation.SimulateMsgDeletePrice(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreatePrice,
			defaultWeightMsgCreatePrice,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oraclesimulation.SimulateMsgCreatePrice(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdatePrice,
			defaultWeightMsgUpdatePrice,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oraclesimulation.SimulateMsgUpdatePrice(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeletePrice,
			defaultWeightMsgDeletePrice,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oraclesimulation.SimulateMsgDeletePrice(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
