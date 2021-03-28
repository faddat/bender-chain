package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateAdd{}

func NewMsgCreateAdd(creator string, post string, title string, body string, ipfs string) *MsgCreateAdd {
	return &MsgCreateAdd{
		Creator: creator,
		Post:    post,
		Title:   title,
		Body:    body,
		Ipfs:    ipfs,
	}
}

func (msg *MsgCreateAdd) Route() string {
	return RouterKey
}

func (msg *MsgCreateAdd) Type() string {
	return "CreateAdd"
}

func (msg *MsgCreateAdd) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateAdd) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAdd) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateAdd{}

func NewMsgUpdateAdd(creator string, id uint64, post string, title string, body string, ipfs string) *MsgUpdateAdd {
	return &MsgUpdateAdd{
		Id:      id,
		Creator: creator,
		Post:    post,
		Title:   title,
		Body:    body,
		Ipfs:    ipfs,
	}
}

func (msg *MsgUpdateAdd) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAdd) Type() string {
	return "UpdateAdd"
}

func (msg *MsgUpdateAdd) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAdd) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAdd) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteAdd{}

func NewMsgDeleteAdd(creator string, id uint64) *MsgDeleteAdd {
	return &MsgDeleteAdd{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteAdd) Route() string {
	return RouterKey
}

func (msg *MsgDeleteAdd) Type() string {
	return "DeleteAdd"
}

func (msg *MsgDeleteAdd) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteAdd) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteAdd) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
