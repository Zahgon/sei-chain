package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func (p Plan) String() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic does basic validation of a Plan
func (p Plan) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// ShouldExecute returns true if the Plan is ready to execute given the current context
func (p Plan) ShouldExecute(ctx sdk.Context) bool { _ = "STUB: not implemented"; return false }

// DueAt is a string representation of when this plan is due to be executed
func (p Plan) DueAt() string { _ = "STUB: not implemented"; return "" }

// UpgradeDetails is a struct that represents the details of an upgrade
// This is held in the Info object of an upgrade Plan
type UpgradeDetails struct {
	UpgradeType string `json:"upgradeType"`
}

// UpgradeDetails parses and returns a details struct from the Info field of a Plan
// The upgrade.pb.go is generated from proto, so this is separated here
func (p Plan) UpgradeDetails() (UpgradeDetails, error) {
	_ = "STUB: not implemented"
	return *new(UpgradeDetails), nil
}

// invalid json, assume no upgrade details

// IsMinorRelease returns true if the upgrade is a minor release
func (ud UpgradeDetails) IsMinorRelease() bool { _ = "STUB: not implemented"; return false }
