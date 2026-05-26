package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func NewGenesisState(
	params Params, fp FeePool, dwis []DelegatorWithdrawInfo, pp sdk.ConsAddress, r []ValidatorOutstandingRewardsRecord,
	acc []ValidatorAccumulatedCommissionRecord, historical []ValidatorHistoricalRewardsRecord,
	cur []ValidatorCurrentRewardsRecord, dels []DelegatorStartingInfoRecord, slashes []ValidatorSlashEventRecord,
) *GenesisState {
	_ = "STUB: not implemented"
	return nil
}

// get raw genesis raw message for testing
func DefaultGenesisState() *GenesisState { _ = "STUB: not implemented"; return nil }

// ValidateGenesis validates the genesis state of distribution genesis input
func ValidateGenesis(gs *GenesisState) error { _ = "STUB: not implemented"; return nil }
