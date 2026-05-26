package ibctesting

import (
	"context"
	"testing"
	"time"

	"github.com/sei-protocol/sei-chain/sei-cosmos/baseapp"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	authtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/types"
	capabilitykeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/keeper"
	capabilitytypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/types"
	stakingkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/keeper"
	clienttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/types"
	channeltypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/04-channel/types"
	commitmenttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/23-commitment/types"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
	ibckeeper "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/keeper"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/types"
	ibctmtypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/light-clients/07-tendermint/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
	tmtypes "github.com/sei-protocol/sei-chain/sei-tendermint/types"

	wasmd "github.com/sei-protocol/sei-chain/sei-wasmd/app"
	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm"
)

// TestChain is a testing struct that wraps a simapp with the last TM Header, the current ABCI
// header and the validators of the TestChain. It also contains a field called ChainID. This
// is the clientID that *other* chains use to refer to this TestChain. The SenderAccount
// is used for delivering transactions through the application state.
// NOTE: the actual application uses an empty chain-id for ease of testing.
type TestChain struct {
	t *testing.T

	Coordinator   *Coordinator
	App           TestingApp
	ChainID       string
	LastHeader    *ibctmtypes.Header // header for last block height committed
	CurrentHeader tmproto.Header     // header for current block height
	QueryServer   types.QueryServer
	TxConfig      client.TxConfig
	Codec         codec.BinaryCodec

	Vals    *tmtypes.ValidatorSet
	Signers []tmtypes.PrivValidator

	senderPrivKey cryptotypes.PrivKey
	SenderAccount authtypes.AccountI

	PendingSendPackets []channeltypes.Packet
	PendingAckPackets  []PacketAck
}

type PacketAck struct {
	Packet channeltypes.Packet
	Ack    []byte
}

type PV struct {
	PrivKey cryptotypes.PrivKey
}

func NewPV() PV { _ = "STUB: not implemented"; return *new(PV) }

// GetPubKey implements PrivValidator interface
func (pv PV) GetPubKey(context.Context) (crypto.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PubKey), nil
}

// SignVote implements PrivValidator interface
func (pv PV) SignVote(_ context.Context, chainID string, vote *tmproto.Vote) error {
	_ = "STUB: not implemented"
	return nil
}

// SignProposal implements PrivValidator interface
func (pv PV) SignProposal(_ context.Context, chainID string, proposal *tmproto.Proposal) error {
	_ = "STUB: not implemented"
	return nil
}

// NewTestChain initializes a new TestChain instance with a single validator set using a
// generated private key. It also creates a sender account to be used for delivering transactions.
//
// The first block height is committed to state in order to allow for client creations on
// counterparty chains. The TestChain will return with a block height starting at 2.
//
// Time management is handled by the Coordinator in order to ensure synchrony between chains.
// Each update of any chain increments the block header time for all chains by 5 seconds.
func NewTestChain(t *testing.T, coord *Coordinator, chainID string, opts ...wasm.Option) *TestChain {
	_ = "STUB: not implemented"
	// generate validator private/public key
	return nil
}

// create validator set with single validator

// generate genesis account

// create current header and call begin block

// create an account to send transactions from

// GetContext returns the current context for the application.
func (chain *TestChain) GetContext() sdk.Context {
	_ = "STUB: not implemented"
	return *new(sdk.Context)
}

// QueryProof performs an abci query with the given key and returns the proto encoded merkle proof
// for the query and the height at which the proof will succeed on a tendermint verifier.
func (chain *TestChain) QueryProof(key []byte) ([]byte, clienttypes.Height) {
	_ = "STUB: not implemented"
	return nil, *new(clienttypes.Height)
}

// QueryProof performs an abci query with the given key and returns the proto encoded merkle proof
// for the query and the height at which the proof will succeed on a tendermint verifier.
func (chain *TestChain) QueryProofAtHeight(key []byte, height int64) ([]byte, clienttypes.Height) {
	_ = "STUB: not implemented"
	return nil, *new(clienttypes.Height)
}

// proof height + 1 is returned as the proof created corresponds to the height the proof
// was created in the IAVL tree. Tendermint and subsequently the clients that rely on it
// have heights 1 above the IAVL tree. Thus we return proof height + 1

// #nosec G115 -- checked above.

// QueryUpgradeProof performs an abci query with the given key and returns the proto encoded merkle proof
// for the query and the height at which the proof will succeed on a tendermint verifier.
func (chain *TestChain) QueryUpgradeProof(key []byte, height uint64) ([]byte, clienttypes.Height) {
	_ = "STUB: not implemented"
	return nil, *new(clienttypes.Height)
}

// #nosec G115 -- checked above.

// proof height + 1 is returned as the proof created corresponds to the height the proof
// was created in the IAVL tree. Tendermint and subsequently the clients that rely on it
// have heights 1 above the IAVL tree. Thus we return proof height + 1

