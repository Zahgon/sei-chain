package types

import (
	"time"

	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
	tmtypes "github.com/sei-protocol/sei-chain/sei-tendermint/types"

	clienttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/types"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

var _ exported.Misbehaviour = &Misbehaviour{}

// Use the same FrozenHeight for all misbehaviour
var FrozenHeight = clienttypes.NewHeight(0, 1)

// NewMisbehaviour creates a new Misbehaviour instance.
func NewMisbehaviour(clientID string, header1, header2 *Header) *Misbehaviour {
	_ = "STUB: not implemented"
	return nil
}

// ClientType is Tendermint light client
func (misbehaviour Misbehaviour) ClientType() string { _ = "STUB: not implemented"; return "" }

// GetClientID returns the ID of the client that committed a misbehaviour.
func (misbehaviour Misbehaviour) GetClientID() string { _ = "STUB: not implemented"; return "" }

// GetTime returns the timestamp at which misbehaviour occurred. It uses the
// maximum value from both headers to prevent producing an invalid header outside
// of the misbehaviour age range.
func (misbehaviour Misbehaviour) GetTime() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// ValidateBasic implements Misbehaviour interface
func (misbehaviour Misbehaviour) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// ValidateBasic on both validators

// Ensure that Height1 is greater than or equal to Height2

// validCommit checks if the given commit is a valid commit from the passed-in validatorset
func validCommit(chainID string, blockID tmtypes.BlockID, commit *tmproto.Commit, valSet *tmproto.ValidatorSet) (err error) {
	_ = "STUB: not implemented"
	return nil
}
