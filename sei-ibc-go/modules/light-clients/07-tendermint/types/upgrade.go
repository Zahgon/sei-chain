package types

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	commitmenttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/23-commitment/types"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

// VerifyUpgradeAndUpdateState checks if the upgraded client has been committed by the current client
// It will zero out all client-specific fields (e.g. TrustingPeriod and verify all data
// in client state that must be the same across all valid Tendermint clients for the new chain.
// VerifyUpgrade will return an error if:
// - the upgradedClient is not a Tendermint ClientState
// - the lastest height of the client state does not have the same revision number or has a greater
// height than the committed client.
//   - the height of upgraded client is not greater than that of current client
//   - the latest height of the new client does not match or is greater than the height in committed client
//   - any Tendermint chain specified parameter in upgraded client such as ChainID, UnbondingPeriod,
//     and ProofSpecs do not match parameters set by committed client
func (cs ClientState) VerifyUpgradeAndUpdateState(
	ctx sdk.Context, cdc codec.BinaryCodec, clientStore sdk.KVStore,
	upgradedClient exported.ClientState, upgradedConsState exported.ConsensusState,
	proofUpgradeClient, proofUpgradeConsState []byte,
) (exported.ClientState, exported.ConsensusState, error) {
	_ = "STUB: not implemented"
	return *new(exported.ClientState), *new(exported.ConsensusState), nil
}

// last height of current counterparty chain must be client's latest height

// upgraded client state and consensus state must be IBC tendermint client state and consensus state
// this may be modified in the future to upgrade to a new IBC tendermint type
// counterparty must also commit to the upgraded consensus state at a sub-path under the upgrade path specified

// unmarshal proofs

// Must prove against latest consensus state to ensure we are verifying against latest upgrade plan
// This verifies that upgrade is intended for the provided revision, since committed client must exist
// at this consensus state

// Verify client proof

// construct clientState Merkle path

// Verify consensus state proof

// construct consensus state Merkle path

// Construct new client state and consensus state
// Relayer chosen client parameters are ignored.
// All chain-chosen parameters come from committed client, all client-chosen parameters
// come from current client.

// The new consensus state is merely used as a trusted kernel against which headers on the new
// chain can be verified. The root is just a stand-in sentinel value as it cannot be known in advance, thus no proof verification will pass.
// The timestamp and the NextValidatorsHash of the consensus state is the blocktime and NextValidatorsHash
// of the last block committed by the old chain. This will allow the first block of the new chain to be verified against
// the last validators of the old chain so long as it is submitted within the TrustingPeriod of this client.
// NOTE: We do not set processed time for this consensus state since this consensus state should not be used for packet verification
// as the root is empty. The next consensus state submitted using update will be usable for packet-verification.

// set metadata for this consensus state

// construct MerklePath for the committed client from upgradePath
func constructUpgradeClientMerklePath(upgradePath []string, lastHeight exported.Height) commitmenttypes.MerklePath {
	_ = "STUB: not implemented"
	// copy all elements from upgradePath except final element
	return *new(commitmenttypes.MerklePath)
}

// append lastHeight and `upgradedClient` to last key of upgradePath and use as lastKey of clientPath
// this will create the IAVL key that is used to store client in upgrade store

// construct MerklePath for the committed consensus state from upgradePath
func constructUpgradeConsStateMerklePath(upgradePath []string, lastHeight exported.Height) commitmenttypes.MerklePath {
	_ = "STUB: not implemented"
	// copy all elements from upgradePath except final element
	return *new(commitmenttypes.MerklePath)
}

// append lastHeight and `upgradedClient` to last key of upgradePath and use as lastKey of clientPath
// this will create the IAVL key that is used to store client in upgrade store
