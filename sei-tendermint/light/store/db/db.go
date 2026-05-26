package db

import (
	"sync"

	dbm "github.com/tendermint/tm-db"

	"github.com/sei-protocol/sei-chain/sei-tendermint/light/store"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

// key prefixes
// NB: Before modifying these, cross-check them with those in
// * internal/store/store.go    [0..4, 13]
// * internal/state/store.go    [5..8, 14]
// * internal/evidence/pool.go  [9..10]
// * light/store/db/db.go       [11..12]
// TODO(sergio): Move all these to their own package.
// TODO: what about these (they already collide):
// * scripts/scmigrate/migrate.go [3]
// * internal/p2p/peermanager.go  [1]
const (
	prefixLightBlock = int64(11)
	prefixSize       = int64(12)
)

type dbs struct {
	db dbm.DB

	mtx  sync.RWMutex
	size uint16
}

// New returns a Store that wraps any DB
// If you want to share one DB across many light clients consider using PrefixDB
func New(db dbm.DB) store.Store { _ = "STUB: not implemented"; return *new(store.Store) }

// retrieve the size of the db

// SaveLightBlock persists LightBlock to the db.
//
// Safe for concurrent use by multiple goroutines.
func (s *dbs) SaveLightBlock(lb *types.LightBlock) error { _ = "STUB: not implemented"; return nil }

// DeleteLightBlockAndValidatorSet deletes the LightBlock from
// the db.
//
// Safe for concurrent use by multiple goroutines.
func (s *dbs) DeleteLightBlock(height int64) error { _ = "STUB: not implemented"; return nil }

// LightBlock retrieves the LightBlock at the given height.
//
// Safe for concurrent use by multiple goroutines.
func (s *dbs) LightBlock(height int64) (*types.LightBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LastLightBlockHeight returns the last LightBlock height stored.
//
// Safe for concurrent use by multiple goroutines.
func (s *dbs) LastLightBlockHeight() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// FirstLightBlockHeight returns the first LightBlock height stored.
//
// Safe for concurrent use by multiple goroutines.
func (s *dbs) FirstLightBlockHeight() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// LightBlockBefore iterates over light blocks until it finds a block before
// the given height. It returns ErrLightBlockNotFound if no such block exists.
//
// Safe for concurrent use by multiple goroutines.
func (s *dbs) LightBlockBefore(height int64) (*types.LightBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Prune prunes header & validator set pairs until there are only size pairs
// left.
//
// Safe for concurrent use by multiple goroutines.
func (s *dbs) Prune(size uint16) error {
	_ = "STUB: not implemented"
	// 1) Check how many we need to prune.
	return nil
}

// nothing to prune

// 2) use an iterator to batch together all the blocks that need to be deleted

// 3) // update size

// 4) write batch deletion to disk

// Size returns the number of header & validator set pairs.
//
// Safe for concurrent use by multiple goroutines.
func (s *dbs) Size() uint16 { _ = "STUB: not implemented"; return 0 }

func (s *dbs) batchDelete(batch dbm.Batch, numToPrune uint16) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *dbs) sizeKey() []byte { _ = "STUB: not implemented"; return nil }

func (s *dbs) lbKey(height int64) []byte { _ = "STUB: not implemented"; return nil }

func (s *dbs) decodeLbKey(key []byte) (height int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func marshalSize(size uint16) []byte { _ = "STUB: not implemented"; return nil }

func unmarshalSize(bz []byte) uint16 { _ = "STUB: not implemented"; return 0 }
