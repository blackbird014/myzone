package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Balances struct {
	Coins sdk.Coins `json:"coins"`
}

func NewBalances() Balances {
	return Balances{Coins: sdk.NewCoins()}
}
