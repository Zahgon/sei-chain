package server

import (
	"net/http"
	"time"

	"github.com/coinbase/rosetta-sdk-go/types"

	crgtypes "github.com/sei-protocol/sei-chain/sei-cosmos/server/rosetta/lib/types"
)

const DefaultRetries = 5
const DefaultRetryWait = 5 * time.Second

// Settings define the rosetta server settings
type Settings struct {
	// Network contains the information regarding the network
	Network *types.NetworkIdentifier
	// Client is the online API handler
	Client crgtypes.Client
	// Listen is the address the handler will listen at
	Listen string
	// Offline defines if the rosetta service should be exposed in offline mode
	Offline bool
	// Retries is the number of readiness checks that will be attempted when instantiating the handler
	// valid only for online API
	Retries int
	// RetryWait is the time that will be waited between retries
	RetryWait time.Duration
}

type Server struct {
	h    http.Handler
	addr string
}

func (h Server) Start() error { _ = "STUB: not implemented"; return nil }

func NewServer(settings Settings) (Server, error) {
	_ = "STUB: not implemented"
	return *new(Server), nil
}

func newOfflineAdapter(settings Settings) (crgtypes.API, error) {
	_ = "STUB: not implemented"
	return *new(crgtypes.API), nil
}

func newOnlineAdapter(settings Settings) (crgtypes.API, error) {
	_ = "STUB: not implemented"
	return *new(crgtypes.API), nil
}
