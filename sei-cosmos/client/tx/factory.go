package tx

import (
	"github.com/spf13/pflag"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/keyring"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/tx/signing"
)

// Factory defines a client transaction factory that facilitates generating and
// signing an application-specific transaction.
type Factory struct {
	keybase            keyring.Keyring
	txConfig           client.TxConfig
	accountRetriever   client.AccountRetriever
	accountNumber      uint64
	sequence           uint64
	gas                uint64
	timeoutHeight      uint64
	gasAdjustment      float64
	chainID            string
	memo               string
	fees               sdk.Coins
	gasPrices          sdk.DecCoins
	signMode           signing.SignMode
	simulateAndExecute bool
}

// NewFactoryCLI creates a new Factory.
func NewFactoryCLI(clientCtx client.Context, flagSet *pflag.FlagSet) Factory {
	_ = "STUB: not implemented"
	return *new(Factory)
}

func (f Factory) AccountNumber() uint64    { _ = "STUB: not implemented"; return 0 }
func (f Factory) Sequence() uint64         { _ = "STUB: not implemented"; return 0 }
func (f Factory) Gas() uint64              { _ = "STUB: not implemented"; return 0 }
func (f Factory) GasAdjustment() float64   { _ = "STUB: not implemented"; return 0 }
func (f Factory) Keybase() keyring.Keyring { _ = "STUB: not implemented"; return *new(keyring.Keyring) }
func (f Factory) ChainID() string          { _ = "STUB: not implemented"; return "" }
func (f Factory) Memo() string             { _ = "STUB: not implemented"; return "" }
func (f Factory) Fees() sdk.Coins          { _ = "STUB: not implemented"; return *new(sdk.Coins) }
func (f Factory) GasPrices() sdk.DecCoins  { _ = "STUB: not implemented"; return *new(sdk.DecCoins) }
func (f Factory) AccountRetriever() client.AccountRetriever {
	_ = "STUB: not implemented"
	return *new(client.AccountRetriever)
}
func (f Factory) TimeoutHeight() uint64 { _ = "STUB: not implemented"; return 0 }

// SimulateAndExecute returns the option to simulate and then execute the transaction
// using the gas from the simulation results
func (f Factory) SimulateAndExecute() bool { _ = "STUB: not implemented"; return false }

// WithTxConfig returns a copy of the Factory with an updated TxConfig.
func (f Factory) WithTxConfig(g client.TxConfig) Factory {
	_ = "STUB: not implemented"
	return *

	// WithAccountRetriever returns a copy of the Factory with an updated AccountRetriever.
	new(Factory)
}

func (f Factory) WithAccountRetriever(ar client.AccountRetriever) Factory {
	_ = "STUB: not implemented"
	return *new(Factory)
}

// WithChainID returns a copy of the Factory with an updated chainID.
func (f Factory) WithChainID(chainID string) Factory {
	_ = "STUB: not implemented"
	return *new(Factory)
}

// WithGas returns a copy of the Factory with an updated gas value.
func (f Factory) WithGas(gas uint64) Factory {
	_ = "STUB: not implemented"
	return *

	// WithFees returns a copy of the Factory with an updated fee.
	new(Factory)
}

func (f Factory) WithFees(fees string) Factory { _ = "STUB: not implemented"; return *new(Factory) }

// WithGasPrices returns a copy of the Factory with updated gas prices.
func (f Factory) WithGasPrices(gasPrices string) Factory {
	_ = "STUB: not implemented"
	return *new(Factory)
}

// WithKeybase returns a copy of the Factory with updated Keybase.
func (f Factory) WithKeybase(keybase keyring.Keyring) Factory {
	_ = "STUB: not implemented"
	return *new(Factory)
}

// WithSequence returns a copy of the Factory with an updated sequence number.
func (f Factory) WithSequence(sequence uint64) Factory {
	_ = "STUB: not implemented"
	return *new(Factory)
}

// WithMemo returns a copy of the Factory with an updated memo.
func (f Factory) WithMemo(memo string) Factory {
	_ = "STUB: not implemented"
	return *

	// WithAccountNumber returns a copy of the Factory with an updated account number.
	new(Factory)
}

func (f Factory) WithAccountNumber(accnum uint64) Factory {
	_ = "STUB: not implemented"
	return *new(Factory)
}

// WithGasAdjustment returns a copy of the Factory with an updated gas adjustment.
func (f Factory) WithGasAdjustment(gasAdj float64) Factory {
	_ = "STUB: not implemented"
	return *new(Factory)
}

// WithSimulateAndExecute returns a copy of the Factory with an updated gas
// simulation value.
func (f Factory) WithSimulateAndExecute(sim bool) Factory {
	_ = "STUB: not implemented"
	return *new(Factory)
}

// SignMode returns the sign mode configured in the Factory
func (f Factory) SignMode() signing.SignMode {
	_ = "STUB: not implemented"

	// WithSignMode returns a copy of the Factory with an updated sign mode value.
	return *new(signing.SignMode)
}

func (f Factory) WithSignMode(mode signing.SignMode) Factory {
	_ = "STUB: not implemented"
	return *new(Factory)
}

// WithTimeoutHeight returns a copy of the Factory with an updated timeout height.
func (f Factory) WithTimeoutHeight(height uint64) Factory {
	_ = "STUB: not implemented"
	return *new(Factory)
}

// BuildUnsignedTx builds a transaction to be signed given a set of messages.
// Once created, the fee, memo, and messages are set.
func (f Factory) BuildUnsignedTx(msgs ...sdk.Msg) (client.TxBuilder, error) {
	_ = "STUB: not implemented"
	return *new(client.TxBuilder), nil
}

//nolint:gosec // bounds checked above

// Derive the fees based on the provided gas prices, where
// fee = ceil(gasPrice * gasLimit).

// PrintUnsignedTx will generate an unsigned transaction and print it to the writer
// specified by ctx.Output. If simulation was requested, the gas will be
// simulated and also printed to the same writer before the transaction is
// printed.
func (f Factory) PrintUnsignedTx(clientCtx client.Context, msgs ...sdk.Msg) error {
	_ = "STUB: not implemented"
	return nil
}

// BuildSimTx creates an unsigned tx with an empty single signature and returns
// the encoded transaction or an error if the unsigned transaction cannot be
// built.
func (f Factory) BuildSimTx(msgs ...sdk.Msg) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create an empty signature literal as the ante handler will populate with a
// sentinel pubkey.

// getSimPK gets the public key to use for building a simulation tx.
// Note, we should only check for keys in the keybase if we are in simulate and execute mode,
// e.g. when using --gas=auto.
// When using --dry-run, we are is simulation mode only and should not check the keybase.
// Ref: https://github.com/cosmos/cosmos-sdk/issues/11283
func (f Factory) getSimPK() (cryptotypes.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(cryptotypes.PubKey), nil
}

// use default public key type

// Use the first element from the list of keys in order to generate a valid
// pubkey that supports multiple algorithms.

// take the first record just for simulation purposes

// Prepare ensures the account defined by ctx.GetFromAddress() exists and
// if the account number and/or the account sequence number are zero (not set),
// they will be queried for and set on the provided Factory. A new Factory with
// the updated fields will be returned.
func (f Factory) Prepare(clientCtx client.Context) (Factory, error) {
	_ = "STUB: not implemented"
	return *new(Factory), nil
}
