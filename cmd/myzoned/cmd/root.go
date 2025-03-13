package cmd

import (
	"myzone/app"
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	"github.com/spf13/cobra"
)

// Version is set during build
const Version = "v1.0.0"

func NewRootCmd() (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:     "myzoned",
		Short:   "MyZone Daemon (server)",
		Version: Version,
	}

	// Add subcommands
	rootCmd.AddCommand(
		VersionCmd(),
		server.StartCmd(app.NewAppCreator(), app.DefaultNodeHome),
	)

	// Add our custom send-with-fee command
	AddSendWithFeeCommand(rootCmd)

	// Add completion command
	rootCmd.AddCommand(&cobra.Command{
		Use:   "completion",
		Short: "Generate completion script",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Root().GenBashCompletion(os.Stdout)
		},
	})

	return rootCmd, nil
}

// VersionCmd returns a CLI command that prints the version.
func VersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the application version",
		Run: func(_ *cobra.Command, _ []string) {
			println(Version)
		},
	}

	return cmd
}
