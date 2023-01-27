package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"blog/x/blog/types"
)

func TestHelpMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateHelp(ctx, &types.MsgCreateHelp{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestHelpMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateHelp
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateHelp{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateHelp{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateHelp{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateHelp(ctx, &types.MsgCreateHelp{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateHelp(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestHelpMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteHelp
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteHelp{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteHelp{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteHelp{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateHelp(ctx, &types.MsgCreateHelp{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteHelp(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
