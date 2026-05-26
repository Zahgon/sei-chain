package types

import (
	"time"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

var _ exported.Header = &Header{}

// ConsensusState returns the updated consensus state associated with the header
func (h Header) ConsensusState() *ConsensusState { _ = "STUB: not implemented"; return nil }

// ClientType defines that the Header is a Tendermint consensus algorithm
func (h Header) ClientType() string { _ = "STUB: not implemented"; return "" }

// GetHeight returns the current height. It returns 0 if the tendermint
// header is nil.
// NOTE: the header.Header is checked to be non nil in ValidateBasic.
func (h Header) GetHeight() exported.Height {
	_ = "STUB: not implemented"
	return *new(exported.Height)
}

// #nosec G115 --- validated before instantiation

// GetTime returns the current block timestamp. It returns a zero time if
// the tendermint header is nil.
// NOTE: the header.Header is checked to be non nil in ValidateBasic.
func (h Header) GetTime() time.Time {
	_ = "STUB: not implemented"
	return *

	// ValidateBasic calls the SignedHeader ValidateBasic function and checks
	// that validatorsets are not nil.
	// NOTE: TrustedHeight and TrustedValidators may be empty when creating client
	// with MsgCreateClient
	new(time.Time)
}

func (h Header) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// TrustedHeight is less than Header for updates and misbehaviour
