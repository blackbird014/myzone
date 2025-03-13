package app

import (
	"encoding/json"
	"io"
	"os"

	"cosmossdk.io/log"
	cmtlog "github.com/cometbft/cometbft/libs/log"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	cmttypes "github.com/cometbft/cometbft/types"
	dbm "github.com/cosmos/cosmos-db"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
)

// NewAppCreator returns a function that creates a new MyZoneApp instance.
// This function returns a servertypes.AppCreator to avoid type conflicts with the SDK.
func NewAppCreator() servertypes.AppCreator {
	return func(
		logger log.Logger,
		db dbm.DB,
		traceStore io.Writer,
		appOpts servertypes.AppOptions,
	) servertypes.Application {
		// Convert the logger if needed
		var cometLogger cmtlog.Logger
		if logger != nil {
			// Simple conversion - in a real app you might want a more sophisticated adapter
			cometLogger = cmtlog.NewTMLogger(os.Stdout)
		}

		return NewMyZoneApp(cometLogger, db, traceStore)
	}
}

// NewAppExporter returns a function that exports the state of the application
func NewAppExporter() servertypes.AppExporter {
	return func(
		logger log.Logger,
		db dbm.DB,
		traceWriter io.Writer,
		height int64,
		forZeroHeight bool,
		jailAllowedAddrs []string,
		opts servertypes.AppOptions,
		modulesToExport []string,
	) (servertypes.ExportedApp, error) {
		// For now, return an empty state
		emptyState, err := json.Marshal(map[string]json.RawMessage{})
		if err != nil {
			return servertypes.ExportedApp{}, err
		}

		return servertypes.ExportedApp{
			AppState:        emptyState,
			Validators:      []cmttypes.GenesisValidator{},
			Height:          height,
			ConsensusParams: cmtproto.ConsensusParams{},
		}, nil
	}
}
