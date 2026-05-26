package network

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/gofrs/flock"
	"github.com/sei-protocol/sei-chain/app/params"
	tmservice "github.com/sei-protocol/sei-chain/sei-tendermint/libs/service"
	tmclient "github.com/sei-protocol/sei-chain/sei-tendermint/rpc/client"
	"google.golang.org/grpc"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/keyring"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/server"
	"github.com/sei-protocol/sei-chain/sei-cosmos/server/api"
	srvconfig "github.com/sei-protocol/sei-chain/sei-cosmos/server/config"
	servertypes "github.com/sei-protocol/sei-chain/sei-cosmos/server/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// package-wide lock (in-process) + file lock (cross-process) so that
// multiple go test binaries don't race on port preallocation.
var (
	lock     = new(sync.Mutex)
	lockPath = filepath.Join(os.TempDir(), "sei-cosmos-test-network.lock")
)

// AppConstructor defines a function which accepts a network configuration and
// creates an ABCI Application to provide to Tendermint.
type AppConstructor = func(val Validator) servertypes.Application

// NewAppConstructor returns a new simapp AppConstructor
func NewAppConstructor(t *testing.T, encodingCfg params.EncodingConfig) AppConstructor {
	_ = "STUB: not implemented"
	return *new(AppConstructor)
}

// Config defines the necessary configuration used to bootstrap and start an
// in-process local testing network.
type Config struct {
	Codec             codec.Codec
	LegacyAmino       *codec.LegacyAmino // TODO: Remove!
	InterfaceRegistry codectypes.InterfaceRegistry

	TxConfig         client.TxConfig
	AccountRetriever client.AccountRetriever
	AppConstructor   AppConstructor             // the ABCI application constructor
	GenesisState     map[string]json.RawMessage // custom genesis state to provide
	TimeoutCommit    time.Duration              // the consensus commitment timeout
	ChainID          string                     // the network chain-id
	NumValidators    int                        // the total number of validators to create and bond
	Mnemonics        []string                   // custom user-provided validator operator mnemonics
	BondDenom        string                     // the staking bond denomination
	MinGasPrices     string                     // the minimum gas prices each validator will accept
	AccountTokens    sdk.Int                    // the amount of unique validator tokens (e.g. 1000node0)
	StakingTokens    sdk.Int                    // the amount of tokens each validator has available to stake
	BondedTokens     sdk.Int                    // the amount of tokens each validator stakes
	PruningStrategy  string                     // the pruning strategy each validator will have
	EnableLogging    bool                       // enable Tendermint logging to STDOUT
	CleanupDir       bool                       // remove base temporary directory during cleanup
	SigningAlgo      string                     // signing algorithm for keys
	EnableGRPCWeb    bool                       // enable gRPC-Web server (off by default)

	KeyringOptions []keyring.Option
}

// DefaultConfig returns a sane default configuration suitable for nearly all
// testing requirements.
func DefaultConfig(t *testing.T) Config { _ = "STUB: not implemented"; return *new(Config) }

type (
	// Network defines a local in-process testing network using SimApp. It can be
	// configured to start any number of validators, each with its own RPC and API
	// clients. Typically, this test network would be used in client and integration
	// testing where user input is expected.
	//
	// Note, due to Tendermint constraints in regards to RPC functionality, there
	// may only be one test network running at a time. Thus, any caller must be
	// sure to Cleanup after testing is finished in order to allow other tests
	// to create networks. In addition, only the first validator will have a valid
	// RPC and API server/client.
	Network struct {
		T          *testing.T
		BaseDir    string
		Validators []*Validator

		Config   Config
		fileLock *flock.Flock
	}

	// Validator defines an in-process Tendermint validator node. Through this object,
	// a client can make RPC and API calls and interact with any client command
	// or handler.
	Validator struct {
		AppConfig  *srvconfig.Config
		ClientCtx  client.Context
		Ctx        *server.Context
		Dir        string
		NodeID     string
		PubKey     cryptotypes.PubKey
		Moniker    string
		APIAddress string
		RPCAddress string
		P2PAddress string
		Address    sdk.AccAddress
		ValAddress sdk.ValAddress
		RPCClient  tmclient.Client
		GoCtx      context.Context

		tmNode   tmservice.Service
		api      *api.Server
		grpc     *grpc.Server
		grpcWeb  *http.Server
		cancelFn context.CancelFunc
	}
)

// New creates a new Network for integration tests.
func New(t *testing.T, cfg Config) *Network {
	_ = "STUB: not implemented"
	// only one caller/test can create and use a network at a time
	return nil
}

// generate private keys, node IDs, and initial transactions

// Only allow the first validator to expose an RPC, API and gRPC
// server/client due to Tendermint in-process constraints.

// save private key seed words

// Arbitrary fee
// Need at least 100386

// Ensure we cleanup incase any test was abruptly halted (e.g. SIGINT) as any
// defer in a test would not be called.

// LatestHeight returns the latest height of the network or an error if the
// query fails or no validators exist.
func (n *Network) LatestHeight() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// WaitForHeight performs a blocking check where it waits for a block to be
// committed after a given block. If that height is not reached within a timeout,
// an error is returned. Regardless, the latest height queried is returned.
func (n *Network) WaitForHeight(h int64) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// WaitForHeightWithTimeout is the same as WaitForHeight except the caller can
// provide a custom timeout.
func (n *Network) WaitForHeightWithTimeout(h int64, t time.Duration) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// WaitForNextBlock waits for the next block to be committed, returning an error
// upon failure.
func (n *Network) WaitForNextBlock() error { _ = "STUB: not implemented"; return nil }

// Cleanup removes the root testing (temporary) directory and stops both the
// Tendermint and API services. It allows other callers to create and start
// test networks. This method must be called when a test is finished, typically
// in a defer.
func (n *Network) Cleanup() { _ = "STUB: not implemented"; return }

// Always cancel the context to avoid leaks, regardless of node state.
