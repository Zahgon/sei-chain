package testslashing

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/slashing/types"
	abcitypes "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/bytes"
)

// TestParams construct default slashing params for tests.
// Have to change these parameters for tests
// lest the tests take forever
func TestParams() types.Params { _ = "STUB: not implemented"; return *new(types.Params) }

func CreateBeginBlockReq(valAddr bytes.HexBytes, power int64, signed bool) abcitypes.RequestBeginBlock {
	_ = "STUB: not implemented"
	return *new(abcitypes.RequestBeginBlock)
}
