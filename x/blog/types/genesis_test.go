package types_test

import (
	"testing"

	"blog/x/blog/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				HelpList: []types.Help{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				HelpCount: 2,
				PostList: []types.Post{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				PostCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated help",
			genState: &types.GenesisState{
				HelpList: []types.Help{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid help count",
			genState: &types.GenesisState{
				HelpList: []types.Help{
					{
						Id: 1,
					},
				},
				HelpCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated post",
			genState: &types.GenesisState{
				PostList: []types.Post{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid post count",
			genState: &types.GenesisState{
				PostList: []types.Post{
					{
						Id: 1,
					},
				},
				PostCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
