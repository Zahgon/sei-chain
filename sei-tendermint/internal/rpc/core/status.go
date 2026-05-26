package core

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

// Status returns Tendermint status including node info, pubkey, latest block
// hash, app hash, block height, current max peer block height, and time.
// More: https://docs.tendermint.com/master/rpc/#/Info/status
//
// TODO(autobahn): several SyncInfo fields remain unpopulated under Autobahn
// because the CometBFT BlockStore / ConsensusReactor are not fed. Not
// blocking integration-test CI today (no reader), but external tooling will
// want them:
//
//   - LatestBlockHash / LatestBlockTime: need Autobahn block-header lookup
//     for the latest global block.
//   - EarliestBlockHash / EarliestAppHash / EarliestBlockHeight /
//     EarliestBlockTime: need Autobahn's pruning-boundary metadata
//     (FirstCommitQC or equivalent).
//   - CatchingUp: currently hardcoded `true` under Autobahn because
//     ConsensusReactor is nil. Needs a sync-state signal from Autobahn.
//   - MaxPeerBlockHeight / TotalSyncedTime / RemainingTime: require an
//     Autobahn equivalent of BlockSyncReactor.
//
// Separately, every RPC handler that uses env.getHeight(env.BlockStore.Height(), …)
// still rejects queries under Autobahn because BlockStore.Height()==0. This
// breaks /block, /block_results, /commit, and the evmrpc endpoints that walk
// through them (eth_getBlockByNumber, eth_gasPrice, etc.). A separate PR
// needs to either route block-data reads through an Autobahn-aware path or
// have Autobahn expose enough block metadata to satisfy the CometBFT Block
// RPC surface.
func (env *Environment) Status(ctx context.Context) (*coretypes.ResultStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Under Autobahn the CometBFT block store isn't populated, so the height
// above is stuck at 0. Pull the live height and app hash from the app —
// ABCIInfo reports both the last height the app committed in FinalizeBlock
// and the matching app hash. Block hash and block time stay empty; see
// the TODO on Status for what's left.
//
// LastCommittedBlockHeight reports the last block consensus has finalized.
// Under CometBFT commit == app-apply in one step, so latestHeight IS the
// committed height. Under Autobahn they can briefly differ; pull the
// consensus-committed number from the GigaRouter's cached CommitQC watch
// (lock-free per-call atomic load).

// Invariant: under Autobahn, consensus finalizes before the app
// executes, so the committed height leads (or equals) the executed
// height. Committed=0 is a legitimate startup window before the first
// CommitQC is loaded; skip the log there.

// Return the very last voting power, not the voting power of this validator
// during the last block.

// this should start as true, if consensus
// hasn't started yet, and then flip to false
// (or true,) depending on what's actually
// happening.

func (env *Environment) validatorAtHeight(h int64) utils.Option[*types.Validator] {
	_ = "STUB: not implemented"
	return nil
}

// Skip the in-memory consensus-state lookup under Autobahn: the CometBFT
// consensus State is never advanced, so GetValidators would nil-deref
// on an unpopulated validator set. The state-store lookup below is kept
// in sync under both engines.
