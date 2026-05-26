package types

import (
	"time"

	ics23 "github.com/confio/ics23/go"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	clienttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/types"
	commitmenttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/23-commitment/types"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

var _ exported.ClientState = (*ClientState)(nil)

// NewClientState creates a new ClientState instance
func NewClientState(
	chainID string, trustLevel Fraction,
	trustingPeriod, ubdPeriod, maxClockDrift time.Duration,
	latestHeight clienttypes.Height, specs []*ics23.ProofSpec,
	upgradePath []string, allowUpdateAfterExpiry, allowUpdateAfterMisbehaviour bool,
) *ClientState {
	_ = "STUB: not implemented"
	return nil
}

// GetChainID returns the chain-id
func (cs ClientState) GetChainID() string {
	_ = "STUB: not implemented"

	// ClientType is tendermint.
	return ""
}

func (cs ClientState) ClientType() string { _ = "STUB: not implemented"; return "" }

// GetLatestHeight returns latest block height.
func (cs ClientState) GetLatestHeight() exported.Height {
	_ = "STUB: not implemented"
	return *

	// Status returns the status of the tendermint client.
	// The client may be:
	// - Active: FrozenHeight is zero and client is not expired
	// - Frozen: Frozen Height is not zero
	// - Expired: the latest consensus state timestamp + trusting period <= current time
	//
	// A frozen client will become expired, so the Frozen status
	// has higher precedence.
	new(exported.Height)
}

func (cs ClientState) Status(
	ctx sdk.Context,
	clientStore sdk.KVStore,
	cdc codec.BinaryCodec,
) exported.Status {
	_ = "STUB: not implemented"
	return *new(exported.Status)
}

// get latest consensus state from clientStore to check for expiry

// if the client state does not have an associated consensus state for its latest height
// then it must be expired

// IsExpired returns whether or not the client has passed the trusting period since the last
// update (in which case no headers are considered valid).
func (cs ClientState) IsExpired(latestTimestamp, now time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

// Validate performs a basic validation of the client state fields.
func (cs ClientState) Validate() error { _ = "STUB: not implemented"; return nil }

// NOTE: the value of tmtypes.MaxChainIDLen may change in the future.
// If this occurs, the code here must account for potential difference
// between the tendermint version being run by the counterparty chain
// and the tendermint version used by this light client.
// https://github.com/cosmos/ibc-go/issues/177

// the latest height revision number must match the chain id revision number

// UpgradePath may be empty, but if it isn't, each key must be non-empty

// GetProofSpecs returns the format the client expects for proof verification
// as a string array specifying the proof type for each position in chained proof
func (cs ClientState) GetProofSpecs() []*ics23.ProofSpec { _ = "STUB: not implemented"; return nil }

// ZeroCustomFields returns a ClientState that is a copy of the current ClientState
// with all client customizable fields zeroed out
func (cs ClientState) ZeroCustomFields() exported.ClientState {
	_ = "STUB: not implemented"
	// copy over all chain-specified fields
	// and leave custom fields empty
	return *new(exported.ClientState)
}

// Initialize will check that initial consensus state is a Tendermint consensus state
// and will store ProcessedTime for initial consensus state as ctx.BlockTime()
func (cs ClientState) Initialize(ctx sdk.Context, _ codec.BinaryCodec, clientStore sdk.KVStore, consState exported.ConsensusState) error {
	_ = "STUB: not implemented"
	return nil
}

// set metadata for initial consensus state.

// VerifyClientState verifies a proof of the client state of the running chain
// stored on the target machine
func (cs ClientState) VerifyClientState(
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	prefix exported.Prefix,
	counterpartyClientIdentifier string,
	proof []byte,
	clientState exported.ClientState,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyClientConsensusState verifies a proof of the consensus state of the
// Tendermint client stored on the target machine.
func (cs ClientState) VerifyClientConsensusState(
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	counterpartyClientIdentifier string,
	consensusHeight exported.Height,
	prefix exported.Prefix,
	proof []byte,
	consensusState exported.ConsensusState,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyConnectionState verifies a proof of the connection state of the
// specified connection end stored on the target machine.
func (cs ClientState) VerifyConnectionState(
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	prefix exported.Prefix,
	proof []byte,
	connectionID string,
	connectionEnd exported.ConnectionI,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyChannelState verifies a proof of the channel state of the specified
// channel end, under the specified port, stored on the target machine.
func (cs ClientState) VerifyChannelState(
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	prefix exported.Prefix,
	proof []byte,
	portID,
	channelID string,
	channel exported.ChannelI,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyPacketCommitment verifies a proof of an outgoing packet commitment at
// the specified port, specified channel, and specified sequence.
func (cs ClientState) VerifyPacketCommitment(
	ctx sdk.Context,
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	delayTimePeriod uint64,
	delayBlockPeriod uint64,
	prefix exported.Prefix,
	proof []byte,
	portID,
	channelID string,
	sequence uint64,
	commitmentBytes []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

// check delay period has passed

// VerifyPacketAcknowledgement verifies a proof of an incoming packet
// acknowledgement at the specified port, specified channel, and specified sequence.
func (cs ClientState) VerifyPacketAcknowledgement(
	ctx sdk.Context,
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	delayTimePeriod uint64,
	delayBlockPeriod uint64,
	prefix exported.Prefix,
	proof []byte,
	portID,
	channelID string,
	sequence uint64,
	acknowledgement []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

// check delay period has passed

// VerifyPacketReceiptAbsence verifies a proof of the absence of an
// incoming packet receipt at the specified port, specified channel, and
// specified sequence.
func (cs ClientState) VerifyPacketReceiptAbsence(
	ctx sdk.Context,
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	delayTimePeriod uint64,
	delayBlockPeriod uint64,
	prefix exported.Prefix,
	proof []byte,
	portID,
	channelID string,
	sequence uint64,
) error {
	_ = "STUB: not implemented"
	return nil
}

// check delay period has passed

// VerifyNextSequenceRecv verifies a proof of the next sequence number to be
// received of the specified channel at the specified port.
func (cs ClientState) VerifyNextSequenceRecv(
	ctx sdk.Context,
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	delayTimePeriod uint64,
	delayBlockPeriod uint64,
	prefix exported.Prefix,
	proof []byte,
	portID,
	channelID string,
	nextSequenceRecv uint64,
) error {
	_ = "STUB: not implemented"
	return nil
}

// check delay period has passed

// verifyDelayPeriodPassed will ensure that at least delayTimePeriod amount of time and delayBlockPeriod number of blocks have passed
// since consensus state was submitted before allowing verification to continue.
func verifyDelayPeriodPassed(ctx sdk.Context, store sdk.KVStore, proofHeight exported.Height, delayTimePeriod, delayBlockPeriod uint64) error {
	_ = "STUB: not implemented"
	// check that executing chain's timestamp has passed consensusState's processed time + delay time period
	return nil
}

//nolint:gosec // overflow checked above

// NOTE: delay time period is inclusive, so if currentTimestamp is validTime, then we return no error

// check that executing chain's height has passed consensusState's processed height + delay block period

// NOTE: delay block period is inclusive, so if currentHeight is validHeight, then we return no error

// produceVerificationArgs perfoms the basic checks on the arguments that are
// shared between the verification functions and returns the unmarshalled
// merkle proof, the consensus state and an error if one occurred.
func produceVerificationArgs(
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	cs ClientState,
	height exported.Height,
	prefix exported.Prefix,
	proof []byte,
) (merkleProof commitmenttypes.MerkleProof, consensusState *ConsensusState, err error) {
	_ = "STUB: not implemented"
	return *new(commitmenttypes.MerkleProof), nil, nil
}
