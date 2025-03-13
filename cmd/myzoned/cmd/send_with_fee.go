package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"myzone/x/mybank/types"
)

// GetSendWithFeeTxCmd returns a CLI command for sending coins with fee burning
func GetSendWithFeeTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-with-fee [to_address] [amount]",
		Short: "Send coins with fee burning (0.1% fee)",
		Long: `Send coins from one account to another with a 0.1% fee that gets burned.
Example:
$ myzoned tx send-with-fee cosmos1... 100uatom
`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			toAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgSendWithFee(
				clientCtx.GetFromAddress().String(),
				toAddr.String(),
				amount,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// AddSendWithFeeCommand adds the send-with-fee command to the provided root command
func AddSendWithFeeCommand(rootCmd *cobra.Command) {
	// Add the command directly to the root command
	rootCmd.AddCommand(GetSendWithFeeTxCmd())
}
