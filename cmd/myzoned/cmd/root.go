package cmd

import (
	"myzone/app"
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/spf13/cobra"
)

func NewRootCmd() (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:   "myzoned",
		Short: "MyZone Daemon (server)",
	}

	server.AddCommands(rootCmd, app.DefaultNodeHome, app.NewAppCreator().(servertypes.AppCreator), nil, nil)
	rootCmd.AddCommand(&cobra.Command{
		Use:   "completion",
		Short: "Generate completion script",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Root().GenBashCompletion(os.Stdout)
		},
	})

	return rootCmd, nil
}
