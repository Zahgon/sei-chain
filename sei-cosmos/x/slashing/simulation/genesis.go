package simulation

// DONTCOVER

import (
	"math/rand"
	"time"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
)

// Simulation parameter constants
const (
	SignedBlocksWindow      = "signed_blocks_window"
	MinSignedPerWindow      = "min_signed_per_window"
	DowntimeJailDuration    = "downtime_jail_duration"
	SlashFractionDoubleSign = "slash_fraction_double_sign"
	SlashFractionDowntime   = "slash_fraction_downtime"
)

// GenSignedBlocksWindow randomized SignedBlocksWindow
func GenSignedBlocksWindow(r *rand.Rand) int64 { _ = "STUB: not implemented"; return 0 }

// GenMinSignedPerWindow randomized MinSignedPerWindow
func GenMinSignedPerWindow(r *rand.Rand) sdk.Dec { _ = "STUB: not implemented"; return *new(sdk.Dec) }

// GenDowntimeJailDuration randomized DowntimeJailDuration
func GenDowntimeJailDuration(r *rand.Rand) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// GenSlashFractionDoubleSign randomized SlashFractionDoubleSign
func GenSlashFractionDoubleSign(r *rand.Rand) sdk.Dec {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

// GenSlashFractionDowntime randomized SlashFractionDowntime
func GenSlashFractionDowntime(r *rand.Rand) sdk.Dec {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

// RandomizedGenState generates a random GenesisState for slashing
func RandomizedGenState(simState *module.SimulationState) { _ = "STUB: not implemented"; return }
