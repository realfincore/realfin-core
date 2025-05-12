package keeper

import (
	"context"

	"realfin/x/oracle/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreatePrice(goCtx context.Context, msg *types.MsgCreatePrice) (*types.MsgCreatePriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetPrice(
		ctx,
		msg.Symbol,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var price = types.Price{
		Creator:   msg.Creator,
		Symbol:    msg.Symbol,
		Price:     msg.Price,
		Timestamp: msg.Timestamp,
	}

	k.SetPrice(
		ctx,
		price,
	)
	return &types.MsgCreatePriceResponse{}, nil
}

func (k msgServer) UpdatePrice(goCtx context.Context, msg *types.MsgUpdatePrice) (*types.MsgUpdatePriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetPrice(
		ctx,
		msg.Symbol,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var price = types.Price{
		Creator:   msg.Creator,
		Symbol:    msg.Symbol,
		Price:     msg.Price,
		Timestamp: msg.Timestamp,
	}

	k.SetPrice(ctx, price)

	return &types.MsgUpdatePriceResponse{}, nil
}

func (k msgServer) DeletePrice(goCtx context.Context, msg *types.MsgDeletePrice) (*types.MsgDeletePriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetPrice(
		ctx,
		msg.Symbol,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemovePrice(
		ctx,
		msg.Symbol,
	)

	return &types.MsgDeletePriceResponse{}, nil
}
