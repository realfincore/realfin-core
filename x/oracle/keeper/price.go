package keeper

import (
	"context"
	"encoding/binary"

	"realfin/x/oracle/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetPriceCount get the total number of price
func (k Keeper) GetPriceCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.PriceCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetPriceCount set the total number of price
func (k Keeper) SetPriceCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.PriceCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendPrice appends a price in the store with a new id and update the count
func (k Keeper) AppendPrice(
	ctx context.Context,
	price types.Price,
) uint64 {
	// Create the price
	count := k.GetPriceCount(ctx)

	// Set the ID of the appended value
	price.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PriceKey))
	appendedValue := k.cdc.MustMarshal(&price)
	store.Set(GetPriceIDBytes(price.Id), appendedValue)

	// Update price count
	k.SetPriceCount(ctx, count+1)

	return count
}

// SetPrice set a specific price in the store
func (k Keeper) SetPrice(ctx context.Context, price types.Price) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PriceKey))
	b := k.cdc.MustMarshal(&price)
	store.Set(GetPriceIDBytes(price.Id), b)
}

// GetPrice returns a price from its id
func (k Keeper) GetPrice(ctx context.Context, id uint64) (val types.Price, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PriceKey))
	b := store.Get(GetPriceIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePrice removes a price from the store
func (k Keeper) RemovePrice(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PriceKey))
	store.Delete(GetPriceIDBytes(id))
}

// GetAllPrice returns all price
func (k Keeper) GetAllPrice(ctx context.Context) (list []types.Price) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PriceKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Price
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetPriceIDBytes returns the byte representation of the ID
func GetPriceIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.PriceKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
