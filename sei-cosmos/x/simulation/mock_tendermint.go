package simulation

import (
	"math/rand"
	"testing"
	"time"

	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	tmbytes "github.com/sei-protocol/sei-chain/sei-tendermint/libs/bytes"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
)

type mockValidator struct {
	val           abci.ValidatorUpdate
	livenessState int
}

func (mv mockValidator) String() string { _ = "STUB: not implemented"; return "" }

type mockValidators map[string]mockValidator

// get mockValidators from abci validators
func newMockValidators(r *rand.Rand, abciVals []abci.ValidatorUpdate, params Params) mockValidators {
	_ = "STUB: not implemented"
	return *new(mockValidators)
}

// TODO describe usage
func (vals mockValidators) getKeys() []string { _ = "STUB: not implemented"; return nil }

// randomProposer picks a random proposer from the current validator set
func (vals mockValidators) randomProposer(r *rand.Rand) tmbytes.HexBytes {
	_ = "STUB: not implemented"
	return *new(tmbytes.HexBytes)
}

//nolint:wsl

// updateValidators mimics Tendermint's update logic.
func updateValidators(
	tb testing.TB,
	r *rand.Rand,
	params Params,
	current map[string]mockValidator,
	updates []abci.ValidatorUpdate,
	event func(route, op, evResult string),
) map[string]mockValidator {
	_ = "STUB: not implemented"
	return nil
}

// validator already exists

// Set this new validator

// RandomRequestBeginBlock generates a list of signing validators according to
// the provided list of validators, signing fraction, and evidence fraction
func RandomRequestBeginBlock(r *rand.Rand, params Params,
	validators mockValidators, pastTimes []time.Time,
	pastVoteInfos [][]abci.VoteInfo,
	event func(route, op, evResult string), header tmproto.Header) abci.RequestBeginBlock {
	_ = "STUB: not implemented"
	return *new(abci.RequestBeginBlock)
}

// spotty connection, 50% probability of success
// See https://github.com/golang/go/issues/23804#issuecomment-365370418
// for reasoning behind computing like this

// offline

// return if no past times

// TODO: Determine capacity before allocation

// Tendermint starts at height 1
// array indices offset by one
