package recipes

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"recipes/testutil/sample"
	recipessimulation "recipes/x/recipes/simulation"
	"recipes/x/recipes/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = recipessimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateRecipe = "op_weight_msg_create_recipe"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateRecipe int = 100

	opWeightMsgCreateData = "op_weight_msg_create_data"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateData int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	recipesGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&recipesGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateRecipe int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateRecipe, &weightMsgCreateRecipe, nil,
		func(_ *rand.Rand) {
			weightMsgCreateRecipe = defaultWeightMsgCreateRecipe
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateRecipe,
		recipessimulation.SimulateMsgCreateRecipe(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateData int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateData, &weightMsgCreateData, nil,
		func(_ *rand.Rand) {
			weightMsgCreateData = defaultWeightMsgCreateData
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateData,
		recipessimulation.SimulateMsgCreateData(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateRecipe,
			defaultWeightMsgCreateRecipe,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				recipessimulation.SimulateMsgCreateRecipe(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateData,
			defaultWeightMsgCreateData,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				recipessimulation.SimulateMsgCreateData(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
