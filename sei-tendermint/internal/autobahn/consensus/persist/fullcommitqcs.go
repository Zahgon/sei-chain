package persist

import (
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

const fullCommitQCsDir = "fullcommitqcs"

// fullCommitQCState is the mutable state protected by FullCommitQCPersister's mutex.
type fullCommitQCState struct {
	iw        utils.Option[*indexedWAL[*types.FullCommitQC]]
	committee *types.Committee
	next      types.GlobalBlockNumber // next expected GlobalRange().First == last QC's GlobalRange().Next
	loaded    []*types.FullCommitQC
}

func (s *fullCommitQCState) persistQC(qc *types.FullCommitQC) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *fullCommitQCState) truncateBefore(n types.GlobalBlockNumber) error {
	_ = "STUB: not implemented"
	return nil
}

// Remove QCs whose range is fully before n. A QC is stale when
// GlobalRange().Next <= n. In practice the scan visits 0–1 entries
// per prune call because pruning advances one block at a time while
// each QC covers many blocks.

// FullCommitQCPersister manages persistence of FullCommitQCs using a WAL.
// Each entry is one FullCommitQC covering a range of global block numbers.
// When stateDir is None, all disk I/O is skipped (no-op mode).
// All public methods are safe for concurrent use.
type FullCommitQCPersister struct {
	state utils.Mutex[*fullCommitQCState]
}

// NewFullCommitQCPersister opens (or creates) a WAL in the fullcommitqcs/
// subdir and replays all persisted entries. Loaded QCs are available via
// ConsumeLoaded. When stateDir is None, returns a no-op persister.
func NewFullCommitQCPersister(stateDir utils.Option[string], committee *types.Committee) (*FullCommitQCPersister, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Next returns the next GlobalBlockNumber expected by the persister
// (i.e., the next QC's GlobalRange().First).
func (gp *FullCommitQCPersister) Next() types.GlobalBlockNumber {
	_ = "STUB: not implemented"
	return *new(types.GlobalBlockNumber)
}

// LoadedFirst returns the first global block number of the first loaded QC,
// or committee.FirstBlock() if empty.
func (gp *FullCommitQCPersister) LoadedFirst() types.GlobalBlockNumber {
	_ = "STUB: not implemented"
	return *new(types.GlobalBlockNumber)
}

// ConsumeLoaded returns QCs loaded from the WAL during construction
// and nils the internal slice so the data is not retained.
func (gp *FullCommitQCPersister) ConsumeLoaded() []*types.FullCommitQC {
	_ = "STUB: not implemented"
	return nil
}

// PersistQC appends a FullCommitQC to the WAL. Duplicates are silently ignored.
// Gaps return an error.
func (gp *FullCommitQCPersister) PersistQC(qc *types.FullCommitQC) error {
	_ = "STUB: not implemented"
	return nil
}

// TruncateBefore removes all QC entries whose range is fully before n.
func (gp *FullCommitQCPersister) TruncateBefore(n types.GlobalBlockNumber) error {
	_ = "STUB: not implemented"
	return nil
}

// Close shuts down the WAL.
func (gp *FullCommitQCPersister) Close() error { _ = "STUB: not implemented"; return nil }

func (s *fullCommitQCState) loadAll() ([]*types.FullCommitQC, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
