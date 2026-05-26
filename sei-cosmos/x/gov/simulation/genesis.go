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
	DepositParamsMinDeposit           = "deposit_params_min_deposit"
	DepositParamsMinExpeditedDeposit  = "deposit_params_min_expedited_deposit"
	DepositParamsDepositPeriod        = "deposit_params_deposit_period"
	VotingParamsVotingPeriod          = "voting_params_voting_period"
	ExpeditedVotingParamsVotingPeriod = "expedited_voting_params_voting_period"
	TallyParamsQuorum                 = "tally_params_quorum"
	TallyParamsExpeditedQuorum        = "tally_params_expedited_quorum"
	TallyParamsThreshold              = "tally_params_threshold"
	TallyParamsExpeditedThreshold     = "tally_params_expedited_threshold"
	TallyParamsVeto                   = "tally_params_veto"
)

// GenDepositParamsDepositPeriod randomized DepositParamsDepositPeriod
func GenDepositParamsDepositPeriod(r *rand.Rand) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// GenDepositParamsMinDeposit randomized DepositParamsMinDeposit
func GenDepositParamsMinDeposit(r *rand.Rand) sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

// GenDepositParamsMinExpeditedDeposit randomized DepositParamsMinExpeditedDeposit
func GenDepositParamsMinExpeditedDeposit(r *rand.Rand) sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

// GenVotingParamsVotingPeriod randomized VotingParamsVotingPeriod
func GenVotingParamsVotingPeriod(r *rand.Rand) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// GenVotingParamsExpeditedVotingPeriod randomized VotingParamsExpeditedVotingPeriod
func GenVotingParamsExpeditedVotingPeriod(r *rand.Rand) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// GenTallyParamsQuorum randomized TallyParamsQuorum
func GenTallyParamsQuorum(r *rand.Rand) sdk.Dec { _ = "STUB: not implemented"; return *new(sdk.Dec) }

// GenTallyExpeditedParamsQuorum randomized TallyParamsExpeditedQuorum
func GenTallyExpeditedParamsQuorum(r *rand.Rand) sdk.Dec {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

// GenTallyParamsThreshold randomized TallyParamsThreshold
func GenTallyParamsThreshold(r *rand.Rand) sdk.Dec { _ = "STUB: not implemented"; return *new(sdk.Dec) }

// GenTallyParamsExpeditedThreshold randomized TallyParamsExpeditedThreshold
func GenTallyParamsExpeditedThreshold(r *rand.Rand) sdk.Dec {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

// GenTallyParamsVeto randomized TallyParamsVeto
func GenTallyParamsVeto(r *rand.Rand) sdk.Dec { _ = "STUB: not implemented"; return *new(sdk.Dec) }

// RandomizedGenState generates a random GenesisState for gov
func RandomizedGenState(simState *module.SimulationState) { _ = "STUB: not implemented"; return }

//nolint:gosec // Intn(100) always returns a non-negative value
