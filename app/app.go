package app

import (
	"io"
	"os"
	"path/filepath"

	mybankkeeper "myzone/x/mybank/keeper"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	abci "github.com/cometbft/cometbft/abci/types"
	cmtlog "github.com/cometbft/cometbft/libs/log"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

const Name = "MyZoneApp"

// DefaultNodeHome default home directories for the application daemon
var DefaultNodeHome string

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, ".myzone")
}

var (
	_ servertypes.Application = (*MyZoneApp)(nil)
)

type MyZoneApp struct {
	*baseapp.BaseApp

	// keepers
	AccountKeeper keeper.AccountKeeper
	BankKeeper    bankkeeper.Keeper
	FeeKeeper     mybankkeeper.FeeKeeper

	// codec and store services
	cdc       codec.Codec
	appCodec  codec.BinaryCodec
	storeKeys map[string]*storetypes.KVStoreKey
	logger    log.Logger
}

// RegisterNodeService implements servertypes.Application
func (app *MyZoneApp) RegisterNodeService(clientCtx client.Context, cfg config.Config) {
	// Empty implementation for now
}

// RegisterTendermintService implements servertypes.Application
func (app *MyZoneApp) RegisterTendermintService(clientCtx client.Context) {
	// Empty implementation for now
}

// RegisterTxService implements servertypes.Application
func (app *MyZoneApp) RegisterTxService(clientCtx client.Context) {
	// Empty implementation for now
}

func NewMyZoneApp(
	cometLogger cmtlog.Logger,
	db dbm.DB,
	traceStore io.Writer,
	baseAppOptions ...func(*baseapp.BaseApp),
) *MyZoneApp {
	// Convert CometBFT logger to Cosmos SDK logger
	logger := log.NewLogger(io.Discard)
	if cometLogger != nil {
		// In a real implementation, you'd use a proper adapter
		// This is just to fix the type error
	}

	bApp := baseapp.NewBaseApp(Name, logger, db, nil, baseAppOptions...)

	// Initialize app instance
	app := &MyZoneApp{
		BaseApp: bApp,
		logger:  logger,
	}

	// In a real implementation, you would initialize:
	// - app.cdc (codec)
	// - app.appCodec (binary codec)
	// - app.storeKeys (store service)
	// - app.AccountKeeper

	// For this demonstration, we're not fully initializing the BankKeeper
	// Instead, we assume the BankKeeper is already initialized elsewhere

	// Initialize the fee keeper that wraps the bank keeper
	// In a real app, you would use the properly initialized BankKeeper
	app.FeeKeeper = mybankkeeper.NewFeeKeeper(app.BankKeeper)

	return app
}

func (app *MyZoneApp) InitChainer(ctx sdk.Context, req *abci.RequestInitChain) (*abci.ResponseInitChain, error) {
	return &abci.ResponseInitChain{}, nil
}

func (app *MyZoneApp) BeginBlocker(ctx sdk.Context) (sdk.BeginBlock, error) {
	return sdk.BeginBlock{}, nil
}

func (app *MyZoneApp) EndBlocker(ctx sdk.Context) (sdk.EndBlock, error) {
	return sdk.EndBlock{}, nil
}

func (app *MyZoneApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	// Empty implementation
}
