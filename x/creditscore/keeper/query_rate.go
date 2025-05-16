package keeper

import (
	"context"

	"realfin/x/creditscore/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RateAll(ctx context.Context, req *types.QueryAllRateRequest) (*types.QueryAllRateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var rates []types.Rate

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	rateStore := prefix.NewStore(store, types.KeyPrefix(types.RateKeyPrefix))

	pageRes, err := query.Paginate(rateStore, req.Pagination, func(key []byte, value []byte) error {
		var rate types.Rate
		if err := k.cdc.Unmarshal(value, &rate); err != nil {
			return err
		}

		rates = append(rates, rate)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRateResponse{Rate: rates, Pagination: pageRes}, nil
}

func (k Keeper) Rate(ctx context.Context, req *types.QueryGetRateRequest) (*types.QueryGetRateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetRate(
		ctx,
		req.Symbol,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetRateResponse{Rate: val}, nil
}
