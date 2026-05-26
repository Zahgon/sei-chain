package ante

import (
	"encoding/hex"

	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/keys/secp256k1"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types/multisig"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/tx/signing"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/legacy/legacytx"
	authsigning "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/signing"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/types"
)

var (
	// simulation signature values used to estimate gas consumption
	key                = make([]byte, secp256k1.PubKeySize)
	simSecp256k1Pubkey = &secp256k1.PubKey{Key: key}
	simSecp256k1Sig    [64]byte

	_ authsigning.SigVerifiableTx = (*legacytx.StdTx)(nil) // assert StdTx implements SigVerifiableTx
)

func init() {
	// This decodes a valid hex string into a sepc256k1Pubkey for use in transaction simulation
	bz, _ := hex.DecodeString("035AD6810A47F073553FF30D2FCC7E0D3B1C0B74B61A1AAA2582344037151E143A")
	copy(key, bz)
	simSecp256k1Pubkey.Key = key
}

// SignatureVerificationGasConsumer is the type of function that is used to both
// consume gas when verifying signatures and also to accept or reject different types of pubkeys
// This is where apps can define their own PubKey
type SignatureVerificationGasConsumer = func(meter sdk.GasMeter, sig signing.SignatureV2, params types.Params) error

// SetPubKeyDecorator sets PubKeys in context for any signer which does not already have pubkey set
// PubKeys must be set in context for all signers before any other sigverify decorators run
// CONTRACT: Tx must implement SigVerifiableTx interface
type SetPubKeyDecorator struct {
	ak AccountKeeper
}

func NewSetPubKeyDecorator(ak AccountKeeper) SetPubKeyDecorator {
	_ = "STUB: not implemented"
	return *new(SetPubKeyDecorator)
}

func (spkd SetPubKeyDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}

// PublicKey was omitted from slice since it has already been set in context

// Only make check if simulate=false

// account already has pubkey set,no need to reset

// Also emit the following events, so that txs can be indexed by these
// indices:
// - signature (via `tx.signature='<sig_as_base64>'`),
// - concat(address,"/",sequence) (via `tx.acc_seq='cosmos1abc...def/42'`).

// Consume parameter-defined amount of gas for each signature according to the passed-in SignatureVerificationGasConsumer function
// before calling the next AnteHandler
// CONTRACT: Pubkeys are set in context for all signers before this decorator runs
// CONTRACT: Tx must implement SigVerifiableTx interface
type SigGasConsumeDecorator struct {
	ak             AccountKeeper
	sigGasConsumer SignatureVerificationGasConsumer
}

func NewSigGasConsumeDecorator(ak AccountKeeper, sigGasConsumer SignatureVerificationGasConsumer) SigGasConsumeDecorator {
	_ = "STUB: not implemented"
	return *new(SigGasConsumeDecorator)
}

func (sgcd SigGasConsumeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}

// stdSigs contains the sequence number, account number, and signatures.
// When simulating, this would just be a 0-length slice.

// In simulate mode the transaction comes with no signatures, thus if the
// account's pubkey is nil, both signature verification and gasKVStore.Set()
// shall consume the largest amount, i.e. it takes more gas to verify
// secp256k1 keys than ed25519 ones.

// make a SignatureV2 with PubKey filled in from above

// Verify all signatures for a tx and return an error if any are invalid. Note,
// the SigVerificationDecorator will not check signatures on ReCheck.
//
// CONTRACT: Pubkeys are set in context for all signers before this decorator runs
// CONTRACT: Tx must implement SigVerifiableTx interface
type SigVerificationDecorator struct {
	ak              AccountKeeper
	signModeHandler authsigning.SignModeHandler
}

func NewSigVerificationDecorator(ak AccountKeeper, signModeHandler authsigning.SignModeHandler) SigVerificationDecorator {
	_ = "STUB: not implemented"
	return *new(SigVerificationDecorator)
}

