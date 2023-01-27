package keeper

import (
	"encoding/binary"

	"blog/x/blog/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetHelpCount get the total number of help
func (k Keeper) GetHelpCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.HelpCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetHelpCount set the total number of help
func (k Keeper) SetHelpCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.HelpCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendHelp appends a help in the store with a new id and update the count
func (k Keeper) AppendHelp(
	ctx sdk.Context,
	help types.Help,
) uint64 {
	// Create the help
	count := k.GetHelpCount(ctx)

	// Set the ID of the appended value
	help.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HelpKey))
	appendedValue := k.cdc.MustMarshal(&help)
	store.Set(GetHelpIDBytes(help.Id), appendedValue)

	// Update help count
	k.SetHelpCount(ctx, count+1)

	return count
}

// SetHelp set a specific help in the store
func (k Keeper) SetHelp(ctx sdk.Context, help types.Help) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HelpKey))
	b := k.cdc.MustMarshal(&help)
	store.Set(GetHelpIDBytes(help.Id), b)
}

// GetHelp returns a help from its id
func (k Keeper) GetHelp(ctx sdk.Context, id uint64) (val types.Help, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HelpKey))
	b := store.Get(GetHelpIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveHelp removes a help from the store
func (k Keeper) RemoveHelp(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HelpKey))
	store.Delete(GetHelpIDBytes(id))
}

// GetAllHelp returns all help
func (k Keeper) GetAllHelp(ctx sdk.Context) (list []types.Help) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HelpKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Help
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetHelpIDBytes returns the byte representation of the ID
func GetHelpIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetHelpIDFromBytes returns ID in uint64 format from a byte array
func GetHelpIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
