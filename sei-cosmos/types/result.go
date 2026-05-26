package types

import (
	"github.com/gogo/protobuf/proto"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	ctypes "github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
)

var cdc = codec.NewLegacyAmino()

func (gi GasInfo) String() string { _ = "STUB: not implemented"; return "" }

func (r Result) String() string { _ = "STUB: not implemented"; return "" }

func (r Result) GetEvents() Events { _ = "STUB: not implemented"; return *new(Events) }

// ABCIMessageLogs represents a slice of ABCIMessageLog.
type ABCIMessageLogs []ABCIMessageLog

func NewABCIMessageLog(i uint32, log string, events Events) ABCIMessageLog {
	_ = "STUB: not implemented"
	return *new(ABCIMessageLog)
}

// String implements the fmt.Stringer interface for the ABCIMessageLogs type.
func (logs ABCIMessageLogs) String() (str string) { _ = "STUB: not implemented"; return "" }

// NewResponseResultTx returns a TxResponse given a ResultTx from tendermint
func NewResponseResultTx(res *ctypes.ResultTx, anyTx *codectypes.Any, timestamp string) *TxResponse {
	_ = "STUB: not implemented"
	return nil
}

// NewResponseFormatBroadcastTxCommit returns a TxResponse given a
// ResultBroadcastTxCommit from tendermint.
func NewResponseFormatBroadcastTxCommit(res *ctypes.ResultBroadcastTxCommit) *TxResponse {
	_ = "STUB: not implemented"
	return nil
}

func newTxResponseCheckTx(res *ctypes.ResultBroadcastTxCommit) *TxResponse {
	_ = "STUB: not implemented"
	return nil
}

func newTxResponseDeliverTx(res *ctypes.ResultBroadcastTxCommit) *TxResponse {
	_ = "STUB: not implemented"
	return nil
}

// NewResponseFormatBroadcastTx returns a TxResponse given a ResultBroadcastTx from tendermint
func NewResponseFormatBroadcastTx(res *ctypes.ResultBroadcastTx) *TxResponse {
	_ = "STUB: not implemented"
	return nil
}

func (r TxResponse) String() string { _ = "STUB: not implemented"; return "" }

// Empty returns true if the response is empty
func (r TxResponse) Empty() bool { _ = "STUB: not implemented"; return false }

func NewSearchTxsResult(totalCount, count, page, limit uint64, txs []*TxResponse) *SearchTxsResult {
	_ = "STUB: not implemented"
	return nil
}

// ParseABCILogs attempts to parse a stringified ABCI tx log into a slice of
// ABCIMessageLog types. It returns an error upon JSON decoding failure.
func ParseABCILogs(logs string) (res ABCIMessageLogs, err error) {
	_ = "STUB: not implemented"
	return *new(ABCIMessageLogs), nil
}

var _, _ codectypes.UnpackInterfacesMessage = SearchTxsResult{}, TxResponse{}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
//
// types.UnpackInterfaces needs to be called for each nested Tx because
// there are generally interfaces to unpack in Tx's
func (s SearchTxsResult) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (r TxResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// GetTx unpacks the Tx from within a TxResponse and returns it
func (r TxResponse) GetTx() Tx { _ = "STUB: not implemented"; return *new(Tx) }

type ResultDecorator interface {
	DecorateSdkResult(*Result)
}

// WrapServiceResult wraps a result from a protobuf RPC service method call in
// a Result object or error. This method takes care of marshaling the res param to
// protobuf and attaching any events on the ctx.EventManager() to the Result.
func WrapServiceResult(ctx Context, res proto.Message, err error) (*Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type DeliverTxHookInput struct {
	EvmTxInfo *abci.EvmTxInfo
	Events    []abci.Event
}
