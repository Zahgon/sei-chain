package verify

import (
	"testing"

	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/signing"
	"github.com/sei-protocol/sei-chain/testutil/processblock"
)

func MintRelease(t *testing.T, app *processblock.App, f BlockRunnable, _ []signing.Tx) BlockRunnable {
	_ = "STUB: not implemented"
	return *new(BlockRunnable)
}

// if minter minted, it must be a new epoch, but not the other way around

//nolint:gosec

//nolint:gosec
