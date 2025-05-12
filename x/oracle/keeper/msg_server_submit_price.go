package keeper

import (
	"context"

	"realfin/x/oracle/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SubmitPrice(goCtx context.Context, msg *types.MsgSubmitPrice) (*types.MsgSubmitPriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate price
	if msg.Price <= 0 {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "price must be positive")
	}

	// Store the price (you may need to implement actual storage logic)
	price := types.Price{
		Symbol:  msg.Symbol,
		Price:   msg.Price,
		Creator: msg.Creator,
	}
	k.Keeper.SetPrice(ctx, price)

	return &types.MsgSubmitPriceResponse{}, nil
}
