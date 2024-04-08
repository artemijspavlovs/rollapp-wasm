package v1

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	hubgentypes "github.com/dymensionxyz/dymension-rdk/x/hub-genesis/types"
)

func GetStoreUpgrades() *storetypes.StoreUpgrades {
	storeUpgrades := storetypes.StoreUpgrades{
		// Set migrations for all new modules
		Added: []string{
			hubgentypes.ModuleName,
		},
	}
	return &storeUpgrades
}

// CreateUpgradeHandler creates an SDK upgrade handler for v3
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		logger := ctx.Logger().With("upgrade", UpgradeName)

		// Start running the module migrations
		logger.Debug("running module migrations ...")
		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}
