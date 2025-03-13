package main

import (
	"myzone/app"             // Line 6
	"myzone/cmd/myzoned/cmd" // Line 7
	"os"

	srvcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, err := cmd.NewRootCmd()
	if err != nil {
		os.Exit(1)
	}

	if err := srvcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