// OnlyLegacyAminoSigners checks SignatureData to see if all
// signers are using SIGN_MODE_LEGACY_AMINO_JSON. If this is the case
// then the corresponding SignatureV2 struct will not have account sequence
// explicitly set, and we should skip the explicit verification of sig.Sequence
// in the SigVerificationDecorator's AnteHandler function.
func OnlyLegacyAminoSigners(sigData signing.SignatureData) bool {
	_ = "STUB: not implemented"
	return false
}

func (svd SigVerificationDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}

// stdSigs contains the sequence number, account number, and signatures.
// When simulating, this would just be a 0-length slice.

// check that signer length and signature length are the same

// retrieve pubkey

// Check account sequence number.

// retrieve signer data

// no need to verify signatures on recheck tx

// If all signers are using SIGN_MODE_LEGACY_AMINO, we rely on VerifySignature to check account sequence number,
// and therefore communicate sequence number as a potential cause of error.

// IncrementSequenceDecorator handles incrementing sequences of all signers.
// Use the IncrementSequenceDecorator decorator to prevent replay attacks. Note,
// there is no need to execute IncrementSequenceDecorator on RecheckTX since
// CheckTx would already bump the sequence number.
//
// NOTE: Since CheckTx and DeliverTx state are managed separately, subsequent and
// sequential txs orginating from the same account cannot be handled correctly in
// a reliable way unless sequence numbers are managed and tracked manually by a
// client. It is recommended to instead use multiple messages in a tx.
type IncrementSequenceDecorator struct {
	ak AccountKeeper
}

func NewIncrementSequenceDecorator(ak AccountKeeper) IncrementSequenceDecorator {
	_ = "STUB: not implemented"
	return *new(IncrementSequenceDecorator)
}

func (isd IncrementSequenceDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}

// increment sequence of all signers

// ValidateSigCountDecorator takes in Params and returns errors if there are too many signatures in the tx for the given params
// otherwise it calls next AnteHandler
// Use this decorator to set parameterized limit on number of signatures in tx
// CONTRACT: Tx must implement SigVerifiableTx interface
type ValidateSigCountDecorator struct {
	ak AccountKeeper
}

func NewValidateSigCountDecorator(ak AccountKeeper) ValidateSigCountDecorator {
	_ = "STUB: not implemented"
	return *new(ValidateSigCountDecorator)
}

func (vscd ValidateSigCountDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}

//nolint:gosec // sigCount is incremented from 0 and checked each iteration, always non-negative

// DefaultSigVerificationGasConsumer is the default implementation of SignatureVerificationGasConsumer. It consumes gas
// for signature verification based upon the public key type. The cost is fetched from the given params and is matched
// by the concrete type.
func DefaultSigVerificationGasConsumer(
	meter sdk.GasMeter, sig signing.SignatureV2, params types.Params,
) error {
	_ = "STUB: not implemented"
	return nil
}

// ConsumeMultisignatureVerificationGas consumes gas from a GasMeter for verifying a multisig pubkey signature
func ConsumeMultisignatureVerificationGas(
	meter sdk.GasMeter, sig *signing.MultiSignatureData, pubkey multisig.PubKey,
	params types.Params, accSeq uint64,
) error {
	_ = "STUB: not implemented"
	return nil
}

// GetSignerAcc returns an account for a given address that is expected to sign
// a transaction.
func GetSignerAcc(ctx sdk.Context, ak AccountKeeper, addr sdk.AccAddress) (types.AccountI, error) {
	_ = "STUB: not implemented"
	return *new(types.AccountI), nil
}

// CountSubKeys counts the total number of keys for a multi-sig public key.
func CountSubKeys(pub cryptotypes.PubKey) int { _ = "STUB: not implemented"; return 0 }

// SignatureDataToBz converts a SignatureData into raw bytes signature.
// For SingleSignatureData, it returns the signature raw bytes.
// For MultiSignatureData, it returns an array of all individual signatures,
// as well as the aggregated signature.
func SignatureDataToBz(data signing.SignatureData) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
