package evmrpc

import (
	"context"
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/sei-protocol/sei-chain/x/evm/types"
)

type contextKey string

const tendermintTraceKey contextKey = "tendermintTrace"
const receiptTraceKey contextKey = "receiptTrace"

type TendermintTraces struct {
	Traces []TendermintTrace `json:"traces"`
}

func (tt *TendermintTraces) MustMarshalToJson() json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

type ReceiptTraces struct {
	Traces []RawResponseReceipt `json:"traces"`
}

func (rt *ReceiptTraces) MustMarshalToJson() json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

type RawResponseReceipt struct {
	BlockNumber       hexutil.Uint64  `json:"blockNumber"`
	ContractAddress   *common.Address `json:"contractAddress"`
	CumulativeGasUsed hexutil.Uint64  `json:"cumulativeGasUsed"`
	EffectiveGasPrice *hexutil.Big    `json:"effectiveGasPrice"`
	From              common.Address  `json:"from"`
	To                *common.Address `json:"to"`
	GasUsed           hexutil.Uint64  `json:"gasUsed"`
	Status            hexutil.Uint    `json:"status"`
	Type              hexutil.Uint    `json:"type"`
	TransactionHash   common.Hash     `json:"transactionHash"`
	TransactionIndex  hexutil.Uint64  `json:"transactionIndex"`
}

type TendermintTrace struct {
	Endpoint  string          `json:"endpoint"`
	Arguments []string        `json:"arguments"`
	Response  json.RawMessage `json:"response"`
}

func WithTendermintTraces(ctx context.Context, traces *TendermintTraces) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func TraceTendermintIfApplicable(ctx context.Context, endpoint string, arguments []string, response interface{}) {
	_ = "STUB: not implemented"
	return
}

func TendermintTracesFromContext(ctx context.Context) *TendermintTraces {
	_ = "STUB: not implemented"
	return nil
}

func WithReceiptTraces(ctx context.Context, traces *ReceiptTraces) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func TraceReceiptIfApplicable(ctx context.Context, receipt *types.Receipt) {
	_ = "STUB: not implemented"
	return
}

func stringifyInt64Ptr(i *int64) string { _ = "STUB: not implemented"; return "" }
