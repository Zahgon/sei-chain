package errors

import (
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
)

const (
	// SuccessABCICode declares an ABCI response use 0 to signal that the
	// processing was successful and no error is returned.
	SuccessABCICode = abci.CodeTypeOK

	// All unclassified errors that do not provide an ABCI code are clubbed
	// under an internal error code and a generic message instead of
	// detailed error string.
	internalABCICodespace        = UndefinedCodespace
	internalABCICode      uint32 = 1
)

// safeIntFromUint64 converts uint64 to int64, capping at math.MaxInt64 to prevent overflow.
func safeIntFromUint64(v uint64) int64 { _ = "STUB: not implemented"; return 0 }

// ABCIInfo returns the ABCI error information as consumed by the tendermint
// client. Returned codespace, code, and log message should be used as a ABCI response.
// Any error that does not provide ABCICode information is categorized as error
// with code 1, codespace UndefinedCodespace
// When not running in a debug mode all messages of errors that do not provide
// ABCICode information are replaced with generic "internal error". Errors
// without an ABCICode information as considered internal.
func ABCIInfo(err error, debug bool) (codespace string, code uint32, log string) {
	_ = "STUB: not implemented"
	return "", 0, ""
}

// ResponseDeliverTx returns an ABCI ResponseDeliverTx object with fields filled in
// from the given error and gas values.
func ResponseDeliverTx(err error, gw, gu uint64, debug bool) abci.ResponseDeliverTx {
	_ = "STUB: not implemented"
	return *new(abci.ResponseDeliverTx)
}

// ResponseDeliverTxWithEvents returns an ABCI ResponseDeliverTx object with fields filled in
// from the given error, gas values and events.
func ResponseDeliverTxWithEvents(err error, gw, gu uint64, events []abci.Event, debug bool) abci.ResponseDeliverTx {
	_ = "STUB: not implemented"
	return *new(abci.ResponseDeliverTx)
}

// QueryResult returns a ResponseQuery from an error. It will try to parse ABCI
// info from the error.
func QueryResult(err error) abci.ResponseQuery {
	_ = "STUB: not implemented"
	return *new(abci.ResponseQuery)
}

// QueryResultWithDebug returns a ResponseQuery from an error. It will try to parse ABCI
// info from the error. It will use debugErrEncoder if debug parameter is true.
// Starting from v0.46, this function will be removed, and be replaced by `QueryResult`.
func QueryResultWithDebug(err error, debug bool) abci.ResponseQuery {
	_ = "STUB: not implemented"
	return *new(abci.ResponseQuery)
}

// The debugErrEncoder encodes the error with a stacktrace.
func debugErrEncoder(err error) string { _ = "STUB: not implemented"; return "" }

func defaultErrEncoder(err error) string { _ = "STUB: not implemented"; return "" }

type coder interface {
	ABCICode() uint32
}

// abciCode tests if given error contains an ABCI code and returns the value of
// it if available. This function is testing for the causer interface as well
// and unwraps the error.
func abciCode(err error) uint32 { _ = "STUB: not implemented"; return 0 }

type codespacer interface {
	Codespace() string
}

// abciCodespace tests if given error contains a codespace and returns the value of
// it if available. This function is testing for the causer interface as well
// and unwraps the error.
func abciCodespace(err error) string { _ = "STUB: not implemented"; return "" }

// errIsNil returns true if value represented by the given error is nil.
//
// Most of the time a simple == check is enough. There is a very narrowed
// spectrum of cases (mostly in tests) where a more sophisticated check is
// required.
func errIsNil(err error) bool { _ = "STUB: not implemented"; return false }
