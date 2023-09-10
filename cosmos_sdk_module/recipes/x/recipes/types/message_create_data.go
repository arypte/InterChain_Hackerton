package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateData = "create_data"

var _ sdk.Msg = &MsgCreateData{}

func NewMsgCreateData(creator string, data string, meta string) *MsgCreateData {
	return &MsgCreateData{
		Creator: creator,
		Data:    data,
		Meta:    meta,
	}
}

func (msg *MsgCreateData) Route() string {
	return RouterKey
}

func (msg *MsgCreateData) Type() string {
	return TypeMsgCreateData
}

func (msg *MsgCreateData) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateData) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
