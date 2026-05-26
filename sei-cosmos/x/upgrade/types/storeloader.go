package types

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/baseapp"
	store "github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
)

// UpgradeStoreLoader is used to prepare baseapp with a fixed StoreLoader
// pattern. This is useful for custom upgrade loading logic.
func UpgradeStoreLoader(upgradeHeight int64, storeUpgrades *store.StoreUpgrades) baseapp.StoreLoader {
	_ = "STUB: not implemented"
	return *new(baseapp.StoreLoader)
}

// Check if the current commit version and upgrade height matches

// Otherwise load default store loader
