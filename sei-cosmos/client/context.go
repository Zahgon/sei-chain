package client

import (
	"io"
	"net/url"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"

	"github.com/gogo/protobuf/proto"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
	rpcclient "github.com/sei-protocol/sei-chain/sei-tendermint/rpc/client"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/keyring"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

type Client = rpcclient.Client
type LocalClient interface {
	Client
	EvmNextPendingNonce(addr common.Address) uint64
	EvmProxy(sender common.Address) (*url.URL, bool)
}

type Context struct {
	FromAddress sdk.AccAddress
	Client      utils.Option[rpcclient.Client]
	ChainID     string
	// Deprecated: Codec codec will be changed to Codec: codec.Codec
	JSONCodec         codec.JSONCodec
	Codec             codec.Codec
	InterfaceRegistry codectypes.InterfaceRegistry
	Input             io.Reader
	Keyring           keyring.Keyring
	KeyringOptions    []keyring.Option
	Output            io.Writer
	OutputFormat      string
	Height            int64
	HomeDir           string
	KeyringDir        string
	From              string
	BroadcastMode     string
	FromName          string
	SignModeStr       string
	UseLedger         bool
	Simulate          bool
	GenerateOnly      bool
	PrintSignedOnly   bool
	Offline           bool
	SkipConfirm       bool
	TxConfig          TxConfig
	AccountRetriever  AccountRetriever
	NodeURI           string
	FeeGranter        sdk.AccAddress
	Viper             *viper.Viper

	// TODO: Deprecated (remove).
	LegacyAmino *codec.LegacyAmino
}

// WithKeyring returns a copy of the context with an updated keyring.
func (ctx Context) WithKeyring(k keyring.Keyring) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithKeyringOptions returns a copy of the context with an updated keyring.
func (ctx Context) WithKeyringOptions(opts ...keyring.Option) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithInput returns a copy of the context with an updated input.
func (ctx Context) WithInput(r io.Reader) Context {
	_ = "STUB: not implemented"
	// convert to a bufio.Reader to have a shared buffer between the keyring and the
	// the Commands, ensuring a read from one advance the read pointer for the other.
	// see https://github.com/cosmos/cosmos-sdk/issues/9566.
	return *new(Context)
}

// Deprecated: WithJSONCodec returns a copy of the Context with an updated JSONCodec.
func (ctx Context) WithJSONCodec(m codec.JSONCodec) Context {
	_ = "STUB: not implemented"

	// since we are using ctx.Codec everywhere in the SDK, for backward compatibility
	// we need to try to set it here as well.
	return *new(Context)
}

// WithCodec returns a copy of the Context with an updated Codec.
func (ctx Context) WithCodec(m codec.Codec) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithLegacyAmino returns a copy of the context with an updated LegacyAmino codec.
// TODO: Deprecated (remove).
func (ctx Context) WithLegacyAmino(cdc *codec.LegacyAmino) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithOutput returns a copy of the context with an updated output writer (e.g. stdout).
func (ctx Context) WithOutput(w io.Writer) Context { _ = "STUB: not implemented"; return *new(Context) }

// WithFrom returns a copy of the context with an updated from address or name.
func (ctx Context) WithFrom(from string) Context { _ = "STUB: not implemented"; return *new(Context) }

// WithOutputFormat returns a copy of the context with an updated OutputFormat field.
func (ctx Context) WithOutputFormat(format string) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithNodeURI returns a copy of the context with an updated node URI.
func (ctx Context) WithNodeURI(nodeURI string) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithHeight returns a copy of the context with an updated height.
func (ctx Context) WithHeight(height int64) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithClient returns a copy of the context with an updated RPC client
// instance.
func (ctx Context) WithClient(client Client) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithUseLedger returns a copy of the context with an updated UseLedger flag.
func (ctx Context) WithUseLedger(useLedger bool) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithChainID returns a copy of the context with an updated chain ID.
func (ctx Context) WithChainID(chainID string) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithHomeDir returns a copy of the Context with HomeDir set.
func (ctx Context) WithHomeDir(dir string) Context { _ = "STUB: not implemented"; return *new(Context) }

// WithKeyringDir returns a copy of the Context with KeyringDir set.
func (ctx Context) WithKeyringDir(dir string) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithGenerateOnly returns a copy of the context with updated GenerateOnly value
func (ctx Context) WithGenerateOnly(generateOnly bool) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithSimulation returns a copy of the context with updated Simulate value
func (ctx Context) WithSimulation(simulate bool) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithOffline returns a copy of the context with updated Offline value.
func (ctx Context) WithOffline(offline bool) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithFromName returns a copy of the context with an updated from account name.
func (ctx Context) WithFromName(name string) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithFromAddress returns a copy of the context with an updated from account
// address.
func (ctx Context) WithFromAddress(addr sdk.AccAddress) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithFeeGranterAddress returns a copy of the context with an updated fee granter account
// address.
func (ctx Context) WithFeeGranterAddress(addr sdk.AccAddress) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithBroadcastMode returns a copy of the context with an updated broadcast
// mode.
func (ctx Context) WithBroadcastMode(mode string) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithSignModeStr returns a copy of the context with an updated SignMode
// value.
func (ctx Context) WithSignModeStr(signModeStr string) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithSkipConfirmation returns a copy of the context with an updated SkipConfirm
// value.
func (ctx Context) WithSkipConfirmation(skip bool) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithTxConfig returns the context with an updated TxConfig
func (ctx Context) WithTxConfig(generator TxConfig) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithAccountRetriever returns the context with an updated AccountRetriever
func (ctx Context) WithAccountRetriever(retriever AccountRetriever) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithInterfaceRegistry returns the context with an updated InterfaceRegistry
func (ctx Context) WithInterfaceRegistry(interfaceRegistry codectypes.InterfaceRegistry) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithViper returns the context with Viper field. This Viper instance is used to read
// client-side config from the config file.
func (ctx Context) WithViper(prefix string) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// PrintString prints the raw string to ctx.Output if it's defined, otherwise to os.Stdout
func (ctx Context) PrintString(str string) error { _ = "STUB: not implemented"; return nil }

// PrintBytes prints the raw bytes to ctx.Output if it's defined, otherwise to os.Stdout.
// NOTE: for printing a complex state object, you should use ctx.PrintOutput
func (ctx Context) PrintBytes(o []byte) error { _ = "STUB: not implemented"; return nil }

// PrintProto outputs toPrint to the ctx.Output based on ctx.OutputFormat which is
// either text or json. If text, toPrint will be YAML encoded. Otherwise, toPrint
// will be JSON encoded using ctx.Codec. An error is returned upon failure.
func (ctx Context) PrintProto(toPrint proto.Message) error {
	_ = "STUB: not implemented"
	// always serialize JSON initially because proto json can't be directly YAML encoded
	return nil
}

// PrintObjectLegacy is a variant of PrintProto that doesn't require a proto.Message type
// and uses amino JSON encoding.
// Deprecated: It will be removed in the near future!
func (ctx Context) PrintObjectLegacy(toPrint any) error { _ = "STUB: not implemented"; return nil }

func (ctx Context) printOutput(out []byte) error { _ = "STUB: not implemented"; return nil }

// handle text format by decoding and re-encoding JSON as YAML

// append new-line for formats besides YAML

// GetFromFields returns a from account address, account name and keyring type, given either an address or key name.
// If clientCtx.Simulate is true the keystore is not accessed and a valid address must be provided
// If clientCtx.GenerateOnly is true the keystore is only accessed if a key name is provided
func GetFromFields(clientCtx Context, kr keyring.Keyring, from string) (sdk.AccAddress, string, keyring.KeyType, error) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), "", *new(keyring.KeyType), nil
}

// NewKeyringFromBackend gets a Keyring object from a backend
func NewKeyringFromBackend(ctx Context, backend string) (keyring.Keyring, error) {
	_ = "STUB: not implemented"
	return *new(keyring.Keyring), nil
}
