package keeper

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

// FeeKeeper wraps the bank keeper and adds fee burning functionality
type FeeKeeper struct {
	bankKeeper bankkeeper.Keeper
	feeRate    math.LegacyDec // Fee rate as a decimal (e.g., 0.001 for 0.1%)
	burnAddr   sdk.AccAddress
}

// NewFeeKeeper creates a new FeeKeeper with a fixed fee rate of 0.001 (0.1%)
func NewFeeKeeper(bankKeeper bankkeeper.Keeper) FeeKeeper {
	// Create a burn address - in a real implementation, this should be a module account
	// or an address that is guaranteed to be unspendable
	burnAddr := sdk.AccAddress([]byte("burn"))

	return FeeKeeper{
		bankKeeper: bankKeeper,
		feeRate:    math.LegacyNewDecWithPrec(1, 3), // 0.001 (0.1%)
		burnAddr:   burnAddr,
	}
}

// SendCoinsWithFee sends coins from one account to another, deducting a fee that gets burnt
func (k FeeKeeper) SendCoinsWithFee(ctx sdk.Context, fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) error {
	// Calculate fee to be burnt
	fee := sdk.NewCoins()
	for _, coin := range amt {
		// Calculate fee amount (amount * fee rate)
		// Convert to Dec, multiply by rate, convert back to Int
		decAmount := math.LegacyNewDecFromInt(coin.Amount)
		feeAmount := decAmount.Mul(k.feeRate).TruncateInt()
		if feeAmount.IsPositive() {
			fee = fee.Add(sdk.NewCoin(coin.Denom, feeAmount))
		}
	}

	// Calculate amount after fee
	amtAfterFee := amt.Sub(fee...)

	// Send coins to recipient
	err := k.bankKeeper.SendCoins(ctx, fromAddr, toAddr, amtAfterFee)
	if err != nil {
		return err
	}

	// Burn the fee by sending it to the burn address
	if !fee.IsZero() {
		err = k.bankKeeper.SendCoins(ctx, fromAddr, k.burnAddr, fee)
		if err != nil {
			return err
		}

		// Emit an event for the fee burn
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				"fee_burn",
				sdk.NewAttribute("sender", fromAddr.String()),
				sdk.NewAttribute("amount", fee.String()),
				sdk.NewAttribute("burn_address", k.burnAddr.String()),
			),
		)
	}

	return nil
}

// SendCoins forwards to the underlying bank keeper's SendCoins method
// This allows using this keeper as a drop-in replacement for the bank keeper
func (k FeeKeeper) SendCoins(ctx sdk.Context, fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) error {
	return k.bankKeeper.SendCoins(ctx, fromAddr, toAddr, amt)
}

// GetBalance forwards to the underlying bank keeper's GetBalance method
func (k FeeKeeper) GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	return k.bankKeeper.GetBalance(ctx, addr, denom)
}

// GetAllBalances forwards to the underlying bank keeper's GetAllBalances method
func (k FeeKeeper) GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins {
	return k.bankKeeper.GetAllBalances(ctx, addr)
}

// GetSupply forwards to the underlying bank keeper's GetSupply method
func (k FeeKeeper) GetSupply(ctx sdk.Context, denom string) sdk.Coin {
	return k.bankKeeper.GetSupply(ctx, denom)
}
