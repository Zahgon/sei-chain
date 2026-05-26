package keeper

import (
	"time"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	authtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
)

// Keeper defines the governance module Keeper
type Keeper struct {
	// The reference to the Paramstore to get and set gov specific params
	paramSpace types.ParamSubspace

	authKeeper   types.AccountKeeper
	bankKeeper   types.BankKeeper
	paramsKeeper types.ParamsKeeper

	// The reference to the DelegationSet and ValidatorSet to get information about validators and delegators
	sk types.StakingKeeper

	// GovHooks
	hooks types.GovHooks

	// The (unexposed) keys used to access the stores from the Context.
	storeKey sdk.StoreKey

	// The codec codec for binary encoding/decoding.
	cdc codec.BinaryCodec

	// Proposal router
	router types.Router
}

// NewKeeper returns a governance keeper. It handles:
// - submitting governance proposals
// - depositing funds into proposals, and activating upon sufficient funds being deposited
// - users voting on proposals, with weight proportional to stake in the system
// - and tallying the result of the vote.
//
// CONTRACT: the parameter Subspace must have the param key table already initialized
func NewKeeper(
	cdc codec.BinaryCodec, key sdk.StoreKey, paramSpace types.ParamSubspace,
	authKeeper types.AccountKeeper, bankKeeper types.BankKeeper, sk types.StakingKeeper,
	paramsKeeper types.ParamsKeeper, rtr types.Router,
) Keeper {
	_ = "STUB: not implemented"

	// ensure governance module account is set
	return *new(Keeper)
}

// It is vital to seal the governance proposal router here as to not allow
// further handlers to be registered after the keeper is created since this
// could create invalid or non-deterministic behavior.

// SetHooks sets the hooks for governance
func (keeper *Keeper) SetHooks(gh types.GovHooks) *Keeper { _ = "STUB: not implemented"; return nil }

// Router returns the gov Keeper's Router
func (keeper Keeper) Router() types.Router {
	_ = "STUB: not implemented"
	return *

	// GetGovernanceAccount returns the governance ModuleAccount
	new(types.Router)
}

func (keeper Keeper) GetGovernanceAccount(ctx sdk.Context) authtypes.ModuleAccountI {
	_ = "STUB: not implemented"
	return *new(authtypes.ModuleAccountI)
}

// ProposalQueues

// InsertActiveProposalQueue inserts a ProposalID into the active proposal queue at endTime
func (keeper Keeper) InsertActiveProposalQueue(ctx sdk.Context, proposalID uint64, endTime time.Time) {
	_ = "STUB: not implemented"
	return
}

// RemoveFromActiveProposalQueue removes a proposalID from the Active Proposal Queue
func (keeper Keeper) RemoveFromActiveProposalQueue(ctx sdk.Context, proposalID uint64, endTime time.Time) {
	_ = "STUB: not implemented"
	return
}

// InsertInactiveProposalQueue Inserts a ProposalID into the inactive proposal queue at endTime
func (keeper Keeper) InsertInactiveProposalQueue(ctx sdk.Context, proposalID uint64, endTime time.Time) {
	_ = "STUB: not implemented"
	return
}

// RemoveFromInactiveProposalQueue removes a proposalID from the Inactive Proposal Queue
func (keeper Keeper) RemoveFromInactiveProposalQueue(ctx sdk.Context, proposalID uint64, endTime time.Time) {
	_ = "STUB: not implemented"
	return
}

// Iterators

// IterateActiveProposalsQueue iterates over the proposals in the active proposal queue
// and performs a callback function
func (keeper Keeper) IterateActiveProposalsQueue(ctx sdk.Context, endTime time.Time, cb func(proposal types.Proposal) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

// IterateInactiveProposalsQueue iterates over the proposals in the inactive proposal queue
// and performs a callback function
func (keeper Keeper) IterateInactiveProposalsQueue(ctx sdk.Context, endTime time.Time, cb func(proposal types.Proposal) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

// ActiveProposalQueueIterator returns an sdk.Iterator for all the proposals in the Active Queue that expire by endTime
func (keeper Keeper) ActiveProposalQueueIterator(ctx sdk.Context, endTime time.Time) sdk.Iterator {
	_ = "STUB: not implemented"
	return *new(sdk.Iterator)
}

// InactiveProposalQueueIterator returns an sdk.Iterator for all the proposals in the Inactive Queue that expire by endTime
func (keeper Keeper) InactiveProposalQueueIterator(ctx sdk.Context, endTime time.Time) sdk.Iterator {
	_ = "STUB: not implemented"
	return *new(sdk.Iterator)
}
