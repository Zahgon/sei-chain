package keeper

import (
	"github.com/sei-protocol/sei-chain/x/oracle/types"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// OrganizeBallotByDenom collects all oracle votes for the period, categorized by the votes' denom parameter
func (k Keeper) OrganizeBallotByDenom(ctx sdk.Context, validatorClaimMap map[string]types.Claim) (votes map[string]types.ExchangeRateBallot) {
	_ = "STUB: not implemented"
	return nil
}

// Organize aggregate votes

// organize ballot only for the active validators

// Make the power of abstain vote zero

// sort created ballot

// ClearBallots clears all tallied votes from the store
func (k Keeper) ClearBallots(ctx sdk.Context, _ uint64) {
	_ = "STUB: not implemented"
	// Clear all aggregate votes
	return
}

// ApplyWhitelist update vote target denom list with params whitelist
func (k Keeper) ApplyWhitelist(ctx sdk.Context, whitelist types.DenomList, voteTargets map[string]types.Denom) {
	_ = "STUB: not implemented"
	// check is there any update in whitelist params
	return
}

// Register meta data to bank module
