package types

import (
	"context"
	"encoding/json"
	"net/http"
	"regexp"
)

// ErrorCode is the type of JSON-RPC error codes.
type ErrorCode int

func (e ErrorCode) String() string { _ = "STUB: not implemented"; return "" }

// Constants defining the standard JSON-RPC error codes.
const (
	CodeParseError     ErrorCode = -32700 // Invalid JSON received by the server
	CodeInvalidRequest ErrorCode = -32600 // The JSON sent is not a valid request object
	CodeMethodNotFound ErrorCode = -32601 // The method does not exist or is unavailable
	CodeInvalidParams  ErrorCode = -32602 // Invalid method parameters
	CodeInternalError  ErrorCode = -32603 // Internal JSON-RPC error
	CodeLagIsHighError ErrorCode = -32604 // Lag is too high error
)

var errorCodeString = map[ErrorCode]string{
	CodeParseError:     "Parse error",
	CodeInvalidRequest: "Invalid request",
	CodeMethodNotFound: "Method not found",
	CodeInvalidParams:  "Invalid params",
	CodeInternalError:  "Internal error",
	CodeLagIsHighError: "Lag is too high",
}

//----------------------------------------
// REQUEST

type RPCRequest struct {
	id json.RawMessage

	Method string
	Params json.RawMessage
}

// NewRequest returns an empty request with the specified ID.
func NewRequest(id int) RPCRequest { _ = "STUB: not implemented"; return *new(RPCRequest) }

// ID returns a string representation of the request ID.
func (req RPCRequest) ID() string { _ = "STUB: not implemented"; return "" }

// IsNotification reports whether req is a notification (has an empty ID).
func (req RPCRequest) IsNotification() bool { _ = "STUB: not implemented"; return false }

type rpcRequestJSON struct {
	V  string          `json:"jsonrpc"` // must be "2.0"
	ID json.RawMessage `json:"id,omitempty"`
	M  string          `json:"method"`
	P  json.RawMessage `json:"params"`
}

// isNullOrEmpty reports whether data is empty or the JSON "null" value.
func isNullOrEmpty(data json.RawMessage) bool { _ = "STUB: not implemented"; return false }

// validID matches the text of a JSON value that is allowed to serve as a
// JSON-RPC request ID. Precondition: Target value is legal JSON.
var validID = regexp.MustCompile(`^(?:".*"|-?\d+)$`)

// UnmarshalJSON decodes a request from a JSON-RPC 2.0 request object.
func (req *RPCRequest) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON marshals a request with the appropriate version tag.
func (req RPCRequest) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (req RPCRequest) String() string { _ = "STUB: not implemented"; return "" }

// MakeResponse constructs a success response to req with the given result.  If
// there is an error marshaling result to JSON, it returns an error response.
func (req RPCRequest) MakeResponse(result interface{}) RPCResponse {
	_ = "STUB: not implemented"
	return *new(RPCResponse)
}

// MakeErrorf constructs an error response to req with the given code and a
// message constructed by formatting msg with args.
func (req RPCRequest) MakeErrorf(code ErrorCode, msg string, args ...interface{}) RPCResponse {
	_ = "STUB: not implemented"
	return *new(RPCResponse)
}

// MakeError constructs an error response to req from the given error value.
// This function will panic if err == nil.
func (req RPCRequest) MakeError(result interface{}, err error) RPCResponse {
	_ = "STUB: not implemented"
	return *new(RPCResponse)
}

// Handle lag is high error specifically to avoid changing the logic for existing endpoints

// SetMethodAndParams updates the method and parameters of req with the given
// values, leaving the ID unchanged.
func (req *RPCRequest) SetMethodAndParams(method string, params interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

//----------------------------------------
// RESPONSE

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data,omitempty"`
}

func (err RPCError) Error() string { _ = "STUB: not implemented"; return "" }

type RPCResponse struct {
	id json.RawMessage

	Result json.RawMessage
	Error  *RPCError
}

// ID returns a representation of the response ID.
func (resp RPCResponse) ID() string { _ = "STUB: not implemented"; return "" }

type rpcResponseJSON struct {
	V  string          `json:"jsonrpc"` // must be "2.0"
	ID json.RawMessage `json:"id,omitempty"`
	R  json.RawMessage `json:"result,omitempty"`
	E  *RPCError       `json:"error,omitempty"`
}

// UnmarshalJSON decodes a response from a JSON-RPC 2.0 response object.
func (resp *RPCResponse) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON marshals a response with the appropriate version tag.
func (resp RPCResponse) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (resp RPCResponse) String() string { _ = "STUB: not implemented"; return "" }

//----------------------------------------

// WSRPCConnection represents a websocket connection.
type WSRPCConnection interface {
	// GetRemoteAddr returns a remote address of the connection.
	GetRemoteAddr() string
	// WriteRPCResponse writes the response onto connection (BLOCKING).
	WriteRPCResponse(context.Context, RPCResponse) error
	// TryWriteRPCResponse tries to write the response onto connection (NON-BLOCKING).
	TryWriteRPCResponse(context.Context, RPCResponse) bool
	// Context returns the connection's context.
	Context() context.Context
}

// CallInfo carries JSON-RPC request metadata for RPC functions invoked via
// JSON-RPC. It can be recovered from the context with GetCallInfo.
type CallInfo struct {
	RPCRequest  *RPCRequest     // non-nil for requests via HTTP or websocket
	HTTPRequest *http.Request   // non-nil for requests via HTTP
	WSConn      WSRPCConnection // non-nil for requests via websocket
}

type callInfoKey struct{}

// WithCallInfo returns a child context of ctx with the ci attached.
func WithCallInfo(ctx context.Context, ci *CallInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// GetCallInfo returns the CallInfo record attached to ctx, or nil if ctx does
// not contain a call record.
func GetCallInfo(ctx context.Context) *CallInfo { _ = "STUB: not implemented"; return nil }

// RemoteAddr returns the remote address (usually a string "IP:port").  If
// neither HTTPRequest nor WSConn is set, an empty string is returned.
//
// For HTTP requests, this reports the request's RemoteAddr.
// For websocket requests, this reports the connection's GetRemoteAddr.
func (ci *CallInfo) RemoteAddr() string { _ = "STUB: not implemented"; return "" }

//----------------------------------------
// SOCKETS

// Determine if its a unix or tcp socket.
// If tcp, must specify the port; `0.0.0.0` will return incorrectly as "unix" since there's no port
// TODO: deprecate
func SocketType(listenAddr string) string { _ = "STUB: not implemented"; return "" }
