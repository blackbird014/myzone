package app

import (
	"io"

	cmtlog "github.com/cometbft/cometbft/libs/log"
	dbm "github.com/cosmos/cosmos-db"
)

// NewAppCreator returns a function that creates a new MyZoneApp instance.
// This is a simplified version to avoid type conflicts with the SDK.
func NewAppCreator() interface{} {
	return func(
		logger cmtlog.Logger,
		db dbm.DB,
		traceStore io.Writer,
		_ interface{},
	) *MyZoneApp {
		return NewMyZoneApp(logger, db, traceStore)
	}
}
