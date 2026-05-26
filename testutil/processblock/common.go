package processblock

import (
	"testing"

	"github.com/sei-protocol/sei-chain/app"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/signing"
	stakingtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
)

type App struct {
	*app.App

	height        int64
	proposer      int
	accToMnemonic map[string]string
	accToSeqDelta map[string]uint64
	lastCtx       sdk.Context
}

func NewTestApp(t *testing.T) *App { _ = "STUB: not implemented"; return nil }

func (a *App) Ctx() sdk.Context {
	_ = "STUB: not implemented"

	// Processes and commits a block of transactions, and return a list of response codes.
	// Assumes all validators voted with equal weight, and there are no byzantine validators.
	// Proposer is rotated among all validators round-robin.
	return *new(sdk.Context)
}

func (a *App) RunBlock(txs []signing.Tx) (resultCodes []uint32) {
	_ = "STUB: not implemented"
	return nil
}

// Commit will set deliver tx ctx to nil so we need to cache it here for testing queries before the next block is FinalizeBlock'ed (which will set deliver tx ctx)

// no needed for application logic

func (a *App) GetVotes() []types.VoteInfo { _ = "STUB: not implemented"; return nil }

func (a *App) GetAllValidators() []stakingtypes.Validator { _ = "STUB: not implemented"; return nil }

func (a *App) GetProposer() stakingtypes.Validator {
	_ = "STUB: not implemented"
	return *new(stakingtypes.Validator)
}

func (a *App) GenerateSignableKey(_ string) (addr sdk.AccAddress) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress)
}

func GenerateRandomPubKey() cryptotypes.PubKey {
	_ = "STUB: not implemented"
	return *new(cryptotypes.PubKey)
}

func generateRandomStringOfLength(len int) string { _ = "STUB: not implemented"; return "" }

func getValAddress(v stakingtypes.Validator) []byte { _ = "STUB: not implemented"; return nil }
