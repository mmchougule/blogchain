package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateHelp = "create_help"
	TypeMsgUpdateHelp = "update_help"
	TypeMsgDeleteHelp = "delete_help"
)

var _ sdk.Msg = &MsgCreateHelp{}

func NewMsgCreateHelp(creator string) *MsgCreateHelp {
	return &MsgCreateHelp{
		Creator: creator,
	}
}

func (msg *MsgCreateHelp) Route() string {
	return RouterKey
}

func (msg *MsgCreateHelp) Type() string {
	return TypeMsgCreateHelp
}

func (msg *MsgCreateHelp) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateHelp) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateHelp) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateHelp{}

func NewMsgUpdateHelp(creator string, id uint64) *MsgUpdateHelp {
	return &MsgUpdateHelp{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgUpdateHelp) Route() string {
	return RouterKey
}

func (msg *MsgUpdateHelp) Type() string {
	return TypeMsgUpdateHelp
}

func (msg *MsgUpdateHelp) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateHelp) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateHelp) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteHelp{}

func NewMsgDeleteHelp(creator string, id uint64) *MsgDeleteHelp {
	return &MsgDeleteHelp{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteHelp) Route() string {
	return RouterKey
}

func (msg *MsgDeleteHelp) Type() string {
	return TypeMsgDeleteHelp
}

func (msg *MsgDeleteHelp) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteHelp) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteHelp) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
