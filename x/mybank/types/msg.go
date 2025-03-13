package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
)

// MsgSend is a basic message for coin transfers
type MsgSend struct {
	FromAddress string    `json:"from_address"`
	ToAddress   string    `json:"to_address"`
	Amount      sdk.Coins `json:"amount"`
}

func (msg MsgSend) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.FromAddress)
	return []sdk.AccAddress{addr}
}

// MsgSendWithFee defines a message for sending coins with a fee
type MsgSendWithFee struct {
	FromAddress string    `json:"from_address"`
	ToAddress   string    `json:"to_address"`
	Amount      sdk.Coins `json:"amount"`
}

// NewMsgSendWithFee creates a new MsgSendWithFee instance
func NewMsgSendWithFee(fromAddr, toAddr string, amount sdk.Coins) *MsgSendWithFee {
	return &MsgSendWithFee{
		FromAddress: fromAddr,
		ToAddress:   toAddr,
		Amount:      amount,
	}
}

// GetSigners returns the signers of the message
func (msg MsgSendWithFee) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.FromAddress)
	return []sdk.AccAddress{addr}
}

// ProtoMessage implements the proto.Message interface
func (msg MsgSendWithFee) ProtoMessage() {}

// Reset implements the proto.Message interface
func (msg *MsgSendWithFee) Reset() {
	*msg = MsgSendWithFee{}
}

// String implements the proto.Message interface
func (msg MsgSendWithFee) String() string {
	return proto.CompactTextString(&msg)
}
