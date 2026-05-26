package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/seilog"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/03-connection/types"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

var logger = seilog.NewLogger("ibc-go", "modules", "core", "03-connection", "keeper")

// ConnOpenInit initialises a connection attempt on chain A. The generated connection identifier
// is returned.
//
// NOTE: Msg validation verifies the supplied identifiers and ensures that the counterparty
// connection identifier is empty.
func (k Keeper) ConnOpenInit(
	ctx sdk.Context,
	clientID string,
	counterparty types.Counterparty, // counterpartyPrefix, counterpartyClientIdentifier
	version *types.Version,
	delayPeriod uint64,
) (string, error) {
	_ = "STUB: not implemented"
	// outbound gating: disallow outbound connection inits when outbound disabled
	return "", nil
}

// connection defines chain A's ConnectionEnd

// ConnOpenTry relays notice of a connection attempt on chain A to chain B (this
// code is executed on chain B).
//
// NOTE:
//   - Here chain A acts as the counterparty
//   - Identifiers are checked on msg validation
func (k Keeper) ConnOpenTry(
	ctx sdk.Context,
	previousConnectionID string, // previousIdentifier
	counterparty types.Counterparty, // counterpartyConnectionIdentifier, counterpartyPrefix and counterpartyClientIdentifier
	delayPeriod uint64,
	clientID string, // clientID of chainA
	clientState exported.ClientState, // clientState that chainA has for chainB
	counterpartyVersions []exported.Version, // supported versions of chain A
	proofInit []byte, // proof that chainA stored connectionEnd in state (on ConnOpenInit)
	proofClient []byte, // proof that chainA stored a light client of chainB
	proofConsensus []byte, // proof that chainA stored chainB's consensus state at consensus height
	proofHeight exported.Height, // height at which relayer constructs proof of A storing connectionEnd in state
	consensusHeight exported.Height, // latest height of chain B which chain A has stored in its chain B client
) (string, error) {
	_ = "STUB: not implemented"
	// inbound gating: disallow inbound connection tries when inbound disabled
	return "", nil
}

// empty connection identifier indicates continuing a previous connection handshake

// ensure that the previous connection exists

// ensure that the existing connection's
// counterparty is chainA and connection is on INIT stage.
// Check that existing connection versions for initialized connection is equal to compatible
// versions for this chain.
// ensure that existing connection's delay period is the same as desired delay period.

// continue with previous connection

// generate a new connection

// validate client parameters of a chainB client stored on chainA

// expectedConnection defines Chain A's ConnectionEnd
// NOTE: chain A's counterparty is chain B (i.e where this code is executed)
// NOTE: chainA and chainB must have the same delay period

// chain B picks a version from Chain A's available versions that is compatible
// with Chain B's supported IBC versions. PickVersion will select the intersection
// of the supported versions and the counterparty versions.

// connection defines chain B's ConnectionEnd

// Check that ChainA committed expectedConnectionEnd to its state

// Check that ChainA stored the clientState provided in the msg

// Check that ChainA stored the correct ConsensusState of chainB at the given consensusHeight

// store connection in chainB state

// ConnOpenAck relays acceptance of a connection open attempt from chain B back
// to chain A (this code is executed on chain A).
//
// NOTE: Identifiers are checked on msg validation.
func (k Keeper) ConnOpenAck(
	ctx sdk.Context,
	connectionID string,
	clientState exported.ClientState, // client state for chainA on chainB
	version *types.Version, // version that ChainB chose in ConnOpenTry
	counterpartyConnectionID string,
	proofTry []byte, // proof that connectionEnd was added to ChainB state in ConnOpenTry
	proofClient []byte, // proof of client state on chainB for chainA
	proofConsensus []byte, // proof that chainB has stored ConsensusState of chainA on its client
	proofHeight exported.Height, // height that relayer constructed proofTry
	consensusHeight exported.Height, // latest height of chainA that chainB has stored on its chainA client
) error {
	_ = "STUB: not implemented"
	// Check that chainB client hasn't stored invalid height
	return nil
}

// Retrieve connection

// Verify the provided version against the previously set connection state

// connection on ChainA must be in INIT or TRYOPEN

// if the connection is INIT then the provided version must be supproted

// if the connection is in TRYOPEN then the version must be the only set version in the
// retreived connection state.

// validate client parameters of a chainA client stored on chainB

// Retrieve chainA's consensus state at consensusheight

// Ensure that ChainB stored expected connectionEnd in its state during ConnOpenTry

// Check that ChainB stored the clientState provided in the msg

// Ensure that ChainB has stored the correct ConsensusState for chainA at the consensusHeight

// Update connection state to Open

// ConnOpenConfirm confirms opening of a connection on chain A to chain B, after
// which the connection is open on both chains (this code is executed on chain B).
//
// NOTE: Identifiers are checked on msg validation.
func (k Keeper) ConnOpenConfirm(
	ctx sdk.Context,
	connectionID string,
	proofAck []byte, // proof that connection opened on ChainA during ConnOpenAck
	proofHeight exported.Height, // height that relayer constructed proofAck
) error {
	_ = "STUB: not implemented"
	// Retrieve connection
	return nil
}

// Check that connection state on ChainB is on state: TRYOPEN

// Check that connection on ChainA is open

// Update ChainB's connection to Open
