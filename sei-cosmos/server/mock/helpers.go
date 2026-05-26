package mock

import (
	"testing"

	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
)

// SetupApp returns an application as well as a clean-up function
// to be used to quickly setup a test case with an app
func SetupApp(t *testing.T) abci.Application {
	_ = "STUB: not implemented"
	return *new(abci.Application)
}
