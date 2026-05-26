package server

import (
	"net/http"
)

// uriReqID is a placeholder ID used for GET requests, which do not receive a
// JSON-RPC request ID from the caller.
const uriReqID = -1

// convert from a function name to the http handler
func makeHTTPHandler(rpcFunc *RPCFunc) func(http.ResponseWriter, *http.Request) {
	_ = "STUB: not implemented"
	return nil
}

func parseURLParams(args []argInfo, req *http.Request) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// isQuotedString reports whether s is enclosed in double quotes.
func isQuotedString(s string) bool { _ = "STUB: not implemented"; return false }

// decodeInteger decodes s into an int64. If s is "double quoted" the quotes
// are removed; otherwise s must be a base-10 digit string.
func decodeInteger(s string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }
