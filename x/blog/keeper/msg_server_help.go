package keeper

import (
	"context"
	"fmt"

	"blog/x/blog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateHelp(goCtx context.Context, msg *types.MsgCreateHelp) (*types.MsgCreateHelpResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var help = types.Help{
		Creator: msg.Creator,
	}

	id := k.AppendHelp(
		ctx,
		help,
	)

	return &types.MsgCreateHelpResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateHelp(goCtx context.Context, msg *types.MsgUpdateHelp) (*types.MsgUpdateHelpResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var help = types.Help{
		Creator: msg.Creator,
		Id:      msg.Id,
	}

	// Checks that the element exists
	val, found := k.GetHelp(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetHelp(ctx, help)

	return &types.MsgUpdateHelpResponse{}, nil
}

func (k msgServer) DeleteHelp(goCtx context.Context, msg *types.MsgDeleteHelp) (*types.MsgDeleteHelpResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetHelp(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveHelp(ctx, msg.Id)

	return &types.MsgDeleteHelpResponse{}, nil
}
