package app

import (
	"embed"
)

//go:embed tags
var f embed.FS

// NOTE: When performing upgrades, make sure to keep / register the handlers
// for both the current (n) and the previous (n-1) upgrade name. There is a bug
// in a missing value in a log statement for which the fix is not released
var upgradesList []string

var LatestUpgrade string

func init() {
	content, err := f.ReadFile("tags")
	if err != nil {
		panic(err)
	}
	upgradesList = parseUpgradesList(string(content))
	LatestUpgrade = upgradesList[len(upgradesList)-1]
}

func parseUpgradesList(list string) []string { _ = "STUB: not implemented"; return nil }

// Upgrades names must be in alphabetical order
// https://github.com/cosmos/cosmos-sdk/issues/11707

// if there is an override list, use that instead, for integration tests
func overrideList() {
	_ = "STUB: not implemented"
	// if there is an override list, use that instead, for integration tests
	return
}

func (app *App) RegisterUpgradeHandlers() {
	_ = "STUB: not implemented"
	// if there is an override list, use that instead, for integration tests
	return
}

// Set params to Distribution here when migrating

// 50 mil

const v606UpgradeHeight = 151573570
