package client

import (
	"context"
	"net"
	"net/http"
	"net/url"
	"sync"

	rpctypes "github.com/sei-protocol/sei-chain/sei-tendermint/rpc/jsonrpc/types"
)

const (
	protoHTTP  = "http"
	protoHTTPS = "https"
	protoWSS   = "wss"
	protoWS    = "ws"
	protoTCP   = "tcp"
	protoUNIX  = "unix"
)

//-------------------------------------------------------------

// Parsed URL structure
type parsedURL struct {
	url.URL

	isUnixSocket bool
}

// Parse URL and set defaults
func newParsedURL(remoteAddr string) (*parsedURL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// default to tcp if nothing specified

// Change protocol to HTTP for unknown protocols and TCP protocol - useful for RPC connections
func (u *parsedURL) SetDefaultSchemeHTTP() {
	_ = "STUB: not implemented"
	// protocol to use for http operations, to support both http and https
	return
}

// known protocols not changed

// default to http for unknown protocols (ex. tcp)

// Get full address without the protocol - useful for Dialer connections
func (u parsedURL) GetHostWithPath() string {
	_ = "STUB: not implemented"
	// Remove protocol, userinfo and # fragment, assume opaque is empty
	return ""
}

// Get a trimmed address - useful for WS connections
func (u parsedURL) GetTrimmedHostWithPath() string {
	_ = "STUB: not implemented"
	// if it's not an unix socket we return the normal URL
	return ""
}

// if it's a unix socket we replace the host slashes with a period
// this is because otherwise the http.Client would think that the
// domain is invalid.

// GetDialAddress returns the endpoint to dial for the parsed URL
func (u parsedURL) GetDialAddress() string {
	_ = "STUB: not implemented"
	// if it's not a unix socket we return the host, example: localhost:443
	return ""
}

// otherwise we return the path of the unix socket, ex /tmp/socket

// Get a trimmed address with protocol - useful as address in RPC connections
func (u parsedURL) GetTrimmedURL() string { _ = "STUB: not implemented"; return "" }

//-------------------------------------------------------------

// A Caller handles the round trip of a single JSON-RPC request.  The
// implementation is responsible for assigning request IDs, marshaling
// parameters, and unmarshaling results.
type Caller interface {
	// Call sends a new request for method to the server with the given
	// parameters. If params == nil, the request has empty parameters.
	// If result == nil, any result value must be discarded without error.
	// Otherwise the concrete value of result must be a pointer.
	Call(ctx context.Context, method string, params, result interface{}) error
}

//-------------------------------------------------------------

// Client is a JSON-RPC client, which sends POST HTTP requests to the
// remote server.
//
// Client is safe for concurrent use by multiple goroutines.
type Client struct {
	address  string
	username string
	password string

	client *http.Client

	mtx       sync.Mutex
	nextReqID int
}

// Both Client and RequestBatch can facilitate calls to the JSON
// RPC endpoint.
var _ Caller = (*Client)(nil)
var _ Caller = (*RequestBatch)(nil)

// New returns a Client pointed at the given address.
// An error is returned on invalid remote. The function panics when remote is nil.
func New(remote string) (*Client, error) { _ = "STUB: not implemented"; return nil, nil }

// NewWithHTTPClient returns a Client pointed at the given address using a
// custom HTTP client. It reports an error if c == nil or if remote is not a
// valid URL.
func NewWithHTTPClient(remote string, c *http.Client) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Call issues a POST HTTP request. Requests are JSON encoded. Content-Type:
// application/json.
func (c *Client) Call(ctx context.Context, method string, params, result interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// NewRequestBatch starts a batch of requests for this client.
func (c *Client) NewRequestBatch() *RequestBatch { _ = "STUB: not implemented"; return nil }

func (c *Client) sendBatch(ctx context.Context, requests []*jsonRPCBufferedRequest) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// serialize the array of requests into a single JSON object

// collect ids to check responses IDs in unmarshalResponseBytesArray

func (c *Client) nextRequestID() int { _ = "STUB: not implemented"; return 0 }

//------------------------------------------------------------------------------------

// jsonRPCBufferedRequest encapsulates a single buffered request, as well as its
// anticipated response structure.
type jsonRPCBufferedRequest struct {
	request rpctypes.RPCRequest
	result  interface{} // The result will be deserialized into this object.
}

// RequestBatch allows us to buffer multiple request/response structures
// into a single batch request. Note that this batch acts like a FIFO queue, and
// is thread-safe.
type RequestBatch struct {
	client *Client

	mtx      sync.Mutex
	requests []*jsonRPCBufferedRequest
}

// Count returns the number of enqueued requests waiting to be sent.
func (b *RequestBatch) Count() int { _ = "STUB: not implemented"; return 0 }

func (b *RequestBatch) enqueue(req *jsonRPCBufferedRequest) { _ = "STUB: not implemented"; return }

// Clear empties out the request batch.
func (b *RequestBatch) Clear() int { _ = "STUB: not implemented"; return 0 }

func (b *RequestBatch) clear() int { _ = "STUB: not implemented"; return 0 }

// Send will attempt to send the current batch of enqueued requests, and then
// will clear out the requests once done. On success, this returns the
// deserialized list of results from each of the enqueued requests.
func (b *RequestBatch) Send(ctx context.Context) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Call enqueues a request to call the given RPC method with the specified
// parameters, in the same way that the `Client.Call` function would.
func (b *RequestBatch) Call(_ context.Context, method string, params, result interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

//-------------------------------------------------------------

func makeHTTPDialer(remoteAddr string) (func(string, string) (net.Conn, error), error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// accept http(s) as an alias for tcp

// DefaultHTTPClient is used to create an http client with some default parameters.
// We overwrite the http.Client.Dial so we can do http over tcp or unix.
// remoteAddr should be fully featured (eg. with tcp:// or unix://).
// An error will be returned in case of invalid remoteAddr.
func DefaultHTTPClient(remoteAddr string) (*http.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set to true to prevent GZIP-bomb DoS attacks
