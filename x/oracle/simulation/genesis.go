package simulation

// DONTCOVER

import (
	"math/rand"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
)

// Simulation parameter constants
const (
	votePeriodKey               = "vote_period"
	voteThresholdKey            = "vote_threshold"
	rewardBandKey               = "reward_band"
	rewardDistributionWindowKey = "reward_distribution_window" //nolint:unused // we will use this later
	slashFractionKey            = "slash_fraction"
	slashWindowKey              = "slash_window"
	minValidPerWindowKey        = "min_valid_per_window"
)

// GenVotePeriod randomized VotePeriod
func GenVotePeriod(r *rand.Rand) uint64 { _ = "STUB: not implemented"; return 0 }

//nolint:gosec

// GenVoteThreshold randomized VoteThreshold
func GenVoteThreshold(r *rand.Rand) sdk.Dec { _ = "STUB: not implemented"; return *new(sdk.Dec) }

// GenRewardBand randomized RewardBand
func GenRewardBand(r *rand.Rand) sdk.Dec { _ = "STUB: not implemented"; return *new(sdk.Dec) }

// GenSlashFraction randomized SlashFraction
func GenSlashFraction(r *rand.Rand) sdk.Dec { _ = "STUB: not implemented"; return *new(sdk.Dec) }

// GenSlashWindow randomized SlashWindow
func GenSlashWindow(r *rand.Rand) uint64 { _ = "STUB: not implemented"; return 0 }

//nolint:gosec

// GenMinValidPerWindow randomized MinValidPerWindow
func GenMinValidPerWindow(r *rand.Rand) sdk.Dec { _ = "STUB: not implemented"; return *new(sdk.Dec) }

// RandomizedGenState generates a random GenesisState for oracle
func RandomizedGenState(simState *module.SimulationState) { _ = "STUB: not implemented"; return }
