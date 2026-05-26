package types

import (
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

var _ exported.Header = &Header{}

// ClientType defines that the Header is a Solo Machine.
func (Header) ClientType() string { _ = "STUB: not implemented"; return "" }

// GetHeight returns the current sequence number as the height.
// Return clientexported.Height to satisfy interface
// Revision number is always 0 for a solo-machine
func (h Header) GetHeight() exported.Height {
	_ = "STUB: not implemented"
	return *new(exported.Height)
}

// GetPubKey unmarshals the new public key into a cryptotypes.PubKey type.
// An error is returned if the new public key is nil or the cached value
// is not a PubKey.
func (h Header) GetPubKey() (cryptotypes.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(cryptotypes.PubKey), nil
}

// ValidateBasic ensures that the sequence, signature and public key have all
// been initialized.
func (h Header) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
