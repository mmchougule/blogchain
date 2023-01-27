package keeper_test

import (
	"testing"

	keepertest "blog/testutil/keeper"
	"blog/testutil/nullify"
	"blog/x/blog/keeper"
	"blog/x/blog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNHelp(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Help {
	items := make([]types.Help, n)
	for i := range items {
		items[i].Id = keeper.AppendHelp(ctx, items[i])
	}
	return items
}

func TestHelpGet(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNHelp(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetHelp(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestHelpRemove(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNHelp(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveHelp(ctx, item.Id)
		_, found := keeper.GetHelp(ctx, item.Id)
		require.False(t, found)
	}
}

func TestHelpGetAll(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNHelp(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllHelp(ctx)),
	)
}

func TestHelpCount(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNHelp(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetHelpCount(ctx))
}
