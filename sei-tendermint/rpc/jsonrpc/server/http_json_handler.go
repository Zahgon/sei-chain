package server

import (
	"html/template"
	"net/http"

	rpctypes "github.com/sei-protocol/sei-chain/sei-tendermint/rpc/jsonrpc/types"
)

// HTTP + JSON handler

const REQUEST_BATCH_SIZE_LIMIT = 10

// jsonrpc calls grab the given method's function info and runs reflect.Call
func makeJSONRPCHandler(funcMap map[string]*RPCFunc) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// For POST requests, reject a non-root URL path. This should not happen
// in the standard configuration, since the wrapper checks the path.

// if its an empty request (like from a browser), just display a list of
// functions

// Ignore notifications, which this service does not support.

func ensureBodyClose(next http.HandlerFunc) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func handleInvalidJSONRPCPaths(next http.HandlerFunc) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

//  we check whether the path is indeed "/", otherwise return a 404 error

// parseRequests parses a JSON-RPC request or request batch from data.
func parseRequests(data []byte) ([]rpctypes.RPCRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// writes a list of available rpc endpoints as an html page
func writeListOfEndpoints(w http.ResponseWriter, funcMap map[string]*RPCFunc) {
	_ = "STUB: not implemented"
	return
}

var listOfEndpoints = template.Must(template.New("list").Parse(`<html>
<head><title>List of RPC Endpoints</title></head>
<body>

<h1>Available RPC endpoints:</h1>

{{if .NoArgs}}
<hr />
<h2>Endpoints with no arguments:</h2>

<ul>
{{range $link := .NoArgs}}  <li><a href="{{$link}}">{{$link}}</a></li>
{{end -}}
</ul>{{end}}

{{if .HasArgs}}
<hr />
<h2>Endpoints that require arguments:</h2>

<ul>
{{range $link := .HasArgs}}  <li><a href="{{$link}}">{{$link}}</a></li>
{{end -}}
</ul>{{end}}

</body></html>`))
