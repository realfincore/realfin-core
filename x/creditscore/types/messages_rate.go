package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateRate{}

func NewMsgCreateRate(
	creator string,
	symbol string,
	rate uint64,
	name string,
	description string,

) *MsgCreateRate {
	return &MsgCreateRate{
		Creator:     creator,
		Symbol:      symbol,
		Rate:        rate,
		Name:        name,
		Description: description,
	}
}

func (msg *MsgCreateRate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateRate{}

func NewMsgUpdateRate(
	creator string,
	symbol string,
	rate uint64,
	name string,
	description string,

) *MsgUpdateRate {
	return &MsgUpdateRate{
		Creator:     creator,
		Symbol:      symbol,
		Rate:        rate,
		Name:        name,
		Description: description,
	}
}

func (msg *MsgUpdateRate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteRate{}

func NewMsgDeleteRate(
	creator string,
	symbol string,

) *MsgDeleteRate {
	return &MsgDeleteRate{
		Creator: creator,
		Symbol:  symbol,
	}
}

func (msg *MsgDeleteRate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
