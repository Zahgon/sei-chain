package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// create a new ValidatorHistoricalRewards
func NewValidatorHistoricalRewards(cumulativeRewardRatio sdk.DecCoins, referenceCount uint32) ValidatorHistoricalRewards {
	_ = "STUB: not implemented"
	return *new(ValidatorHistoricalRewards)
}

// create a new ValidatorCurrentRewards
func NewValidatorCurrentRewards(rewards sdk.DecCoins, period uint64) ValidatorCurrentRewards {
	_ = "STUB: not implemented"
	return *new(ValidatorCurrentRewards)
}

// return the initial accumulated commission (zero)
func InitialValidatorAccumulatedCommission() ValidatorAccumulatedCommission {
	_ = "STUB: not implemented"
	return *new(ValidatorAccumulatedCommission)
}

// create a new ValidatorSlashEvent
func NewValidatorSlashEvent(validatorPeriod uint64, fraction sdk.Dec) ValidatorSlashEvent {
	_ = "STUB: not implemented"
	return *new(ValidatorSlashEvent)
}

func (vs ValidatorSlashEvents) String() string { _ = "STUB: not implemented"; return "" }
