package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePrice{}

func NewMsgCreatePrice(
	creator string,
	symbol string,
	price uint64,
	timestamp int32,

) *MsgCreatePrice {
	return &MsgCreatePrice{
		Creator:   creator,
		Symbol:    symbol,
		Price:     price,
		Timestamp: timestamp,
	}
}

func (msg *MsgCreatePrice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdatePrice{}

func NewMsgUpdatePrice(
	creator string,
	symbol string,
	price uint64,
	timestamp int32,

) *MsgUpdatePrice {
	return &MsgUpdatePrice{
		Creator:   creator,
		Symbol:    symbol,
		Price:     price,
		Timestamp: timestamp,
	}
}

func (msg *MsgUpdatePrice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeletePrice{}

func NewMsgDeletePrice(
	creator string,
	symbol string,

) *MsgDeletePrice {
	return &MsgDeletePrice{
		Creator: creator,
		Symbol:  symbol,
	}
}

func (msg *MsgDeletePrice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
