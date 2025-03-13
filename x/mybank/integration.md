# Integrating FeeKeeper with Your Application

This document describes how to integrate the `FeeKeeper` with your Cosmos SDK application to add a fee burning functionality to coin transfers.

## Overview

The `FeeKeeper` wraps the standard bank module's keeper and adds a fee burning feature. When coins are transferred using the `SendCoinsWithFee` method, a small fee (0.1% by default) is deducted and sent to a burn address.

## Integration Steps

### 1. Initialize the FeeKeeper in app.go

In your `app.go` file, add the following after initializing the bank keeper:

```go
// Initialize the bank keeper
bankKeeper := bankkeeper.NewBaseKeeper(
    appCodec,
    keys[banktypes.StoreKey],
    app.AccountKeeper,
    app.GetSubspace(banktypes.ModuleName),
    app.ModuleAccountAddrs(),
)

// Initialize the fee keeper that wraps the bank keeper
feeKeeper := mybankkeeper.NewFeeKeeper(bankKeeper)
```

### 2. Use the FeeKeeper in Your Module or Command

You can use the `FeeKeeper` in various parts of your application, such as custom modules or CLI commands:

```go
// Example usage in a message handler
func HandleSendWithFeeMsg(ctx sdk.Context, feeKeeper mybankkeeper.FeeKeeper, msg types.MsgSendWithFee) error {
    fromAddr, _ := sdk.AccAddressFromBech32(msg.FromAddress)
    toAddr, _ := sdk.AccAddressFromBech32(msg.ToAddress)
    
    // Send coins with fee burning
    err := feeKeeper.SendCoinsWithFee(ctx, fromAddr, toAddr, msg.Amount)
    if err != nil {
        return err
    }
    
    return nil
}
```

### 3. Add CLI Commands

You can expose the fee burning functionality through CLI commands:

```go
// Example CLI command
func GetSendWithFeeTxCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "send-with-fee [from_key_or_address] [to_address] [amount]",
        Short: "Send coins from one account to another with fee burning",
        Args:  cobra.ExactArgs(3),
        RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx, err := client.GetClientTxContext(cmd)
            if err != nil {
                return err
            }
            
            fromAddr, _ := sdk.AccAddressFromBech32(args[0])
            toAddr, _ := sdk.AccAddressFromBech32(args[1])
            amount, _ := sdk.ParseCoinsNormalized(args[2])
            
            msg := types.NewMsgSendWithFee(fromAddr.String(), toAddr.String(), amount)
            
            return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
        },
    }
    
    return cmd
}
```

## Example Usage

Here's an example of how to use the `FeeKeeper` in a transaction:

1. For a transfer of 100 tokens, the fee would be 0.1 tokens (0.1%).
2. The recipient receives 99.9 tokens.
3. The 0.1 token fee is sent to the burn address.

This functionality can be used for various purposes, such as reducing the total supply of tokens over time or funding community initiatives. 