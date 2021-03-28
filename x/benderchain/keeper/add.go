package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/faddat/bender-chain/x/benderchain/types"
	"strconv"
)

// GetAddCount get the total number of add
func (k Keeper) GetAddCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddCountKey))
	byteKey := types.KeyPrefix(types.AddCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to iint64
		panic("cannot decode count")
	}

	return count
}

// SetAddCount set the total number of add
func (k Keeper) SetAddCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddCountKey))
	byteKey := types.KeyPrefix(types.AddCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendAdd appends a add in the store with a new id and update the count
func (k Keeper) AppendAdd(
	ctx sdk.Context,
	creator string,
	post string,
	title string,
	body string,
	ipfs string,
) uint64 {
	// Create the add
	count := k.GetAddCount(ctx)
	var add = types.Add{
		Creator: creator,
		Id:      count,
		Post:    post,
		Title:   title,
		Body:    body,
		Ipfs:    ipfs,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&add)
	store.Set(GetAddIDBytes(add.Id), appendedValue)

	// Update add count
	k.SetAddCount(ctx, count+1)

	return count
}

// SetAdd set a specific add in the store
func (k Keeper) SetAdd(ctx sdk.Context, add types.Add) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddKey))
	b := k.cdc.MustMarshalBinaryBare(&add)
	store.Set(GetAddIDBytes(add.Id), b)
}

// GetAdd returns a add from its id
func (k Keeper) GetAdd(ctx sdk.Context, id uint64) types.Add {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddKey))
	var add types.Add
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetAddIDBytes(id)), &add)
	return add
}

// HasAdd checks if the add exists in the store
func (k Keeper) HasAdd(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddKey))
	return store.Has(GetAddIDBytes(id))
}

// GetAddOwner returns the creator of the add
func (k Keeper) GetAddOwner(ctx sdk.Context, id uint64) string {
	return k.GetAdd(ctx, id).Creator
}

// RemoveAdd removes a add from the store
func (k Keeper) RemoveAdd(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddKey))
	store.Delete(GetAddIDBytes(id))
}

// GetAllAdd returns all add
func (k Keeper) GetAllAdd(ctx sdk.Context) (list []types.Add) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Add
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAddIDBytes returns the byte representation of the ID
func GetAddIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetAddIDFromBytes returns ID in uint64 format from a byte array
func GetAddIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
