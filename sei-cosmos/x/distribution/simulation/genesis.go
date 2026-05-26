package simulation

// DONTCOVER

import (
	"math/rand"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
)

// Simulation parameter constants
const (
	CommunityTax        = "community_tax"
	BaseProposerReward  = "base_proposer_reward"
	BonusProposerReward = "bonus_proposer_reward"
	WithdrawEnabled     = "withdraw_enabled"
)

// GenCommunityTax randomized CommunityTax
func GenCommunityTax(r *rand.Rand) sdk.Dec { _ = "STUB: not implemented"; return *new(sdk.Dec) }

// GenBaseProposerReward randomized BaseProposerReward
func GenBaseProposerReward(r *rand.Rand) sdk.Dec { _ = "STUB: not implemented"; return *new(sdk.Dec) }

// GenBonusProposerReward randomized BonusProposerReward
func GenBonusProposerReward(r *rand.Rand) sdk.Dec { _ = "STUB: not implemented"; return *new(sdk.Dec) }

// GenWithdrawEnabled returns a randomized WithdrawEnabled parameter.
func GenWithdrawEnabled(r *rand.Rand) bool { _ = "STUB: not implemented"; return false }

// 95% chance of withdraws being enabled

// RandomizedGenState generates a random GenesisState for distribution
func RandomizedGenState(simState *module.SimulationState) { _ = "STUB: not implemented"; return }
