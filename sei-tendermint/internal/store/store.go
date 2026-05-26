package store

import (
	"github.com/gogo/protobuf/proto"
	dbm "github.com/tendermint/tm-db"

	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

/*
BlockStore is a simple low level store for blocks.

There are three types of information stored:
  - BlockMeta:   Meta information about each block
  - Block part:  Parts of each block, aggregated w/ PartSet
  - Commit:      The commit part of each block, for gossiping precommit votes

Currently the precommit signatures are duplicated in the Block parts as
well as the Commit.  In the future this may change, perhaps by moving
the Commit data outside the Block. (TODO)

The store can be assumed to contain all contiguous blocks between base and height (inclusive).

// NOTE: BlockStore methods will panic if they encounter errors
// deserializing loaded data, indicating probable corruption on disk.
*/
type BlockStore struct {
	db dbm.DB
}

// NewBlockStore returns a new BlockStore with the given DB,
// initialized to the last height that was committed to the DB.
func NewBlockStore(db dbm.DB) *BlockStore { _ = "STUB: not implemented"; return nil }

// Base returns the first known contiguous block height, or 0 for empty block stores.
func (bs *BlockStore) Base() int64 { _ = "STUB: not implemented"; return 0 }

// Height returns the last known contiguous block height, or 0 for empty block stores.
func (bs *BlockStore) Height() int64 { _ = "STUB: not implemented"; return 0 }

// Size returns the number of blocks in the block store.
func (bs *BlockStore) Size() int64 { _ = "STUB: not implemented"; return 0 }

// LoadBase atomically loads the base block meta, or returns nil if no base is found.
func (bs *BlockStore) LoadBaseMeta() *types.BlockMeta { _ = "STUB: not implemented"; return nil }

// LoadBlock returns the block with the given height.
// If no block is found for that height, it returns nil.
func (bs *BlockStore) LoadBlock(height int64) *types.Block { _ = "STUB: not implemented"; return nil }

// If the part is missing (e.g. since it has been deleted after we
// loaded the block meta) we consider the whole block to be missing.

// NOTE: The existence of meta should imply the existence of the
// block. So, make sure meta is only saved after blocks are saved.

// LoadBlockByHash returns the block with the given hash.
// If no block is found for that hash, it returns nil.
// Panics if it fails to parse height associated with the given hash.
func (bs *BlockStore) LoadBlockByHash(hash []byte) *types.Block {
	_ = "STUB: not implemented"
	return nil
}

// LoadBlockMetaByHash returns the blockmeta who's header corresponds to the given
// hash. If none is found, returns nil.
func (bs *BlockStore) LoadBlockMetaByHash(hash []byte) *types.BlockMeta {
	_ = "STUB: not implemented"
	return nil
}

// LoadBlockPart returns the Part at the given index
// from the block at the given height.
// If no part is found for the given height and index, it returns nil.
func (bs *BlockStore) LoadBlockPart(height int64, index int) *types.Part {
	_ = "STUB: not implemented"
	return nil
}

// LoadBlockMeta returns the BlockMeta for the given height.
// If no block is found for the given height, it returns nil.
func (bs *BlockStore) LoadBlockMeta(height int64) *types.BlockMeta {
	_ = "STUB: not implemented"
	return nil
}

// LoadBlockCommit returns the Commit for the given height.
// This commit consists of the +2/3 and other Precommit-votes for block at `height`,
// and it comes from the block.LastCommit for `height+1`.
// If no commit is found for the given height, it returns nil.
func (bs *BlockStore) LoadBlockCommit(height int64) *types.Commit {
	_ = "STUB: not implemented"
	return nil
}

// LoadSeenCommit returns the last locally seen Commit before being
// cannonicalized. This is useful when we've seen a commit, but there
// has not yet been a new block at `height + 1` that includes this
// commit in its block.LastCommit.
func (bs *BlockStore) LoadSeenCommit() *types.Commit { _ = "STUB: not implemented"; return nil }

// PruneBlocks removes block up to (but not including) a height. It returns the number of blocks pruned.
func (bs *BlockStore) PruneBlocks(height int64) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// when removing the block meta, use the hash to remove the hash key at the same time

// unmarshal block meta

// delete the hash key corresponding to the block meta's hash

// remove block meta first as this is used to indicate whether the block exists.
// For this reason, we also use ony block meta as a measure of the amount of blocks pruned

// pruneRange is a generic function for deleting a range of values based on the lowest
// height up to but excluding retainHeight. For each key/value pair, an optional hook can be
// executed before the deletion itself is made. pruneRange will use batch delete to delete
// keys in batches of at most 1000 keys.
func (bs *BlockStore) pruneRange(
	start []byte,
	end []byte,
	preDeletionHook func(key, value []byte, batch dbm.Batch) error,
) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// loop until we have finished iterating over all the keys by writing, opening a new batch
// and incrementing through the next range of keys.

// once we looped over all keys we do a final flush to disk

// batchDelete runs an iterator over a set of keys, first preforming a pre deletion hook before adding it to the batch.
// The function ends when either 1000 keys have been added to the batch or the iterator has reached the end.
func (bs *BlockStore) batchDelete(
	batch dbm.Batch,
	start, end []byte,
	preDeletionHook func(key, value []byte, batch dbm.Batch) error,
) (uint64, []byte, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

// SaveBlock persists the given block, blockParts, and seenCommit to the underlying db.
// blockParts: Must be parts of the block
// seenCommit: The +2/3 precommits that were seen which committed at height.
//
//	If all the nodes restart after committing a block,
//	we need this to reload the precommits to catch-up nodes to the
//	most recent height.  Otherwise they'd stall at H-1.
func (bs *BlockStore) SaveBlock(block *types.Block, blockParts *types.PartSet, seenCommit *types.Commit) {
	_ = "STUB: not implemented"
	return
}

func (bs *BlockStore) saveBlockToBatch(batch dbm.Batch, block *types.Block, blockParts *types.PartSet, seenCommit *types.Commit) error {
	_ = "STUB: not implemented"
	return nil
}

// Save block parts. This must be done before the block meta, since callers
// typically load the block meta first as an indication that the block exists
// and then go on to load block parts - we must make sure the block is
// complete as soon as the block meta is written.

// Save seen commit (seen +2/3 precommits for block)

func (bs *BlockStore) saveBlockPart(height int64, index int, part *types.Part, batch dbm.Batch) {
	_ = "STUB: not implemented"
	return
}

// SaveSeenCommit saves a seen commit, used by e.g. the state sync reactor when bootstrapping node.
func (bs *BlockStore) SaveSeenCommit(height int64, seenCommit *types.Commit) error {
	_ = "STUB: not implemented"
	return nil
}

func (bs *BlockStore) SaveSignedHeader(sh *types.SignedHeader, blockID types.BlockID) error {
	_ = "STUB: not implemented"
	// first check that the block store doesn't already have the block
	return nil
}

// FIXME: saving signed headers although necessary for proving evidence,
// doesn't have complete parity with block meta's thus block size and num
// txs are filled with negative numbers. We should aim to find a solution to
// this.

func (bs *BlockStore) Close() error { _ = "STUB: not implemented"; return nil }

//---------------------------------- KEY ENCODING -----------------------------------------

// key prefixes
// NB: Before modifying these, cross-check them with those in
// * internal/store/store.go    [0..4, 13]
// * internal/state/store.go    [5..8, 14]
// * internal/evidence/pool.go  [9..10]
// * light/store/db/db.go       [11..12]
// TODO(thane): Move all these to their own package.
// TODO: what about these (they already collide):
// * scripts/scmigrate/migrate.go [3] --> Looks OK, as it is also called "SeenCommit"
// * internal/p2p/peermanager.go  [1]
const (
	// prefixes are unique across all tm db's
	prefixBlockMeta   = int64(0)
	prefixBlockPart   = int64(1)
	prefixBlockCommit = int64(2)
	prefixSeenCommit  = int64(3)
	prefixBlockHash   = int64(4)
	prefixExtCommit   = int64(13)
)

func blockMetaKey(height int64) []byte { _ = "STUB: not implemented"; return nil }

func decodeBlockMetaKey(key []byte) (height int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func blockPartKey(height int64, partIndex int) []byte { _ = "STUB: not implemented"; return nil }

func blockCommitKey(height int64) []byte { _ = "STUB: not implemented"; return nil }

func seenCommitKey() []byte { _ = "STUB: not implemented"; return nil }

//lint:ignore U1000 Method retained for completeness
func extCommitKey(height int64) []byte { _ = "STUB: not implemented"; return nil }

func blockHashKey(hash []byte) []byte { _ = "STUB: not implemented"; return nil }

//-----------------------------------------------------------------------------

// mustEncode proto encodes a proto.message and panics if fails
func mustEncode(pb proto.Message) []byte { _ = "STUB: not implemented"; return nil }

//-----------------------------------------------------------------------------

// DeleteLatestBlock removes the block pointed to by height,
// lowering height by one.
func (bs *BlockStore) DeleteLatestBlock() error { _ = "STUB: not implemented"; return nil }

// delete what we can, skipping what's already missing, to ensure partial
// blocks get deleted fully.

// delete last, so as to not leave keys built on meta.BlockID dangling
