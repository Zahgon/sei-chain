package evmrpc

import (
	"encoding/json"
	"net/http"
)

// seiLegacyHTTPMaxBody matches github.com/ethereum/go-ethereum/rpc.defaultBodyLimit (5MiB), the
// default HTTP request body cap used by rpc.Server before ServeHTTP. The legacy gate must not
// read more than the inner JSON-RPC stack will accept (see rpc.Server.SetHTTPBodyLimit).
const seiLegacyHTTPMaxBody = 5 * 1024 * 1024

const (
	invalidRequestCode  = -32600
	seiLegacyNotEnabled = -32601
	internalErrorCode   = -32603
)

// wrapSeiLegacyHTTP wraps the EVM JSON-RPC HTTP handler to enforce [evm].enabled_legacy_sei_apis for
// gated sei_* and sei2_* methods. Disallowed calls get a JSON-RPC error without invoking the inner handler.
// Single-object allowed calls pass through unchanged; batches forward a filtered subset and merge inner
// results back by JSON-RPC id. Deprecation header on successful forwards of gated methods. nil allowlist = no wrap.
func wrapSeiLegacyHTTP(inner http.Handler, allowlist map[string]struct{}) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

type seiLegacyHTTPGate struct {
	inner     http.Handler
	allowlist map[string]struct{}
}

func (g *seiLegacyHTTPGate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	// Read the body once; delegate JSON-RPC validation to the inner handler. We only intercept
	// when we can parse JSON-RPC and the method is a gated sei_* / sei2_* name.
	return
}

func (g *seiLegacyHTTPGate) serveInnerWithBody(w http.ResponseWriter, r *http.Request, body []byte) {
	_ = "STUB: not implemented"
	return
}

func orNullID(id json.RawMessage) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

func (g *seiLegacyHTTPGate) handleSingle(w http.ResponseWriter, r *http.Request, body []byte) {
	_ = "STUB: not implemented"
	return
}

// Non-gated methods (eth_*, web3_*, net_*, etc.) need no recording or
// header injection — pass straight through so the gzip handler writes
// directly to the real http.ResponseWriter.

// Prevent the inner gzip handler from compressing into the recorder;
// we need plain JSON so copyHTTPHeader does not propagate a stale
// Content-Encoding: gzip for a body that is replayed uncompressed.

type jsonrpcMessage struct {
	Method string          `json:"method"`
	ID     json.RawMessage `json:"id"`
}

// logic taken from hasValidID in https://github.com/ethereum/go-ethereum/blob/master/rpc/json.go
// null is valid: it is used in error responses per JSON-RPC 2.0 §5
func (m *jsonrpcMessage) hasValidID() bool { _ = "STUB: not implemented"; return false }

func (g *seiLegacyHTTPGate) handleBatch(w http.ResponseWriter, r *http.Request, body []byte) {
	_ = "STUB: not implemented"
	return
}

// Batch element is not a JSON object, or has an invalid (object/array) id; synthesize -32600 and do not forward.

// Fast path: every element is forwarded (nothing blocked/invalid) and none
// are gated sei_*/sei2_* methods.  Skip the recorder so the gzip handler
// writes directly to the real http.ResponseWriter — same fix as handleSingle.

// Prevent the inner gzip handler from compressing into the recorder;
// mergeSeiLegacyHTTPBatch needs plain JSON to unmarshal inner results.

const seiLegacyBatchInternalErr = "invalid or incomplete JSON-RPC batch response from server"

const seiLegacyBatchInvalidReqMsg = "Invalid Request"

// seiLegacyBatchResponsesNoForward builds the batch JSON-RPC response when nothing is forwarded to the
// inner server: invalid slots yield -32600, blocked gated methods yield gate errors, and notifications
// (requests with no "id" member) are omitted (JSON-RPC 2.0: no response for notifications, including in batches).
func seiLegacyBatchResponsesNoForward(
	invalidReq []bool,
	blockedErr []error,
	ids []json.RawMessage,
	lenMsgs int,
) []json.RawMessage {
	_ = "STUB: not implemented"
	return nil
}

// mergeSeiLegacyHTTPBatch merges inner batch results with gate/invalid slots. Output is ordered like the
// original batch but omits entries for JSON-RPC notifications (no "id" member), per JSON-RPC 2.0 batch rules.
// synthIDs holds the unique synthetic ID assigned to each forwarded non-notification request (nil for all
// others). The inner server echoes these synthetic IDs back, so idToIdx is always collision-free regardless
// of duplicate or null original IDs. Original IDs are restored in the output via patchJSONRPCResponseIDIfNeeded.
func mergeSeiLegacyHTTPBatch(
	invalidReq []bool,
	blocked []bool,
	blockedErr []error,
	ids []json.RawMessage,
	synthIDs []json.RawMessage,
	lenMsgs int,
	innerBody []byte,
) []json.RawMessage {
	_ = "STUB: not implemented"
	return nil
}

// Skip malformed inner entry; the matching slot will fall through to
// the internalErrorCode branch in the merge loop below.

func rpcIDKey(id json.RawMessage) string { _ = "STUB: not implemented"; return "" }

// isJSONRPCNotificationID is true when the request omits "id" (JSON-RPC Notification).
// "id": null is discouraged but valid per JSON-RPC 2.0 and MUST receive a response like any other id.
func isJSONRPCNotificationID(id json.RawMessage) bool { _ = "STUB: not implemented"; return false }

func jsonRPCObjectIDKey(raw json.RawMessage) (idField json.RawMessage, hasKey bool, err error) {
	_ = "STUB: not implemented"
	return *new(json.RawMessage), false, nil
}

func marshalJSONRPCError(id json.RawMessage, code int, message string) []byte {
	_ = "STUB: not implemented"
	return nil
}

// patchJSONRPCResponseIDIfNeeded replaces response "id" with the client's id as raw JSON (avoids float64 rounding).
func patchJSONRPCResponseIDIfNeeded(resp json.RawMessage, wantID json.RawMessage) []byte {
	_ = "STUB: not implemented"
	return nil
}

func copyHTTPHeader(dst, src http.Header) { _ = "STUB: not implemented"; return }

// writeJSONRPCBatchResponse writes a JSON-RPC batch response. Per JSON-RPC 2.0, if there are no
// response objects, the server must not return an empty JSON array — use an empty HTTP body instead.
func writeJSONRPCBatchResponse(w http.ResponseWriter, code int, arr []json.RawMessage) {
	_ = "STUB: not implemented"
	return
}

// setJSONObjectID returns obj with its "id" field replaced by newID.
// Returns obj unchanged if it cannot be parsed as a JSON object.
func setJSONObjectID(obj json.RawMessage, newID json.RawMessage) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

func writeJSONArrayResponse(w http.ResponseWriter, code int, arr []json.RawMessage) {
	_ = "STUB: not implemented"
	return
}

func writeSeiLegacyBlocked(w http.ResponseWriter, id json.RawMessage, gateErr error) {
	_ = "STUB: not implemented"
	return
}

func marshalBlockedResponse(id json.RawMessage, gateErr error) []byte {
	_ = "STUB: not implemented"
	return nil
}
