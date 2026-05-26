package app

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
)

var ERC20ApprovalTopic = common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
var ERC20TransferTopic = common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
var ERC721TransferTopic = common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
var ERC721ApprovalTopic = common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
var ERC721ApproveAllTopic = common.HexToHash("0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31")
var ERC1155TransferSingleTopic = common.HexToHash("0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62")
var ERC1155TransferBatchTopic = common.HexToHash("0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb")
var ERC1155ApprovalForAllTopic = common.HexToHash("0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31")
var ERC1155URITopic = common.HexToHash("0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b")
var EmptyHash = common.HexToHash("0x0")
var TrueHash = common.HexToHash("0x1")

type AllowanceResponse struct {
	Allowance sdk.Int         `json:"allowance"`
	Expires   json.RawMessage `json:"expires"`
}

func getOwnerEventKey(contractAddr string, tokenID string) string {
	_ = "STUB: not implemented"
	return ""
}

func (app *App) AddCosmosEventsToEVMReceiptIfApplicable(ctx sdk.Context, tx sdk.Tx, checksum [32]byte, response sdk.DeliverTxHookInput) {
	_ = "STUB: not implemented"
	// hooks will only be called if DeliverTx is successful
	return
}

// Note: txs with a very large number of WASM events may run out of gas due to
// additional gas consumption from EVM receipt generation and event translation

// unfortunately CW721 transfer events differ from ERC721 transfer events
// in that CW721 include sender (which can be different than owner) whereas
// ERC721 always include owner. The following logic refer to the owner
// event emitted before the transfer and use that instead to populate the
// synthetic ERC721 event.

// check if there is a ERC721 pointer to contract Addr

// check if there is a ERC1155 pointer to contract Addr

//nolint:gosec

//nolint:gosec

//nolint:gosec
//nolint:gosec

// we don't create shell receipt for failed Cosmos tx since there is no event anyway

// use the first signer as the `from`

func (app *App) translateCW20Event(ctx sdk.Context, wasmEvent abci.Event, pointerAddr common.Address, contractAddr string) (res []*ethtypes.Log) {
	_ = "STUB: not implemented"
	return nil
}

func (app *App) translateCW721Event(ctx sdk.Context, wasmEvent abci.Event, pointerAddr common.Address, contractAddr string,
	ownerEventsMap map[string][]abci.Event, cw721TransferCounterMap map[string]int) (res []*ethtypes.Log) {
	_ = "STUB: not implemented"
	return nil
}

func (app *App) translateCW1155Event(ctx sdk.Context, wasmEvent abci.Event, pointerAddr common.Address, contractAddr string) (res []*ethtypes.Log) {
	_ = "STUB: not implemented"
	return nil
}

func (app *App) GetEvmAddressHash(ctx sdk.Context, addrStr string) common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
}

func GetEventsOfType(rdtx sdk.DeliverTxHookInput, ty string) (res []abci.Event) {
	_ = "STUB: not implemented"
	return nil
}

func GetAttributeValue(event abci.Event, attribute string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (app *App) GetActionsFromWasmEvent(ctx sdk.Context, event abci.Event) (actions []*Action) {
	_ = "STUB: not implemented"
	return nil
}

func safeBigIntFromString(s string) *big.Int { _ = "STUB: not implemented"; return nil }

type Action struct {
	Type      string
	Amount    *big.Int
	Amounts   []*big.Int
	TokenId   *big.Int
	TokenIds  []*big.Int
	Sender    common.Hash
	Recipient common.Hash
	Spender   common.Hash
	Operator  common.Hash
	Owner     common.Hash
	From      common.Hash
	To        common.Hash
}
