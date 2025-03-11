package keeper

import (
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/keeper"
)

type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey storetypes.StoreKey // Fixed: use storetypes.StoreKey
	ak       keeper.AccountKeeper
}

func NewKeeper(cdc codec.BinaryCodec, storeKey storetypes.StoreKey, ak keeper.AccountKeeper) Keeper {
	return Keeper{cdc: cdc, storeKey: storeKey, ak: ak}
}

func (k Keeper) SendCoins(ctx sdk.Context, fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) error {
	store := ctx.KVStore(k.storeKey)

	// Get sender's coins (raw bytes to sdk.Coins)
	var fromCoins sdk.Coins
	if bz := store.Get(fromAddr); bz != nil {
		parsedCoins, err := sdk.ParseCoinsNormalized(string(bz))
		if err != nil {
			return err
		}
		fromCoins = parsedCoins
	}

	// Check sufficient funds
	if !fromCoins.IsAllGTE(amt) {
		return sdkerrors.ErrInsufficientFunds
	}

	// Update sender's balance
	fromCoins = fromCoins.Sub(amt...)
	bz := []byte(fromCoins.String())
	store.Set(fromAddr, bz)

	// Get recipient's coins
	var toCoins sdk.Coins
	if bz := store.Get(toAddr); bz != nil {
		parsedCoins, err := sdk.ParseCoinsNormalized(string(bz))
		if err != nil {
			return err
		}
		toCoins = parsedCoins
	}

	// Update recipient's balance
	toCoins = toCoins.Add(amt...)
	bz = []byte(toCoins.String())
	store.Set(toAddr, bz)

	return nil
}
