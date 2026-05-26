package types

import (
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

var _ exported.ConsensusState = &ConsensusState{}

// ClientType returns Solo Machine type.
func (ConsensusState) ClientType() string { _ = "STUB: not implemented"; return "" }

// GetTimestamp returns zero.
func (cs ConsensusState) GetTimestamp() uint64 { _ = "STUB: not implemented"; return 0 }

// GetRoot returns nil since solo machines do not have roots.
func (cs ConsensusState) GetRoot() exported.Root {
	_ = "STUB: not implemented"

	// GetPubKey unmarshals the public key into a cryptotypes.PubKey type.
	// An error is returned if the public key is nil or the cached value
	// is not a PubKey.
	return *new(exported.Root)
}

func (cs ConsensusState) GetPubKey() (cryptotypes.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(cryptotypes.PubKey), nil
}

// ValidateBasic defines basic validation for the solo machine consensus state.
func (cs ConsensusState) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