// #nosec G115 -- checked above.

// QueryConsensusStateProof performs an abci query for a consensus state
// stored on the given clientID. The proof and consensusHeight are returned.
func (chain *TestChain) QueryConsensusStateProof(clientID string) ([]byte, clienttypes.Height) {
	_ = "STUB: not implemented"
	return nil, *new(clienttypes.Height)
}

// NextBlock sets the last header to the current header and increments the current header to be
// at the next block height. It does not update the time as that is handled by the Coordinator.
//
// CONTRACT: this function must only be called after app.Commit() occurs
func (chain *TestChain) NextBlock() {
	_ = "STUB: not implemented"
	// set the last header to the current header
	// use nil trusted fields
	return
}

// increment the current header

// NOTE: the time is increased by the coordinator to maintain time synchrony amongst
// chains.

// wasmApp.BeginBlock(wasmApp.GetContextForDeliverTx([]byte{}), abci.RequestBeginBlock{Header: chain.CurrentHeader})

// sendMsgs delivers a transaction through the application without returning the result.
func (chain *TestChain) sendMsgs(msgs ...sdk.Msg) error { _ = "STUB: not implemented"; return nil }

// SendMsgs delivers a transaction through the application. It updates the senders sequence
// number and updates the TestChain's headers. It returns the result and error if one
// occurred.
func (chain *TestChain) SendMsgs(msgs ...sdk.Msg) (*sdk.Result, error) {
	_ = "STUB: not implemented"
	// ensure the chain has the latest time
	return nil, nil
}

// SignAndDeliver calls app.Commit()

// increment sequence for successful transaction execution

func (chain *TestChain) captureIBCEvents(r *sdk.Result) { _ = "STUB: not implemented"; return }

// Keep a queue on the chain that we can relay in tests

// Keep a queue on the chain that we can relay in tests

// GetClientState retrieves the client state for the provided clientID. The client is
// expected to exist otherwise testing will fail.
func (chain *TestChain) GetClientState(clientID string) exported.ClientState {
	_ = "STUB: not implemented"
	return *new(exported.ClientState)
}

// GetConsensusState retrieves the consensus state for the provided clientID and height.
// It will return a success boolean depending on if consensus state exists or not.
func (chain *TestChain) GetConsensusState(clientID string, height exported.Height) (exported.ConsensusState, bool) {
	_ = "STUB: not implemented"
	return *new(exported.ConsensusState), false
}

