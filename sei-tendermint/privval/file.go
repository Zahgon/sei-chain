package privval

import (
	"context"
	"encoding/json"
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto"
	tmbytes "github.com/sei-protocol/sei-chain/sei-tendermint/libs/bytes"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

// TODO: type ?
const (
	stepNone      int8 = 0 // Used to distinguish the initial state
	stepPropose   int8 = 1
	stepPrevote   int8 = 2
	stepPrecommit int8 = 3
)

// A vote is either stepPrevote or stepPrecommit.
func voteToStep(vote *tmproto.Vote) (int8, error) { _ = "STUB: not implemented"; return 0, nil }

//-------------------------------------------------------------------------------

// FilePVKey stores the immutable part of PrivValidator.
type FilePVKey struct {
	Address types.Address
	PubKey  crypto.PubKey
	PrivKey crypto.PrivKey

	filePath string
}

type filePVKeyJSON struct {
	Address types.Address   `json:"address"`
	PubKey  json.RawMessage `json:"pub_key"`
	PrivKey json.RawMessage `json:"priv_key"`
}

func (pvKey FilePVKey) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (pvKey *FilePVKey) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// Save persists the FilePVKey to its filePath.
func (pvKey FilePVKey) Save() error { _ = "STUB: not implemented"; return nil }

// Write pubkey in autobahn-compatible format alongside the key file.
// TODO: use atypes.PublicKey.String() directly to avoid duplicating the "validator:" prefix.

//-------------------------------------------------------------------------------

// FilePVLastSignState stores the mutable part of PrivValidator.
type FilePVLastSignState struct {
	Height    int64            `json:"height,string"`
	Round     int32            `json:"round"`
	Step      int8             `json:"step"`
	Signature []byte           `json:"signature,omitempty"`
	SignBytes tmbytes.HexBytes `json:"signbytes,omitempty"`

	filePath string
}

func (lss *FilePVLastSignState) reset() { _ = "STUB: not implemented"; return }

// checkHRS checks the given height, round, step (HRS) against that of the
// FilePVLastSignState. It returns an error if the arguments constitute a regression,
// or if they match but the SignBytes are empty.
// The returned boolean indicates whether the last Signature should be reused -
// it returns true if the HRS matches the arguments and the SignBytes are not empty (indicating
// we have already signed for this HRS, and can reuse the existing signature).
// It panics if the HRS matches the arguments, there's a SignBytes, but no Signature.
func (lss *FilePVLastSignState) checkHRS(height int64, round int32, step int8) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Save persists the FilePvLastSignState to its filePath.
func (lss *FilePVLastSignState) Save() error { _ = "STUB: not implemented"; return nil }

//-------------------------------------------------------------------------------

// FilePV implements PrivValidator using data persisted to disk
// to prevent double signing.
// NOTE: the directories containing pv.Key.filePath and pv.LastSignState.filePath must already exist.
// It includes the LastSignature and LastSignBytes so we don't lose the signature
// if the process crashes after signing but before the resulting consensus message is processed.
type FilePV struct {
	Key           FilePVKey
	LastSignState FilePVLastSignState
}

var _ types.PrivValidator = (*FilePV)(nil)

// NewFilePV generates a new validator from the given key and paths.
func NewFilePV(privKey crypto.PrivKey, keyFilePath, stateFilePath string) *FilePV {
	_ = "STUB: not implemented"
	return nil
}

// GenFilePV generates a new validator with randomly generated private key
// and sets the filePaths, but does not call Save().
func GenFilePV(keyFilePath, stateFilePath, keyType string) (*FilePV, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LoadFilePV loads a FilePV from the filePaths.  The FilePV handles double
// signing prevention by persisting data to the stateFilePath.  If either file path
// does not exist, the program will exit.
func LoadFilePV(keyFilePath, stateFilePath string) (*FilePV, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LoadFilePVEmptyState loads a FilePV from the given keyFilePath, with an empty LastSignState.
// If the keyFilePath does not exist, the program will exit.
func LoadFilePVEmptyState(keyFilePath, stateFilePath string) (*FilePV, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If loadState is true, we load from the stateFilePath. Otherwise, we use an empty LastSignState.
func loadFilePV(keyFilePath, stateFilePath string, loadState bool) (*FilePV, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// overwrite pubkey and address for convenience

// LoadOrGenFilePV loads a FilePV from the given filePaths
// or else generates a new one and saves it to the filePaths.
func LoadOrGenFilePV(keyFilePath, stateFilePath string) (*FilePV, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAddress returns the address of the validator.
// Implements PrivValidator.
func (pv *FilePV) GetAddress() types.Address {
	_ = "STUB: not implemented"
	return *

	// GetPubKey returns the public key of the validator.
	// Implements PrivValidator.
	new(types.Address)
}

func (pv *FilePV) GetPubKey(ctx context.Context) (crypto.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PubKey), nil
}

// SignVote signs a canonical representation of the vote, along with the
// chainID. Implements PrivValidator.
func (pv *FilePV) SignVote(ctx context.Context, chainID string, vote *tmproto.Vote) error {
	_ = "STUB: not implemented"
	return nil
}

// SignProposal signs a canonical representation of the proposal, along with
// the chainID. Implements PrivValidator.
func (pv *FilePV) SignProposal(ctx context.Context, chainID string, proposal *tmproto.Proposal) error {
	_ = "STUB: not implemented"
	return nil
}

// Save persists the FilePV to disk.
func (pv *FilePV) Save() error { _ = "STUB: not implemented"; return nil }

// Reset resets all fields in the FilePV.
// NOTE: Unsafe!
func (pv *FilePV) Reset() error { _ = "STUB: not implemented"; return nil }

// String returns a string representation of the FilePV.
func (pv *FilePV) String() string { _ = "STUB: not implemented"; return "" }

//------------------------------------------------------------------------------------

// signVote checks if the vote is good to sign and sets the vote signature.
// It may need to set the timestamp as well if the vote is otherwise the same as
// a previously signed vote (ie. we crashed after signing but before the vote hit the WAL).
func (pv *FilePV) signVote(chainID string, vote *tmproto.Vote) error {
	_ = "STUB: not implemented"
	return nil
}

// We might crash before writing to the wal,
// causing us to try to re-sign for the same HRS.
// If signbytes are the same, use the last signature.
// If they only differ by timestamp, use last timestamp and signature
// Otherwise, return error

// Compares the canonicalized votes (i.e. without vote extensions
// or vote extension signatures).

// It passed the checks. Sign the vote

// signProposal checks if the proposal is good to sign and sets the proposal signature.
func (pv *FilePV) signProposal(chainID string, proposal *tmproto.Proposal) error {
	_ = "STUB: not implemented"
	return nil
}

// We might crash before writing to the wal,
// causing us to try to re-sign for the same HRS.
// If signbytes are the same, use the last signature.

// It passed the checks. Sign the proposal

// Persist height/round/step and signature
func (pv *FilePV) saveSigned(height int64, round int32, step int8, signBytes []byte, sig []byte) error {
	_ = "STUB: not implemented"
	return nil
}

//-----------------------------------------------------------------------------------------

// Returns the timestamp from the lastSignBytes.
// Returns true if the only difference in the votes is their timestamp.
// Performs these checks on the canonical votes (excluding the vote extension
// and vote extension signatures).
func checkVotesOnlyDifferByTimestamp(lastSignBytes, newSignBytes []byte) (time.Time, bool, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), false, nil
}

// set the times to the same value and check equality
