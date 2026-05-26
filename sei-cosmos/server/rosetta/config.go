package rosetta

import (
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/spf13/pflag"

	crg "github.com/sei-protocol/sei-chain/sei-cosmos/server/rosetta/lib/server"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
)

// configuration defaults constants
const (
	// DefaultBlockchain defines the default blockchain identifier name
	DefaultBlockchain = "app"
	// DefaultAddr defines the default rosetta binding address
	DefaultAddr = ":8080"
	// DefaultRetries is the default number of retries
	DefaultRetries = 5
	// DefaultTendermintEndpoint is the default value for the tendermint endpoint
	DefaultTendermintEndpoint = "localhost:26657"
	// DefaultGRPCEndpoint is the default value for the gRPC endpoint
	DefaultGRPCEndpoint = "localhost:9090"
	// DefaultNetwork defines the default network name
	DefaultNetwork = "network"
	// DefaultOffline defines the default offline value
	DefaultOffline = false
)

// configuration flags
const (
	FlagBlockchain         = "blockchain"
	FlagNetwork            = "network"
	FlagTendermintEndpoint = "tendermint"
	FlagGRPCEndpoint       = "grpc"
	FlagAddr               = "addr"
	FlagRetries            = "retries"
	FlagOffline            = "offline"
)

// Config defines the configuration of the rosetta server
type Config struct {
	// Blockchain defines the blockchain name
	// defaults to DefaultBlockchain
	Blockchain string
	// Network defines the network name
	Network string
	// TendermintRPC defines the endpoint to connect to
	// tendermint RPC, specifying 'tcp://' before is not
	// required, usually it's at port 26657 of the
	TendermintRPC string
	// GRPCEndpoint defines the cosmos application gRPC endpoint
	// usually it is located at 9090 port
	GRPCEndpoint string
	// Addr defines the default address to bind the rosetta server to
	// defaults to DefaultAddr
	Addr string
	// Retries defines the maximum number of retries
	// rosetta will do before quitting
	Retries int
	// Offline defines if the server must be run in offline mode
	Offline bool
	// Codec overrides the default data and construction api client codecs
	Codec *codec.ProtoCodec
	// InterfaceRegistry overrides the default data and construction api interface registry
	InterfaceRegistry codectypes.InterfaceRegistry
}

// NetworkIdentifier returns the network identifier given the configuration
func (c *Config) NetworkIdentifier() *types.NetworkIdentifier {
	_ = "STUB: not implemented"
	return nil
}

// validate validates a configuration and sets
// its defaults in case they were not provided
func (c *Config) validate() error { _ = "STUB: not implemented"; return nil }

// these are must

// these are optional but it must be online

// WithCodec extends the configuration with a predefined Codec
func (c *Config) WithCodec(ir codectypes.InterfaceRegistry, cdc *codec.ProtoCodec) {
	_ = "STUB: not implemented"
	return
}

// FromFlags gets the configuration from flags
func FromFlags(flags *pflag.FlagSet) (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

func ServerFromConfig(conf *Config) (crg.Server, error) {
	_ = "STUB: not implemented"
	return *new(crg.Server), nil
}

// SetFlags sets the configuration flags to the given flagset
func SetFlags(flags *pflag.FlagSet) { _ = "STUB: not implemented"; return }
