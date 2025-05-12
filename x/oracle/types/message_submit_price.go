package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSubmitPrice{}

func NewMsgSubmitPrice(creator string, symbol string, price uint64) *MsgSubmitPrice {
	return &MsgSubmitPrice{
		Creator: creator,
		Symbol:  symbol,
		Price:   price,
	}
}

func (msg *MsgSubmitPrice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
