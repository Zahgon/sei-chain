// Package rest provides HTTP types and primitives for REST
// requests validation and responses handling.
package rest

import (
	"encoding/json"
	"net/http"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const (
	DefaultPage    = 1
	DefaultLimit   = 30             // should be consistent with tendermint/tendermint/rpc/core/pipe.go:19
	TxMinHeightKey = "tx.minheight" // Inclusive minimum height filter
	TxMaxHeightKey = "tx.maxheight" // Inclusive maximum height filter
)

// ResponseWithHeight defines a response object type that wraps an original
// response with a height.
type ResponseWithHeight struct {
	Height int64           `json:"height"`
	Result json.RawMessage `json:"result"`
}

// NewResponseWithHeight creates a new ResponseWithHeight instance
func NewResponseWithHeight(height int64, result json.RawMessage) ResponseWithHeight {
	_ = "STUB: not implemented"
	return *new(ResponseWithHeight)
}

// ParseResponseWithHeight returns the raw result from a JSON-encoded
// ResponseWithHeight object.
func ParseResponseWithHeight(cdc *codec.LegacyAmino, bz []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GasEstimateResponse defines a response definition for tx gas estimation.
type GasEstimateResponse struct {
	GasEstimate uint64 `json:"gas_estimate"`
}

// BaseReq defines a structure that can be embedded in other request structures
// that all share common "base" fields.
type BaseReq struct {
	From          string       `json:"from"`
	Memo          string       `json:"memo"`
	ChainID       string       `json:"chain_id"`
	AccountNumber uint64       `json:"account_number"`
	Sequence      uint64       `json:"sequence"`
	TimeoutHeight uint64       `json:"timeout_height"`
	Fees          sdk.Coins    `json:"fees"`
	GasPrices     sdk.DecCoins `json:"gas_prices"`
	Gas           string       `json:"gas"`
	GasAdjustment string       `json:"gas_adjustment"`
	Simulate      bool         `json:"simulate"`
}

// NewBaseReq creates a new basic request instance and sanitizes its values
func NewBaseReq(
	from, memo, chainID string, gas, gasAdjustment string, accNumber, seq uint64,
	fees sdk.Coins, gasPrices sdk.DecCoins, simulate bool,
) BaseReq {
	_ = "STUB: not implemented"
	return *new(BaseReq)
}

// Sanitize performs basic sanitization on a BaseReq object.
func (br BaseReq) Sanitize() BaseReq { _ = "STUB: not implemented"; return *new(BaseReq) }

// ValidateBasic performs basic validation of a BaseReq. If custom validation
// logic is needed, the implementing request handler should perform those
// checks manually.
func (br BaseReq) ValidateBasic(w http.ResponseWriter) bool {
	_ = "STUB: not implemented"
	return false
}

// both fees and gas prices were provided

// neither fees or gas prices were provided

// ReadRESTReq reads and unmarshals a Request's body to the the BaseReq struct.
// Writes an error response to ResponseWriter and returns false if errors occurred.
func ReadRESTReq(w http.ResponseWriter, r *http.Request, cdc *codec.LegacyAmino, req interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

// ErrorResponse defines the attributes of a JSON error response.
type ErrorResponse struct {
	Code  int    `json:"code,omitempty"`
	Error string `json:"error"`
}

// NewErrorResponse creates a new ErrorResponse instance.
func NewErrorResponse(code int, err string) ErrorResponse {
	_ = "STUB: not implemented"
	return *new(ErrorResponse)
}

// CheckError takes care of writing an error response if err is not nil.
// Returns false when err is nil; it returns true otherwise.
func CheckError(w http.ResponseWriter, status int, err error) bool {
	_ = "STUB: not implemented"
	return false
}

// CheckBadRequestError attaches an error message to an HTTP 400 BAD REQUEST response.
// Returns false when err is nil; it returns true otherwise.
func CheckBadRequestError(w http.ResponseWriter, err error) bool {
	_ = "STUB: not implemented"
	return false
}

// CheckInternalServerError attaches an error message to an HTTP 500 INTERNAL SERVER ERROR response.
// Returns false when err is nil; it returns true otherwise.
func CheckInternalServerError(w http.ResponseWriter, err error) bool {
	_ = "STUB: not implemented"
	return false
}

// CheckNotFoundError attaches an error message to an HTTP 404 NOT FOUND response.
// Returns false when err is nil; it returns true otherwise.
func CheckNotFoundError(w http.ResponseWriter, err error) bool {
	_ = "STUB: not implemented"
	return false
}

// WriteErrorResponse prepares and writes a HTTP error
// given a status code and an error message.
func WriteErrorResponse(w http.ResponseWriter, status int, err string) {
	_ = "STUB: not implemented"
	return
}

// WriteSimulationResponse prepares and writes an HTTP
// response for transactions simulations.
func WriteSimulationResponse(w http.ResponseWriter, cdc *codec.LegacyAmino, gas uint64) {
	_ = "STUB: not implemented"
	return
}

// ParseUint64OrReturnBadRequest converts s to a uint64 value.
func ParseUint64OrReturnBadRequest(w http.ResponseWriter, s string) (n uint64, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// ParseFloat64OrReturnBadRequest converts s to a float64 value. It returns a
// default value, defaultIfEmpty, if the string is empty.
func ParseFloat64OrReturnBadRequest(w http.ResponseWriter, s string, defaultIfEmpty float64) (n float64, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// ParseQueryHeightOrReturnBadRequest sets the height to execute a query if set by the http request.
// It returns false if there was an error parsing the height.
func ParseQueryHeightOrReturnBadRequest(w http.ResponseWriter, clientCtx client.Context, r *http.Request) (client.Context, bool) {
	_ = "STUB: not implemented"
	return *new(client.Context), false
}

// PostProcessResponseBare post processes a body similar to PostProcessResponse
// except it does not wrap the body and inject the height.
func PostProcessResponseBare(w http.ResponseWriter, ctx client.Context, body interface{}) {
	_ = "STUB: not implemented"
	return
}

// PostProcessResponse performs post processing for a REST response. The result
// returned to clients will contain two fields, the height at which the resource
// was queried at and the original result.
func PostProcessResponse(w http.ResponseWriter, ctx client.Context, resp interface{}) {
	_ = "STUB: not implemented"
	return
}

// LegacyAmino used intentionally for REST

// ParseHTTPArgsWithLimit parses the request's URL and returns a slice containing
// all arguments pairs. It separates page and limit used for pagination where a
// default limit can be provided.
func ParseHTTPArgsWithLimit(r *http.Request, defaultLimit int) (tags []string, page, limit int, err error) {
	_ = "STUB: not implemented"
	return nil, 0, 0, nil
}

// ParseHTTPArgs parses the request's URL and returns a slice containing all
// arguments pairs. It separates page and limit used for pagination.
func ParseHTTPArgs(r *http.Request) (tags []string, page, limit int, err error) {
	_ = "STUB: not implemented"
	return nil, 0, 0, nil
}

// ParseQueryParamBool parses the given param to a boolean. It returns false by
// default if the string is not parseable to bool.
func ParseQueryParamBool(r *http.Request, param string) bool {
	_ = "STUB: not implemented"
	return false
}

// GetRequest defines a wrapper around an HTTP GET request with a provided URL.
// An error is returned if the request or reading the body fails.
func GetRequest(url string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,
		// nolint:gosec
		nil
}

// PostRequest defines a wrapper around an HTTP POST request with a provided URL and data.
// An error is returned if the request or reading the body fails.
func PostRequest(url string, contentType string, data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// nolint:gosec
