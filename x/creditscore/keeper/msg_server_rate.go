package keeper

import (
	"context"

	"realfin/x/creditscore/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateRate(goCtx context.Context, msg *types.MsgCreateRate) (*types.MsgCreateRateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetRate(
		ctx,
		msg.Symbol,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var rate = types.Rate{
		Creator:     msg.Creator,
		Symbol:      msg.Symbol,
		Rate:        msg.Rate,
		Name:        msg.Name,
		Description: msg.Description,
	}

	k.SetRate(
		ctx,
		rate,
	)
	return &types.MsgCreateRateResponse{}, nil
}

func (k msgServer) UpdateRate(goCtx context.Context, msg *types.MsgUpdateRate) (*types.MsgUpdateRateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetRate(
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

	var rate = types.Rate{
		Creator:     msg.Creator,
		Symbol:      msg.Symbol,
		Rate:        msg.Rate,
		Name:        msg.Name,
		Description: msg.Description,
	}

	k.SetRate(ctx, rate)

	return &types.MsgUpdateRateResponse{}, nil
}

func (k msgServer) DeleteRate(goCtx context.Context, msg *types.MsgDeleteRate) (*types.MsgDeleteRateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetRate(
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

	k.RemoveRate(
		ctx,
		msg.Symbol,
	)

	return &types.MsgDeleteRateResponse{}, nil
}