// GetValsAtHeight will return the validator set of the chain at a given height. It will return
// a success boolean depending on if the validator set exists or not at that height.
func (chain *TestChain) GetValsAtHeight(height int64) (*tmtypes.ValidatorSet, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// GetAcknowledgement retrieves an acknowledgement for the provided packet. If the
// acknowledgement does not exist then testing will fail.
func (chain *TestChain) GetAcknowledgement(packet exported.PacketI) []byte {
	_ = "STUB: not implemented"
	return nil
}

// GetPrefix returns the prefix for used by a chain in connection creation
func (chain *TestChain) GetPrefix() commitmenttypes.MerklePrefix {
	_ = "STUB: not implemented"
	return *new(commitmenttypes.MerklePrefix)
}

// ConstructUpdateTMClientHeader will construct a valid 07-tendermint Header to update the
// light client on the source chain.
func (chain *TestChain) ConstructUpdateTMClientHeader(counterparty *TestChain, clientID string) (*ibctmtypes.Header, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConstructUpdateTMClientHeader will construct a valid 07-tendermint Header to update the
// light client on the source chain.
func (chain *TestChain) ConstructUpdateTMClientHeaderWithTrustedHeight(counterparty *TestChain, clientID string, trustedHeight clienttypes.Height) (*ibctmtypes.Header, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Relayer must query for LatestHeight on client to get TrustedHeight if the trusted height is not set

// Once we get TrustedHeight from client, we must query the validators from the counterparty chain
// If the LatestHeight == LastHeader.Height, then TrustedValidators are current validators
// If LatestHeight < LastHeader.Height, we can query the historical validator set from HistoricalInfo

// NOTE: We need to get validators from counterparty at height: trustedHeight+1
// since the last trusted validators for a header at height h
// is the NextValidators at h+1 committed to in header h by
// NextValidatorsHash

// #nosec G115 -- checked above

// inject trusted fields into last header
// for now assume revision number is 0

// ExpireClient fast forwards the chain's block time by the provided amount of time which will
// expire any clients with a trusting period less than or equal to this amount of time.
func (chain *TestChain) ExpireClient(amount time.Duration) { _ = "STUB: not implemented"; return }

// CurrentTMClientHeader creates a TM header using the current header parameters
// on the chain. The trusted fields in the header are set to nil.
func (chain *TestChain) CurrentTMClientHeader() *ibctmtypes.Header {
	_ = "STUB: not implemented"
	return nil
}

// CreateTMClientHeader creates a TM header to update the TM client. Args are passed in to allow
// caller flexibility to use params that differ from the chain.
func (chain *TestChain) CreateTMClientHeader(chainID string, blockHeight int64, trustedHeight clienttypes.Height, timestamp time.Time, tmValSet, tmTrustedVals *tmtypes.ValidatorSet, signers []tmtypes.PrivValidator) *ibctmtypes.Header {
	_ = "STUB: not implemented"
	return nil
}

//nolint:staticcheck

// #nosec G115 -- validator set size is checked above

// The trusted fields may be nil. They may be filled before relaying messages to a client.
// The relayer is responsible for querying client and injecting appropriate trusted fields.

// MakeBlockID copied unimported test functions from tmtypes to use them here
func MakeBlockID(hash []byte, partSetSize uint32, partSetHash []byte) tmtypes.BlockID {
	_ = "STUB: not implemented"
	return *new(tmtypes.BlockID)
}

// CreateSortedSignerArray takes two PrivValidators, and the corresponding Validator structs
// (including voting power). It returns a signer array of PrivValidators that matches the
// sorting of ValidatorSet.
// The sorting is first by .VotingPower (descending), with secondary index of .Address (ascending).
func CreateSortedSignerArray(altPrivVal, suitePrivVal tmtypes.PrivValidator,
	altVal, suiteVal *tmtypes.Validator,
) []tmtypes.PrivValidator {
	_ = "STUB: not implemented"
	return nil
}

// CreatePortCapability binds and claims a capability for the given portID if it does not
// already exist. This function will fail testing on any resulting error.
// NOTE: only creation of a capbility for a transfer or mock port is supported
// Other applications must bind to the port in InitGenesis or modify this code.
func (chain *TestChain) CreatePortCapability(scopedKeeper capabilitykeeper.ScopedKeeper, portID string) {
	_ = "STUB: not implemented"
	// ensure the chain has the latest time
	// check if the portId is already binded, if not bind it
	return
}

// create capability using the IBC capability keeper

// claim capability using the scopedKeeper

// GetPortCapability returns the port capability for the given portID. The capability must
// exist, otherwise testing will fail.
func (chain *TestChain) GetPortCapability(portID string) *capabilitytypes.Capability {
	_ = "STUB: not implemented"
	return nil
}

// CreateChannelCapability binds and claims a capability for the given portID and channelID
// if it does not already exist. This function will fail testing on any resulting error. The
// scoped keeper passed in will claim the new capability.
func (chain *TestChain) CreateChannelCapability(scopedKeeper capabilitykeeper.ScopedKeeper, portID, channelID string) {
	_ = "STUB: not implemented"
	// ensure the chain has the latest time
	return
}

// check if the portId is already binded, if not bind it

// GetChannelCapability returns the channel capability for the given portID and channelID.
// The capability must exist, otherwise testing will fail.
func (chain *TestChain) GetChannelCapability(portID, channelID string) *capabilitytypes.Capability {
	_ = "STUB: not implemented"
	return nil
}

func (chain *TestChain) Balance(acc sdk.AccAddress, denom string) sdk.Coin {
	_ = "STUB: not implemented"
	return *new(sdk.Coin)
}

func (chain *TestChain) AllBalances(acc sdk.AccAddress) sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

func (chain TestChain) GetTestSupport() *wasmd.TestSupport { _ = "STUB: not implemented"; return nil }

var _ TestingApp = TestingAppDecorator{}

type TestingAppDecorator struct {
	*wasmd.WasmApp
	t *testing.T
}

func NewTestingAppDecorator(t *testing.T, wasmApp *wasmd.WasmApp) *TestingAppDecorator {
	_ = "STUB: not implemented"
	return nil
}

func (a TestingAppDecorator) GetBaseApp() *baseapp.BaseApp { _ = "STUB: not implemented"; return nil }

func (a TestingAppDecorator) GetStakingKeeper() stakingkeeper.Keeper {
	_ = "STUB: not implemented"
	return *new(stakingkeeper.Keeper)
}

func (a TestingAppDecorator) GetIBCKeeper() *ibckeeper.Keeper {
	_ = "STUB: not implemented"
	return nil
}

func (a TestingAppDecorator) GetScopedIBCKeeper() capabilitykeeper.ScopedKeeper {
	_ = "STUB: not implemented"
	return *new(capabilitykeeper.ScopedKeeper)
}

func (a TestingAppDecorator) GetTxConfig() client.TxConfig {
	_ = "STUB: not implemented"
	return *new(client.TxConfig)
}

func (a TestingAppDecorator) TestSupport() *wasmd.TestSupport {
	_ = "STUB: not implemented"
	return nil
}
