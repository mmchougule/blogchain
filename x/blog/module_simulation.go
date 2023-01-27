package blog

import (
	"math/rand"

	"blog/testutil/sample"
	blogsimulation "blog/x/blog/simulation"
	"blog/x/blog/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = blogsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateHelp = "op_weight_msg_help"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateHelp int = 100

	opWeightMsgUpdateHelp = "op_weight_msg_help"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateHelp int = 100

	opWeightMsgDeleteHelp = "op_weight_msg_help"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteHelp int = 100

	opWeightMsgCreatePost = "op_weight_msg_post"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePost int = 100

	opWeightMsgUpdatePost = "op_weight_msg_post"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdatePost int = 100

	opWeightMsgDeletePost = "op_weight_msg_post"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeletePost int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	blogGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		HelpList: []types.Help{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		HelpCount: 2,
		PostList: []types.Post{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		PostCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&blogGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateHelp int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateHelp, &weightMsgCreateHelp, nil,
		func(_ *rand.Rand) {
			weightMsgCreateHelp = defaultWeightMsgCreateHelp
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateHelp,
		blogsimulation.SimulateMsgCreateHelp(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateHelp int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateHelp, &weightMsgUpdateHelp, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateHelp = defaultWeightMsgUpdateHelp
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateHelp,
		blogsimulation.SimulateMsgUpdateHelp(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteHelp int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteHelp, &weightMsgDeleteHelp, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteHelp = defaultWeightMsgDeleteHelp
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteHelp,
		blogsimulation.SimulateMsgDeleteHelp(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreatePost int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreatePost, &weightMsgCreatePost, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePost = defaultWeightMsgCreatePost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePost,
		blogsimulation.SimulateMsgCreatePost(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdatePost int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdatePost, &weightMsgUpdatePost, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePost = defaultWeightMsgUpdatePost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePost,
		blogsimulation.SimulateMsgUpdatePost(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeletePost int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeletePost, &weightMsgDeletePost, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePost = defaultWeightMsgDeletePost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeletePost,
		blogsimulation.SimulateMsgDeletePost(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
