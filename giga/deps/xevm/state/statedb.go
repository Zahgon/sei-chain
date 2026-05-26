package state

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/stateless"
	"github.com/ethereum/go-ethereum/core/tracing"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	ethutils "github.com/ethereum/go-ethereum/trie/utils"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("giga", "deps", "xevm", "state")

// Initialized for each transaction individually
type DBImpl struct {
	// ctx is the single CacheMultiStore context used for all KV mutations within this stateDB.
	ctx sdk.Context
	// committedCtx is the pre-stateDB context, used for GetCommittedState reads and event flushing.
	committedCtx sdk.Context

	// validRevisions tracks snapshot points (journal index) for RevertToSnapshot.
	validRevisions []revision
	nextRevisionId int

	// snapshottedEventManagers holds EMs from prior snapshots that survived (not reverted).
	snapshottedEventManagers []*sdk.EventManager

	tempState *TemporaryState
	journal   []journalEntry

	// If err is not nil at the end of the execution, the transaction will be rolled
	// back.
	err error
	// whenever this is set, the same error would also cause EVM to revert, which is
	// why we don't put it in `tempState`, since we still want to be able to access it later.
	precompileErr error

	// a temporary address that collects fees for this particular transaction so that there is
	// no single bottleneck for fee collection. Its account state and balance will be deleted
	// before the block commits
	coinbaseAddress    sdk.AccAddress
	coinbaseEvmAddress common.Address

	k          EVMKeeper
	simulation bool

	// for cases like bank.send_native, we want to suppress transfer events
	eventsSuppressed bool

	logger *tracing.Hooks
}

func NewDBImpl(ctx sdk.Context, k EVMKeeper, simulation bool) *DBImpl {
	_ = "STUB: not implemented"
	return nil
}

// Create a single CacheMultiStore layer for all KV mutations within this stateDB.

func (s *DBImpl) DisableEvents() { _ = "STUB: not implemented"; return }

func (s *DBImpl) EnableEvents() { _ = "STUB: not implemented"; return }

func (s *DBImpl) SetLogger(logger *tracing.Hooks) {
	_ = "STUB: not implemented"

	// for interface compliance
	return
}

func (s *DBImpl) SetEVM(evm *vm.EVM) {
	_ = "STUB: not implemented"

	// AddPreimage records a SHA3 preimage seen by the VM.
	// AddPreimage performs a no-op since the EnablePreimageRecording flag is disabled
	// on the vm.Config during state transitions. No store trie preimages are written
	// to the database.
	return
}

func (s *DBImpl) AddPreimage(_ common.Hash, _ []byte) { _ = "STUB: not implemented"; return }

func (s *DBImpl) Cleanup() { _ = "STUB: not implemented"; return }

func (s *DBImpl) CleanupForTracer() {
	_ = "STUB: not implemented"
	// Reset back to the committed (pre-stateDB) state by discarding the CMS layer.
	return
}

// Re-create the CMS layer for the tracer.

// ResetForTracer resets in-memory state for a new transaction without flushing
// the CacheMultiStore hierarchy. This is safe for concurrent use when copies of
// this statedb are being read from other goroutines, since it never calls
// CacheMultiStore.Write() on any shared store layer.
func (s *DBImpl) ResetForTracer() { _ = "STUB: not implemented"; return }

func (s *DBImpl) Finalize() (surplus sdk.Int, err error) {
	_ = "STUB: not implemented"
	return *new(sdk.Int), nil
}

// delete state of self-destructed accounts

// Write the single CMS layer to the underlying store.

// Emit all surviving events (from snapshots + current) to the committed ctx's EventManager.

// Backward-compatibility functions
func (s *DBImpl) Error() error { _ = "STUB: not implemented"; return nil }

func (s *DBImpl) GetStorageRoot(common.Address) common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
}

func (s *DBImpl) Copy() vm.StateDB { _ = "STUB: not implemented"; return *new(vm.StateDB) }

func (s *DBImpl) Finalise(bool) { _ = "STUB: not implemented"; return }

func (s *DBImpl) Commit(uint64, bool, bool) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

func (s *DBImpl) SetTxContext(common.Hash, int) {
	_ = "STUB: not implemented"
	// noop
	return
}

func (s *DBImpl) AccessEvents() *vm.AccessEvents {
	_ = "STUB: not implemented"

	// CreateContract marks the account as created for EIP-6780 purposes.
	// This is called regardless of whether the account previously existed
	// (e.g., prefunded addresses), ensuring that contracts created and
	// self-destructed in the same transaction are properly destroyed.
	return nil
}

func (s *DBImpl) CreateContract(acc common.Address) { _ = "STUB: not implemented"; return }

func (s *DBImpl) PointCache() *ethutils.PointCache { _ = "STUB: not implemented"; return nil }

func (s *DBImpl) Witness() *stateless.Witness { _ = "STUB: not implemented"; return nil }

func (s *DBImpl) IntermediateRoot(bool) common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
}

func (s *DBImpl) TxIndex() int { _ = "STUB: not implemented"; return 0 }

func (s *DBImpl) Preimages() map[common.Hash][]byte { _ = "STUB: not implemented"; return nil }

func (s *DBImpl) SetPrecompileError(err error) { _ = "STUB: not implemented"; return }

func (s *DBImpl) GetPrecompileError() error { _ = "STUB: not implemented"; return nil }

// ** TEST ONLY FUNCTIONS **//
func (s *DBImpl) Err() error { _ = "STUB: not implemented"; return nil }

func (s *DBImpl) WithErr(err error) { _ = "STUB: not implemented"; return }

func (s *DBImpl) Ctx() sdk.Context { _ = "STUB: not implemented"; return *new(sdk.Context) }

func (s *DBImpl) WithCtx(ctx sdk.Context) {
	_ = "STUB: not implemented"

	// in-memory state that's generated by a specific
	// EVM snapshot in a single transaction
	return
}

type TemporaryState struct {
	logs                  []*ethtypes.Log
	transientStates       map[string]map[string]common.Hash
	transientAccounts     map[string][]byte
	transientModuleStates map[string][]byte
	transientAccessLists  *accessList
	surplus               sdk.Int // in wei
}

func NewTemporaryState() *TemporaryState { _ = "STUB: not implemented"; return nil }

func (ts *TemporaryState) DeepCopy() *TemporaryState { _ = "STUB: not implemented"; return nil }

func GetDBImpl(vmsdb vm.StateDB) *DBImpl { _ = "STUB: not implemented"; return nil }
