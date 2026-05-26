package tx

import (
	"context"
	"net/http"

	gogogrpc "github.com/gogo/protobuf/grpc"
	"github.com/spf13/pflag"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/rest"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/tx"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/tx/signing"
	authsigning "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/signing"
)

// GenerateOrBroadcastTxCLI will either generate and print and unsigned transaction
// or sign it and broadcast it returning an error upon failure.
func GenerateOrBroadcastTxCLI(ctx context.Context, clientCtx client.Context, flagSet *pflag.FlagSet, msgs ...sdk.Msg) error {
	_ = "STUB: not implemented"
	return nil
}

// GenerateOrBroadcastTxWithFactory will either generate and print and unsigned transaction
// or sign it and broadcast it returning an error upon failure.
func GenerateOrBroadcastTxWithFactory(ctx context.Context, clientCtx client.Context, txf Factory, msgs ...sdk.Msg) error {
	_ = "STUB: not implemented"
	// Validate all msgs before generating or broadcasting the tx.
	// We were calling ValidateBasic separately in each CLI handler before.
	// Right now, we're factorizing that call inside this function.
	// ref: https://github.com/cosmos/cosmos-sdk/pull/9236#discussion_r623803504
	return nil
}

// GenerateTx will generate an unsigned transaction and print it to the writer
// specified by ctx.Output. If simulation was requested, the gas will be
// simulated and also printed to the same writer before the transaction is
// printed.
func GenerateTx(clientCtx client.Context, txf Factory, msgs ...sdk.Msg) error {
	_ = "STUB: not implemented"
	return nil
}

// BroadcastTx attempts to generate, sign and broadcast a transaction with the
// given set of messages. It will also simulate gas requirements if necessary.
// It will return an error upon failure.
func broadcastTx(ctx context.Context, clientCtx client.Context, txf Factory, msgs ...sdk.Msg) error {
	_ = "STUB: not implemented"
	return nil
}

// broadcast to a Tendermint node

// WriteGeneratedTxResponse writes a generated unsigned transaction to the
// provided http.ResponseWriter. It will simulate gas costs if requested by the
// BaseReq. Upon any error, the error will be written to the http.ResponseWriter.
// Note that this function returns the legacy StdTx Amino JSON format for compatibility
// with legacy clients.
// Deprecated: We are removing Amino soon.
func WriteGeneratedTxResponse(
	clientCtx client.Context, w http.ResponseWriter, br rest.BaseReq, msgs ...sdk.Msg,
) {
	_ = "STUB: not implemented"
	return
}

// BuildUnsignedTx builds a transaction to be signed given a set of messages. The
// transaction is initially created via the provided factory's generator. Once
// created, the fee, memo, and messages are set.
func BuildUnsignedTx(txf Factory, msgs ...sdk.Msg) (client.TxBuilder, error) {
	_ = "STUB: not implemented"
	return *new(client.TxBuilder), nil
}

// BuildSimTx creates an unsigned tx with an empty single signature and returns
// the encoded transaction or an error if the unsigned transaction cannot be
// built.
func BuildSimTx(txf Factory, msgs ...sdk.Msg) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// CalculateGas simulates the execution of a transaction and returns the
	// simulation response obtained by the query and the adjusted gas amount.
}

func CalculateGas(
	clientCtx gogogrpc.ClientConn, txf Factory, msgs ...sdk.Msg,
) (*tx.SimulateResponse, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// prepareFactory ensures the account defined by ctx.GetFromAddress() exists and
// if the account number and/or the account sequence number are zero (not set),
// they will be queried for and set on the provided Factory. A new Factory with
// the updated fields will be returned.
func prepareFactory(clientCtx client.Context, txf Factory) (Factory, error) {
	_ = "STUB: not implemented"
	return *new(Factory), nil
}

// SignWithPrivKey signs a given tx with the given private key, and returns the
// corresponding SignatureV2 if the signing is successful.
func SignWithPrivKey(
	signMode signing.SignMode, signerData authsigning.SignerData,
	txBuilder client.TxBuilder, priv cryptotypes.PrivKey, txConfig client.TxConfig,
	accSeq uint64,
) (signing.SignatureV2, error) {
	_ = "STUB: not implemented"
	return *new(signing.SignatureV2), nil
}

// Generate the bytes to be signed.

// Sign those bytes

// Construct the SignatureV2 struct

func checkMultipleSigners(mode signing.SignMode, tx authsigning.Tx) error {
	_ = "STUB: not implemented"
	return nil
}

// Sign signs a given tx with a named key. The bytes signed over are canconical.
// The resulting signature will be added to the transaction builder overwriting the previous
// ones if overwrite=true (otherwise, the signature will be appended).
// Signing a transaction with mutltiple signers in the DIRECT mode is not supprted and will
// return an error.
// An error is returned upon failure.
func Sign(txf Factory, name string, txBuilder client.TxBuilder, overwriteSig bool) error {
	_ = "STUB: not implemented"
	return nil
}

// use the SignModeHandler's default mode if unspecified

// For SIGN_MODE_DIRECT, calling SetSignatures calls setSignerInfos on
// TxBuilder under the hood, and SignerInfos is needed to generated the
// sign bytes. This is the reason for setting SetSignatures here, with a
// nil signature.
//
// Note: this line is not needed for SIGN_MODE_LEGACY_AMINO, but putting it
// also doesn't affect its generated sign bytes, so for code's simplicity
// sake, we put it here.

//nolint:prealloc // It's behind a flag.

// Generate the bytes to be signed.

// Sign those bytes

// Construct the SignatureV2 struct

// GasEstimateResponse defines a response definition for tx gas estimation.
type GasEstimateResponse struct {
	GasEstimate uint64 `json:"gas_estimate" yaml:"gas_estimate"`
}

func (gr GasEstimateResponse) String() string { _ = "STUB: not implemented"; return "" }
