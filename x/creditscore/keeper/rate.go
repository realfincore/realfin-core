package keeper

import (
	"context"

	"realfin/x/creditscore/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetRate set a specific rate in the store from its index
func (k Keeper) SetRate(ctx context.Context, rate types.Rate) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RateKeyPrefix))
	b := k.cdc.MustMarshal(&rate)
	store.Set(types.RateKey(
		rate.Symbol,
	), b)
}

// GetRate returns a rate from its index
func (k Keeper) GetRate(
	ctx context.Context,
	symbol string,

) (val types.Rate, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RateKeyPrefix))

	b := store.Get(types.RateKey(
		symbol,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRate removes a rate from the store
func (k Keeper) RemoveRate(
	ctx context.Context,
	symbol string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RateKeyPrefix))
	store.Delete(types.RateKey(
		symbol,
	))
}

// GetAllRate returns all rate
func (k Keeper) GetAllRate(ctx context.Context) (list []types.Rate) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RateKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Rate
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
