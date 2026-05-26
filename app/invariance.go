package app

import (
	"context"

	servertypes "github.com/sei-protocol/sei-chain/sei-cosmos/server/types"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

type LightInvarianceConfig struct {
	SupplyEnabled bool `mapstructure:"supply_enabled"`
}

var DefaultLightInvarianceConfig = LightInvarianceConfig{
	SupplyEnabled: true,
}

const (
	flagSupplyEnabled = "light_invariance.supply_enabled"
)

func ReadLightInvarianceConfig(opts servertypes.AppOptions) (LightInvarianceConfig, error) {
	_ = "STUB: not implemented"
	return *new(LightInvarianceConfig), nil
}

// copy

func (app *App) LightInvarianceChecks(ctx context.Context, cms sdk.CommitMultiStore, config LightInvarianceConfig) {
	_ = "STUB: not implemented"
	// Skip invariance checks when mock_balances is enabled since we fake balances
	// without updating the actual store, which would fail the supply check.
	return
}

func (app *App) LightInvarianceTotalSupply(ctx context.Context, cms sdk.CommitMultiStore) {
	_ = "STUB: not implemented"
	return
}

// TODO(PLT-327): remove once app_lightinvariance_supply_duration_seconds verified

// invalid key; ignore
// TODO(PLT-327): remove once app_lightinvariance_supply_invalid_key_total verified

// invalid key length; ignore
// TODO(PLT-327): remove once app_lightinvariance_supply_invalid_key_total verified

// TODO(PLT-327): remove once app_lightinvariance_supply_unmarshal_failure_total verified

// TODO(PLT-327): remove once app_lightinvariance_supply_unmarshal_failure_total verified

// TODO(PLT-327): remove once app_lightinvariance_supply_invalid_key_total verified

// TODO(PLT-327): remove once app_lightinvariance_supply_unmarshal_failure_total verified

// TODO(PLT-327): remove once app_lightinvariance_supply_unmarshal_failure_total verified

// TODO(PLT-327): remove once app_lightinvariance_supply_unmarshal_failure_total verified

// TODO(PLT-327): remove once app_lightinvariance_supply_unmarshal_failure_total verified

// Formula: useiDiff = useiPreTotal - useiPostTotal - weiDiffInUsei + supplyChanged
// If money is conserved, this should be zero
// useiPreTotal - useiPostTotal = how much usei left balances (negative means usei entered balances)
// weiDiffInUsei = how much usei was moved to wei balances
// supplyChanged = how much new usei was minted
