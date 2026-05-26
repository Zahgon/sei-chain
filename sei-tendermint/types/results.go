package types

import (
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto/merkle"
)

// ABCIResults wraps the deliver tx results to return a proof.
type ABCIResults []*abci.ResponseDeliverTx

// NewResults strips non-deterministic fields from ResponseDeliverTx responses
// and returns ABCIResults.
func NewResults(responses []*abci.ResponseDeliverTx) ABCIResults {
	_ = "STUB: not implemented"
	return *new(ABCIResults)
}

func (a ABCIResults) Hash() []byte { _ = "STUB: not implemented"; return nil }

func (a ABCIResults) ProveResult(i int) merkle.Proof {
	_ = "STUB: not implemented"
	return *new(merkle.Proof)
}

func (a ABCIResults) toByteSlices() [][]byte { _ = "STUB: not implemented"; return nil }

func deterministicResponseDeliverTx(response *abci.ResponseDeliverTx) *abci.ResponseDeliverTx {
	_ = "STUB: not implemented"
	return nil
}
