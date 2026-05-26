package utils

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
)

const (
	defaultPage  = 1
	defaultLimit = 30 // should be consistent with tendermint/tendermint/rpc/core/pipe.go:19
)

// Proposer contains metadata of a governance proposal used for querying a
// proposer.
type Proposer struct {
	ProposalID uint64 `json:"proposal_id" yaml:"proposal_id"`
	Proposer   string `json:"proposer" yaml:"proposer"`
}

// NewProposer returns a new Proposer given id and proposer
func NewProposer(proposalID uint64, proposer string) Proposer {
	_ = "STUB: not implemented"
	return *new(Proposer)
}

func (p Proposer) String() string { _ = "STUB: not implemented"; return "" }

// QueryDepositsByTxQuery will query for deposits via a direct txs tags query. It
// will fetch and build deposits directly from the returned txs and return a
// JSON marshalled result or any error that occurred.
//
// NOTE: SearchTxs is used to facilitate the txs query which does not currently
// support configurable pagination.
func QueryDepositsByTxQuery(ctx context.Context, clientCtx client.Context, params types.QueryProposalParams) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// initial deposit was submitted with proposal, so must be queried separately
		nil
}

// Query legacy Msgs event action

// Query proto Msgs event action

// QueryVotesByTxQuery will query for votes via a direct txs tags query. It
// will fetch and build votes directly from the returned txs and return a JSON
// marshalled result or any error that occurred.
func QueryVotesByTxQuery(ctx context.Context, clientCtx client.Context, params types.QueryProposalVotesParams) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// query interrupted either if we collected enough votes or tx indexer run out of relevant txs

// Search for both (legacy) votes and weighted votes.

// Query legacy Vote Msgs

// Query Vote proto Msgs

// Query legacy VoteWeighted Msgs

// Query VoteWeighted proto Msgs

// QueryVoteByTxQuery will query for a single vote via a direct txs tags query.
func QueryVoteByTxQuery(ctx context.Context, clientCtx client.Context, params types.QueryVoteParams) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Query legacy Vote Msgs

// Query Vote proto Msgs

// Query legacy VoteWeighted Msgs

// Query VoteWeighted proto Msgs

// there should only be a single vote under the given conditions

// QueryDepositByTxQuery will query for a single deposit via a direct txs tags
// query.
func QueryDepositByTxQuery(ctx context.Context, clientCtx client.Context, params types.QueryDepositParams) ([]byte, error) {
	_ = "STUB: not implemented"

	// initial deposit was submitted with proposal, so must be queried separately
	return nil, nil
}

// Query legacy Msgs event action

// Query proto Msgs event action

// there should only be a single deposit under the given conditions

// QueryProposerByTxQuery will query for a proposer of a governance proposal by
// ID.
func QueryProposerByTxQuery(ctx context.Context, clientCtx client.Context, proposalID uint64) (Proposer, error) {
	_ = "STUB: not implemented"
	return *new(Proposer), nil
}

// Query legacy Msgs event action

// Query proto Msgs event action

// there should only be a single proposal under the given conditions

// QueryProposalByID takes a proposalID and returns a proposal
func QueryProposalByID(proposalID uint64, clientCtx client.Context, queryRoute string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// combineEvents queries txs by events with all events from each event group,
// and combines all those events together.
//
// Tx are indexed in tendermint via their Msgs `Type()`, which can be:
// - via legacy Msgs (amino or proto), their `Type()` is a custom string,
// - via ADR-031 proto msgs, their `Type()` is the protobuf FQ method name.
// In searching for events, we search for both `Type()`s, and we use the
// `combineEvents` function here to merge events.
func combineEvents(ctx context.Context, clientCtx client.Context, page int, eventGroups ...[]string) (*sdk.SearchTxsResult, error) {
	_ = "STUB: not implemented"
	// Only the Txs field will be populated in the final SearchTxsResult.
	return nil, nil
}

// queryInitialDepositByTxQuery will query for a initial deposit of a governance proposal by
// ID.
func queryInitialDepositByTxQuery(ctx context.Context, clientCtx client.Context, proposalID uint64) (types.Deposit, error) {
	_ = "STUB: not implemented"
	return *new(types.Deposit), nil
}

// Query legacy Msgs event action

// Query proto Msgs event action

// there should only be a single proposal under the given conditions
