package client

import (
	"bufio"
	"io"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client/tx"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// GasEstimateResponse defines a response definition for tx gas estimation.
type GasEstimateResponse struct {
	GasEstimate uint64 `json:"gas_estimate" yaml:"gas_estimate"`
}

func (gr GasEstimateResponse) String() string { _ = "STUB: not implemented"; return "" }

// PrintUnsignedStdTx builds an unsigned StdTx and prints it to os.Stdout.
func PrintUnsignedStdTx(txBldr tx.Factory, clientCtx client.Context, msgs []sdk.Msg) error {
	_ = "STUB: not implemented"
	return nil
}

// SignTx signs a transaction managed by the TxBuilder using a `name` key stored in Keybase.
// The new signature is appended to the TxBuilder when overwrite=false or overwritten otherwise.
// Don't perform online validation or lookups if offline is true.
func SignTx(txFactory tx.Factory, clientCtx client.Context, name string, txBuilder client.TxBuilder, offline, overwriteSig bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Ledger and Multisigs only support LEGACY_AMINO_JSON signing.

// SignTxWithSignerAddress attaches a signature to a transaction.
// Don't perform online validation or lookups if offline is true, else
// populate account and sequence numbers from a foreign account.
// This function should only be used when signing with a multisig. For
// normal keys, please use SignTx directly.
func SignTxWithSignerAddress(txFactory tx.Factory, clientCtx client.Context, addr sdk.AccAddress,
	name string, txBuilder client.TxBuilder, offline, overwrite bool) (err error) {
	_ = "STUB: not implemented"
	// Multisigs only support LEGACY_AMINO_JSON signing.
	return nil
}

// check whether the address is a signer

// Read and decode a StdTx from the given filename.  Can pass "-" to read from stdin.
func ReadTxFromFile(ctx client.Context, filename string) (tx sdk.Tx, err error) {
	_ = "STUB: not implemented"
	return *new(sdk.Tx), nil
}

// NewBatchScanner returns a new BatchScanner to read newline-delimited StdTx transactions from r.
func NewBatchScanner(cfg client.TxConfig, r io.Reader) *BatchScanner {
	_ = "STUB: not implemented"
	return nil
}

// BatchScanner provides a convenient interface for reading batch data such as a file
// of newline-delimited JSON encoded StdTx.
type BatchScanner struct {
	*bufio.Scanner
	theTx        sdk.Tx
	cfg          client.TxConfig
	unmarshalErr error
}

// Tx returns the most recent Tx unmarshalled by a call to Scan.
func (bs BatchScanner) Tx() sdk.Tx {
	_ = "STUB: not implemented"

	// UnmarshalErr returns the first unmarshalling error that was encountered by the scanner.
	return *new(sdk.Tx)
}

func (bs BatchScanner) UnmarshalErr() error { _ = "STUB: not implemented"; return nil }

// Scan advances the Scanner to the next line.
func (bs *BatchScanner) Scan() bool { _ = "STUB: not implemented"; return false }

func populateAccountFromState(
	txBldr tx.Factory, clientCtx client.Context, addr sdk.AccAddress,
) (tx.Factory, error) {
	_ = "STUB: not implemented"
	return *new(tx.Factory), nil
}

// GetTxEncoder return tx encoder from global sdk configuration if ones is defined.
// Otherwise returns encoder with default logic.
func GetTxEncoder(cdc *codec.LegacyAmino) (encoder sdk.TxEncoder) {
	_ = "STUB: not implemented"
	return *new(sdk.TxEncoder)
}

func ParseQueryResponse(bz []byte) (sdk.SimulationResponse, error) {
	_ = "STUB: not implemented"
	return *new(sdk.SimulationResponse), nil
}

func isTxSigner(user sdk.AccAddress, signers []sdk.AccAddress) bool {
	_ = "STUB: not implemented"
	return false
}
