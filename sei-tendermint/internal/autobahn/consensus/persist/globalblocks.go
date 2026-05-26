package persist

import (
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

const globalBlocksDir = "globalblocks"

// LoadedGlobalBlock is a block loaded from disk during state restoration.
// Also used as the internal WAL entry type. Block doesn't carry its
// GlobalBlockNumber (that's assigned by the ordering layer), so we embed
// it in each WAL entry to make entries self-describing.
type LoadedGlobalBlock struct {
	Number types.GlobalBlockNumber
	Block  *types.Block
}

// loadedGlobalBlockCodec serializes LoadedGlobalBlock as [8-byte LE number][proto block].
type loadedGlobalBlockCodec struct{}

func (loadedGlobalBlockCodec) Marshal(e LoadedGlobalBlock) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (loadedGlobalBlockCodec) Unmarshal(raw []byte) (LoadedGlobalBlock, error) {
	_ = "STUB: not implemented"
	return *new(LoadedGlobalBlock), nil
}

// globalBlockState is the mutable state protected by GlobalBlockPersister's mutex.
type globalBlockState struct {
	iw        utils.Option[*indexedWAL[LoadedGlobalBlock]]
	committee *types.Committee
	next      types.GlobalBlockNumber
	loaded    []LoadedGlobalBlock
}

func (s *globalBlockState) persistBlock(n types.GlobalBlockNumber, block *types.Block) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *globalBlockState) truncateBefore(n types.GlobalBlockNumber) error {
	_ = "STUB: not implemented"
	return nil
}

// GlobalBlockPersister manages persistence of globally-ordered blocks using a WAL.
// Each entry embeds its GlobalBlockNumber since Block doesn't carry it.
// When stateDir is None, all disk I/O is skipped (no-op mode).
// All public methods are safe for concurrent use.
type GlobalBlockPersister struct {
	state utils.Mutex[*globalBlockState]
}

// NewGlobalBlockPersister opens (or creates) a WAL in the globalblocks/ subdir
// and replays all persisted entries. Loaded blocks are available via
// ConsumeLoaded. When stateDir is None, returns a no-op persister.
func NewGlobalBlockPersister(stateDir utils.Option[string], committee *types.Committee) (*GlobalBlockPersister, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: avoid loading all blocks on startup; cache only the last N blocks
// (e.g. 1000) in memory instead.

// Next returns the next GlobalBlockNumber expected by the persister.
func (gp *GlobalBlockPersister) Next() types.GlobalBlockNumber {
	_ = "STUB: not implemented"
	return *new(types.GlobalBlockNumber)
}

// LoadedFirst returns the first loaded block number, or committee.FirstBlock() if empty.
func (gp *GlobalBlockPersister) LoadedFirst() types.GlobalBlockNumber {
	_ = "STUB: not implemented"
	return *new(types.GlobalBlockNumber)
}

// ConsumeLoaded returns blocks loaded from the WAL during construction
// and nils the internal slice so the data is not retained.
func (gp *GlobalBlockPersister) ConsumeLoaded() []LoadedGlobalBlock {
	_ = "STUB: not implemented"
	return nil
}

// PersistBlock appends a block to the WAL. Duplicates are silently ignored.
// Gaps return an error.
func (gp *GlobalBlockPersister) PersistBlock(n types.GlobalBlockNumber, block *types.Block) error {
	_ = "STUB: not implemented"
	return nil
}

// TruncateBefore removes all entries before n from the WAL.
func (gp *GlobalBlockPersister) TruncateBefore(n types.GlobalBlockNumber) error {
	_ = "STUB: not implemented"
	return nil
}

// TruncateAfter removes block entries >= n from the WAL, updates the
// persister cursor, and trims loaded data. Called by DataWAL.reconcile
// to remove blocks persisted without corresponding QCs.
func (gp *GlobalBlockPersister) TruncateAfter(n types.GlobalBlockNumber) error {
	_ = "STUB: not implemented"
	return nil
}

// Close shuts down the WAL.
func (gp *GlobalBlockPersister) Close() error { _ = "STUB: not implemented"; return nil }

func (s *globalBlockState) loadAll() ([]LoadedGlobalBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
