package mock

import (
	"context"

	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
	tmtypes "github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

var _ tmtypes.PrivValidator = PV{}

// MockPV implements PrivValidator without any safety or persistence.
// Only use it for testing.
type PV struct {
	PrivKey cryptotypes.PrivKey
}

func NewPV() PV { _ = "STUB: not implemented"; return *new(PV) }

// GetPubKey implements PrivValidator interface
func (pv PV) GetPubKey(_ context.Context) (crypto.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PubKey), nil
}

// SignVote implements PrivValidator interface
func (pv PV) SignVote(_ context.Context, chainID string, vote *tmproto.Vote) error {
	_ = "STUB: not implemented"
	return nil
}

// SignProposal implements PrivValidator interface
func (pv PV) SignProposal(_ context.Context, chainID string, proposal *tmproto.Proposal) error {
	_ = "STUB: not implemented"
	return nil
}
