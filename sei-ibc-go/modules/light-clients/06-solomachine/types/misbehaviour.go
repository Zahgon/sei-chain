package types

import (
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

var _ exported.Misbehaviour = &Misbehaviour{}

// ClientType is a Solo Machine light client.
func (misbehaviour Misbehaviour) ClientType() string { _ = "STUB: not implemented"; return "" }

// GetClientID returns the ID of the client that committed a misbehaviour.
func (misbehaviour Misbehaviour) GetClientID() string { _ = "STUB: not implemented"; return "" }

// Type implements Evidence interface.
func (misbehaviour Misbehaviour) Type() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic implements Evidence interface.
func (misbehaviour Misbehaviour) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// misbehaviour signatures cannot be identical

// message data signed cannot be identical

// ValidateBasic ensures that the signature and data fields are non-empty.
func (sd SignatureAndData) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
