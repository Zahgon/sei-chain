package cli

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/distribution/types"
)

// ParseCommunityPoolSpendProposalWithDeposit reads and parses a CommunityPoolSpendProposalWithDeposit from a file.
func ParseCommunityPoolSpendProposalWithDeposit(cdc codec.JSONCodec, proposalFile string) (types.CommunityPoolSpendProposalWithDeposit, error) {
	_ = "STUB: not implemented"
	return *new(types.CommunityPoolSpendProposalWithDeposit), nil
}
