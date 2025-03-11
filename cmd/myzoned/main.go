package main

import (
	"myzone/app"             // Line 6
	"myzone/cmd/myzoned/cmd" // Line 7
	"os"

	srvcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

const Version = "v1.0.0"

func main() {
	rootCmd, err := cmd.NewRootCmd()
	if err != nil {
		os.Exit(1)
	}
	rootCmd.Version = Version

	if err := srvcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
