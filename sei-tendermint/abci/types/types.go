package types

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gogo/protobuf/jsonpb"
)

const (
	CodeTypeOK uint32 = 0
)

// IsOK returns true if Code is OK.
func (r ResponseCheckTx) IsOK() bool { _ = "STUB: not implemented"; return false }

func (r ResponseCheckTx) Err() error { _ = "STUB: not implemented"; return nil }

// IsErr returns true if Code is something other than OK.
func (r ResponseCheckTx) IsErr() bool { _ = "STUB: not implemented"; return false }

// IsOK returns true if Code is OK.
func (r ResponseDeliverTx) IsOK() bool { _ = "STUB: not implemented"; return false }

// IsErr returns true if Code is something other than OK.
func (r ResponseDeliverTx) IsErr() bool { _ = "STUB: not implemented"; return false }

// IsOK returns true if Code is OK.
func (r ExecTxResult) IsOK() bool { _ = "STUB: not implemented"; return false }

// IsErr returns true if Code is something other than OK.
func (r ExecTxResult) IsErr() bool { _ = "STUB: not implemented"; return false }

// IsOK returns true if Code is OK.
func (r ResponseQuery) IsOK() bool { _ = "STUB: not implemented"; return false }

// IsErr returns true if Code is something other than OK.
func (r ResponseQuery) IsErr() bool { _ = "STUB: not implemented"; return false }

func (r ResponseProcessProposal) IsAccepted() bool { _ = "STUB: not implemented"; return false }

func (r ResponseProcessProposal) IsStatusUnknown() bool { _ = "STUB: not implemented"; return false }

//---------------------------------------------------------------------------
// override JSON marshaling so we emit defaults (ie. disable omitempty)

var (
	jsonpbMarshaller = jsonpb.Marshaler{
		EnumsAsInts:  true,
		EmitDefaults: true,
	}
	jsonpbUnmarshaller = jsonpb.Unmarshaler{}
)

func (r *ResponseCheckTx) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *ResponseCheckTx) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (r *ResponseDeliverTx) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ResponseDeliverTx) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (r *ResponseQuery) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *ResponseQuery) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (r *ResponseCommit) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *ResponseCommit) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (r *EventAttribute) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *EventAttribute) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// validatorUpdateJSON is the JSON encoding of a validator update.
//
// It handles translation of public keys from the protobuf representation to
// the legacy Amino-compatible format expected by RPC clients.
type validatorUpdateJSON struct {
	PubKey json.RawMessage `json:"pub_key,omitempty"`
	Power  int64           `json:"power,string"`
}

func (v *ValidatorUpdate) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (v *ValidatorUpdate) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// Some compile time assertions to ensure we don't
// have accidental runtime surprises later on.

// jsonEncodingRoundTripper ensures that asserted
// interfaces implement both MarshalJSON and UnmarshalJSON
type jsonRoundTripper interface {
	json.Marshaler
	json.Unmarshaler
}

var _ jsonRoundTripper = (*ResponseCommit)(nil)
var _ jsonRoundTripper = (*ResponseQuery)(nil)
var _ jsonRoundTripper = (*ResponseDeliverTx)(nil)
var _ jsonRoundTripper = (*ResponseCheckTx)(nil)

var _ jsonRoundTripper = (*EventAttribute)(nil)

// -----------------------------------------------
// construct Result data

// deterministicExecTxResult constructs a copy of response that omits
// non-deterministic fields. The input response is not modified.
func deterministicExecTxResult(response *ExecTxResult) *ExecTxResult {
	_ = "STUB: not implemented"
	return nil
}

// MarshalTxResults encodes the the TxResults as a list of byte
// slices. It strips off the non-deterministic pieces of the TxResults
// so that the resulting data can be used for hash comparisons and used
// in Merkle proofs.
func MarshalTxResults(r []*ExecTxResult) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type PendingTxCheckerResponse int

const (
	Accepted PendingTxCheckerResponse = iota
	Rejected
	Pending
)

// ResponseCheckTxV2 response type contains non-protobuf fields, so non-local ABCI clients will not be able
// to utilize the new fields in V2 type (but still be backwards-compatible)
type ResponseCheckTxV2 struct {
	*ResponseCheckTx

	// helper properties for prioritization in mempool
	EVMNonce uint64
	// EVM and sei addresses are both derived from the sender's public key.
	// TODO(gprusak): include just the secp256k1 public key and let the CheckTx caller derive evm/sei address on their own.
	EVMSenderAddress   common.Address
	SeiSenderAddress   []byte
	IsEVM              bool
	EVMRequiredBalance *big.Int
}

type CheckTxTypeV2 int32

const (
	CheckTxTypeV2New CheckTxTypeV2 = iota
	CheckTxTypeV2Recheck
)

type RequestCheckTxV2 struct {
	Tx   []byte
	Type CheckTxTypeV2
}

type RequestDeliverTxV2 struct {
	Tx          []byte
	SigVerified bool
}

type RequestGetTxPriorityHintV2 struct {
	Tx []byte
}

type TxResultV2 struct {
	Height int64
	Index  uint32
	Tx     []byte
	Result ExecTxResult
}
