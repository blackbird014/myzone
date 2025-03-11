package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgSend struct {
	FromAddress string    `json:"from_address"`
	ToAddress   string    `json:"to_address"`
	Amount      sdk.Coins `json:"amount"`
}

func (msg MsgSend) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.FromAddress)
	return []sdk.AccAddress{addr}
}
