package legacytx

import (
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/tx/signing"
)

// Interface implementation checks
var (
	_ sdk.Tx                             = (*StdTx)(nil)
	_ sdk.TxWithMemo                     = (*StdTx)(nil)
	_ sdk.FeeTx                          = (*StdTx)(nil)
	_ codectypes.UnpackInterfacesMessage = (*StdTx)(nil)

	_ codectypes.UnpackInterfacesMessage = (*StdSignature)(nil)
)

// StdFee includes the amount of coins paid in fees and the maximum
// gas to be used by the transaction. The ratio yields an effective "gasprice",
// which must be above some miminum to be accepted into the mempool.
// [Deprecated]
type StdFee struct {
	Amount sdk.Coins `json:"amount" yaml:"amount"`
	Gas    uint64    `json:"gas" yaml:"gas"`
}

// Deprecated: NewStdFee returns a new instance of StdFee
func NewStdFee(gas uint64, amount sdk.Coins) StdFee { _ = "STUB: not implemented"; return *new(StdFee) }

// GetGas returns the fee's (wanted) gas.
func (fee StdFee) GetGas() uint64 {
	_ = "STUB: not implemented"

	// GetAmount returns the fee's amount.
	return 0
}

func (fee StdFee) GetAmount() sdk.Coins {
	_ = "STUB: not implemented"

	// Bytes returns the encoded bytes of a StdFee.
	return *new(sdk.Coins)
}

func (fee StdFee) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// GasPrices returns the gas prices for a StdFee.
//
// NOTE: The gas prices returned are not the true gas prices that were
// originally part of the submitted transaction because the fee is computed
// as fee = ceil(gasWanted * gasPrices).
func (fee StdFee) GasPrices() sdk.DecCoins { _ = "STUB: not implemented"; return *new(sdk.DecCoins) }

//nolint:gosec G115 -- bounds checked above

// StdTx is the legacy transaction format for wrapping a Msg with Fee and Signatures.
// It only works with Amino, please prefer the new protobuf Tx in types/tx.
// NOTE: the first signature is the fee payer (Signatures must not be nil).
// Deprecated
type StdTx struct {
	Msgs          []sdk.Msg      `json:"msg" yaml:"msg"`
	Fee           StdFee         `json:"fee" yaml:"fee"`
	Signatures    []StdSignature `json:"signatures" yaml:"signatures"`
	Memo          string         `json:"memo" yaml:"memo"`
	TimeoutHeight uint64         `json:"timeout_height" yaml:"timeout_height"`
	GasEstimate   uint64         `json:"gas_estimate" yaml:"gas_estimate"`
}

// Deprecated
func NewStdTx(msgs []sdk.Msg, fee StdFee, sigs []StdSignature, memo string) StdTx {
	_ = "STUB: not implemented"
	return *new(StdTx)
}

// GetMsgs returns the all the transaction's messages.
func (tx StdTx) GetMsgs() []sdk.Msg {
	_ = "STUB: not implemented"

	// ValidateBasic does a simple and lightweight validation check that doesn't
	// require access to any other information.
	return nil
}

func (tx StdTx) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (tx StdTx) GetGasEstimate() uint64 { _ = "STUB: not implemented"; return 0 }

// Deprecated: AsAny implements intoAny. It doesn't work for protobuf serialization,
// so it can't be saved into protobuf configured storage. We are using it only for API
// compatibility.
func (tx *StdTx) AsAny() *codectypes.Any { _ = "STUB: not implemented"; return nil }

// GetSigners returns the addresses that must sign the transaction.
// Addresses are returned in a deterministic order.
// They are accumulated from the GetSigners method for each Msg
// in the order they appear in tx.GetMsgs().
// Duplicate addresses will be omitted.
func (tx StdTx) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// GetMemo returns the memo
func (tx StdTx) GetMemo() string {
	_ = "STUB: not implemented"

	// GetTimeoutHeight returns the transaction's timeout height (if set).
	return ""
}

func (tx StdTx) GetTimeoutHeight() uint64 { _ = "STUB: not implemented"; return 0 }

// GetSignatures returns the signature of signers who signed the Msg.
// CONTRACT: Length returned is same as length of
// pubkeys returned from MsgKeySigners, and the order
// matches.
// CONTRACT: If the signature is missing (ie the Msg is
// invalid), then the corresponding signature is
// .Empty().
func (tx StdTx) GetSignatures() [][]byte { _ = "STUB: not implemented"; return nil }

// GetSignaturesV2 implements SigVerifiableTx.GetSignaturesV2
func (tx StdTx) GetSignaturesV2() ([]signing.SignatureV2, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPubkeys returns the pubkeys of signers if the pubkey is included in the signature
// If pubkey is not included in the signature, then nil is in the slice instead
func (tx StdTx) GetPubKeys() ([]cryptotypes.PubKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetGas returns the Gas in StdFee
func (tx StdTx) GetGas() uint64 {
	_ = "STUB: not implemented"

	// GetFee returns the FeeAmount in StdFee
	return 0
}

func (tx StdTx) GetFee() sdk.Coins {
	_ = "STUB: not implemented"
	return *

	// FeePayer returns the address that is responsible for paying fee
	// StdTx returns the first signer as the fee payer
	// If no signers for tx, return empty address
	new(sdk.Coins)
}

func (tx StdTx) FeePayer() sdk.AccAddress { _ = "STUB: not implemented"; return *new(sdk.AccAddress) }

// FeeGranter always returns nil for StdTx
func (tx StdTx) FeeGranter() sdk.AccAddress { _ = "STUB: not implemented"; return *new(sdk.AccAddress) }

func (tx StdTx) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// Signatures contain PubKeys, which need to be unpacked.
