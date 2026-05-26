package simulation

import (
	"math/rand"
	"testing"

	"github.com/sei-protocol/sei-chain/sei-cosmos/baseapp"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	simtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
)

const DefaultGenTxGas = 10000000

func getTestingMode(tb testing.TB) (testingMode bool, t *testing.T, b *testing.B) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

// getBlockSize returns a block size as determined from the transition matrix.
// It targets making average block size the provided parameter. The three
// states it moves between are:
//   - "over stuffed" blocks with average size of 2 * avgblocksize,
//   - normal sized blocks, hitting avgBlocksize on average,
//   - and empty blocks, with no txs / only txs scheduled from the past.
func getBlockSize(r *rand.Rand, params Params, lastBlockSizeState, avgBlockSize int) (state, blockSize int) {
	_ = "STUB: not implemented"
	// TODO: Make default blocksize transition matrix actually make the average
	// blocksize equal to avgBlockSize.
	return 0, 0
}

func mustMarshalJSONIndent(o interface{}) []byte { _ = "STUB: not implemented"; return nil }

// OperationInput is a struct that holds all the needed values to generate a tx and deliver it
type OperationInput struct {
	R               *rand.Rand
	App             *baseapp.BaseApp
	TxGen           client.TxConfig
	Cdc             *codec.ProtoCodec
	Msg             sdk.Msg
	MsgType         string
	CoinsSpentInMsg sdk.Coins
	Context         sdk.Context
	SimAccount      simtypes.Account
	AccountKeeper   AccountKeeper
	Bankkeeper      BankKeeper
	ModuleName      string
}

// GenAndDeliverTxWithRandFees generates a transaction with a random fee and delivers it.
func GenAndDeliverTxWithRandFees(txCtx OperationInput) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
	_ = "STUB: not implemented"
	return *new(simtypes.OperationMsg), nil, nil
}

// GenAndDeliverTx generates a transactions and delivers it.
func GenAndDeliverTx(txCtx OperationInput, fees sdk.Coins) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
	_ = "STUB: not implemented"
	return *new(simtypes.OperationMsg), nil, nil
}

func GenTx(gen client.TxConfig, msgs []sdk.Msg, feeAmt sdk.Coins, gas uint64, chainID string, accNums, accSeqs []uint64, priv ...cryptotypes.PrivKey) (sdk.Tx, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Tx), nil
}

// create a random length memo

// 1st round: set SignatureV2 with empty signatures, to set correct
// signer infos.

// 2nd round: once all signer infos are set, every signer can sign.
