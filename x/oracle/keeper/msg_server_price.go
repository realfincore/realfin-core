package keeper

import (
	"context"
	"fmt"

	"realfin/x/oracle/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreatePrice(goCtx context.Context, msg *types.MsgCreatePrice) (*types.MsgCreatePriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var price = types.Price{
		Creator: msg.Creator,
		Symbol:  msg.Symbol,
		Price:   msg.Price,
	}

	id := k.AppendPrice(
		ctx,
		price,
	)

	return &types.MsgCreatePriceResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdatePrice(goCtx context.Context, msg *types.MsgUpdatePrice) (*types.MsgUpdatePriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var price = types.Price{
		Creator: msg.Creator,
		Id:      msg.Id,
		Symbol:  msg.Symbol,
		Price:   msg.Price,
	}

	// Checks that the element exists
	val, found := k.GetPrice(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetPrice(ctx, price)

	return &types.MsgUpdatePriceResponse{}, nil
}

func (k msgServer) DeletePrice(goCtx context.Context, msg *types.MsgDeletePrice) (*types.MsgDeletePriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetPrice(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemovePrice(ctx, msg.Id)

	return &types.MsgDeletePriceResponse{}, nil
}
