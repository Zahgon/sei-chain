package types

import (
	"time"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
)

// Default parameter namespace
const (
	DefaultSignedBlocksWindow   = int64(108000) // ~12 hours based on 0.4s block times
	DefaultDowntimeJailDuration = 60 * 10 * time.Second
)

var (
	DefaultMinSignedPerWindow = sdk.NewDecWithPrec(5, 2)
	// No Slashing Fraction by default
	DefaultSlashFractionDoubleSign = sdk.NewDec(0)
	DefaultSlashFractionDowntime   = sdk.NewDec(0)
)

// Parameter store keys
var (
	KeySignedBlocksWindow      = []byte("SignedBlocksWindow")
	KeyMinSignedPerWindow      = []byte("MinSignedPerWindow")
	KeyDowntimeJailDuration    = []byte("DowntimeJailDuration")
	KeySlashFractionDoubleSign = []byte("SlashFractionDoubleSign")
	KeySlashFractionDowntime   = []byte("SlashFractionDowntime")
)

// ParamKeyTable for slashing module
func ParamKeyTable() paramtypes.KeyTable {
	_ = "STUB: not implemented"
	return *new(paramtypes.KeyTable)
}

// NewParams creates a new Params object
func NewParams(
	signedBlocksWindow int64, minSignedPerWindow sdk.Dec, downtimeJailDuration time.Duration,
	slashFractionDoubleSign, slashFractionDowntime sdk.Dec,
) Params {
	_ = "STUB: not implemented"
	return *new(Params)
}

// ParamSetPairs - Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	_ = "STUB: not implemented"
	return *new(paramtypes.ParamSetPairs)
}

// DefaultParams defines the parameters for this module
func DefaultParams() Params { _ = "STUB: not implemented"; return *new(Params) }

func validateSignedBlocksWindow(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateMinSignedPerWindow(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateDowntimeJailDuration(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateSlashFractionDoubleSign(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateSlashFractionDowntime(i interface{}) error { _ = "STUB: not implemented"; return nil }
