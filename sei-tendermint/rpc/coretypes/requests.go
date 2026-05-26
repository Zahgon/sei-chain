package coretypes

import (
	"encoding/json"
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/bytes"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

type RequestSubscribe struct {
	Query string `json:"query"`
}

type RequestUnsubscribe struct {
	Query string `json:"query"`
}

type RequestBlockchainInfo struct {
	MinHeight Int64 `json:"minHeight"`
	MaxHeight Int64 `json:"maxHeight"`
}

type RequestGenesisChunked struct {
	Chunk Int64 `json:"chunk"`
}

type RequestBlockInfo struct {
	Height *Int64 `json:"height"`
}

type RequestBlockByHash struct {
	Hash bytes.HexBytes `json:"hash"`
}

type RequestCheckTx struct {
	Tx types.Tx `json:"tx"`
}

type RequestRemoveTx struct {
	TxHash types.TxHash `json:"txkey"`
}

type RequestTx struct {
	Hash  bytes.HexBytes `json:"hash"`
	Prove bool           `json:"prove"`
}

type RequestTxSearch struct {
	Query   string `json:"query"`
	Prove   bool   `json:"prove"`
	Page    *Int64 `json:"page"`
	PerPage *Int64 `json:"per_page"`
	OrderBy string `json:"order_by"`
}

type RequestBlockSearch struct {
	Query   string `json:"query"`
	Page    *Int64 `json:"page"`
	PerPage *Int64 `json:"per_page"`
	OrderBy string `json:"order_by"`
}

type RequestValidators struct {
	Height  *Int64 `json:"height"`
	Page    *Int64 `json:"page"`
	PerPage *Int64 `json:"per_page"`
}

type RequestConsensusParams struct {
	Height *Int64 `json:"height"`
}

type RequestUnconfirmedTxs struct {
	Page    *Int64 `json:"page"`
	PerPage *Int64 `json:"per_page"`
}

type RequestBroadcastTx struct {
	Tx types.Tx `json:"tx"`
}

type RequestABCIQuery struct {
	Path   string         `json:"path"`
	Data   bytes.HexBytes `json:"data"`
	Height Int64          `json:"height"`
	Prove  bool           `json:"prove"`
}

type RequestBroadcastEvidence struct {
	Evidence types.Evidence
}

type requestBroadcastEvidenceJSON struct {
	Evidence json.RawMessage `json:"evidence"`
}

func (r RequestBroadcastEvidence) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RequestBroadcastEvidence) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// RequestEvents is the argument for the "/events" RPC endpoint.
type RequestEvents struct {
	// Optional filter spec. If nil or empty, all items are eligible.
	Filter *EventFilter `json:"filter"`

	// The maximum number of eligible items to return.
	// If zero or negative, the server will report a default number.
	MaxItems int `json:"maxItems"`

	// Return only items after this cursor. If empty, the limit is just
	// before the the beginning of the event log.
	After string `json:"after"`

	// Return only items before this cursor.  If empty, the limit is just
	// after the head of the event log.
	Before string `json:"before"`

	// Wait for up to this long for events to be available.
	WaitTime time.Duration `json:"waitTime"`
}

// An EventFilter specifies which events are selected by an /events request.
type EventFilter struct {
	Query string `json:"query"`
}

// Int64 is a wrapper for int64 that encodes to JSON as a string and can be
// decoded from either a string or a number value.
type Int64 int64

func (z *Int64) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (z Int64) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// IntPtr returns a pointer to the value of *z as an int, or nil if z == nil.
func (z *Int64) IntPtr() *int { _ = "STUB: not implemented"; return nil }

// Int64Ptr returns an *Int64 that points to the same value as v, or nil.
func Int64Ptr(v *int) *Int64 { _ = "STUB: not implemented"; return nil }
