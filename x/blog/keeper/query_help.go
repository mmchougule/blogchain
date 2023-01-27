package keeper

import (
	"context"

	"blog/x/blog/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) HelpAll(goCtx context.Context, req *types.QueryAllHelpRequest) (*types.QueryAllHelpResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var helps []types.Help
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	helpStore := prefix.NewStore(store, types.KeyPrefix(types.HelpKey))

	pageRes, err := query.Paginate(helpStore, req.Pagination, func(key []byte, value []byte) error {
		var help types.Help
		if err := k.cdc.Unmarshal(value, &help); err != nil {
			return err
		}

		helps = append(helps, help)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllHelpResponse{Help: helps, Pagination: pageRes}, nil
}

func (k Keeper) Help(goCtx context.Context, req *types.QueryGetHelpRequest) (*types.QueryGetHelpResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	help, found := k.GetHelp(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetHelpResponse{Help: help}, nil
}
