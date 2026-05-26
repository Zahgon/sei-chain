package multiversion

import (
	"sync"

	"github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/occ"
	db "github.com/tendermint/tm-db"
)

type MultiVersionStore interface {
	GetLatest(key []byte) (value MultiVersionValueItem)
	GetLatestBeforeIndex(index int, key []byte) (value MultiVersionValueItem)
	Has(index int, key []byte) bool
	WriteLatestToStore()
	SetWriteset(index int, incarnation int, writeset WriteSet)
	InvalidateWriteset(index int, incarnation int)
	SetEstimatedWriteset(index int, incarnation int, writeset WriteSet)
	GetAllWritesetKeys() map[int][]string
	CollectIteratorItems(index int) *db.MemDB
	SetReadset(index int, readset ReadSet)
	GetReadset(index int) ReadSet
	ClearReadset(index int)
	VersionedIndexedStore(index int, incarnation int, abortChannel chan occ.Abort) *VersionIndexedStore
	SetIterateset(index int, iterateset Iterateset)
	GetIterateset(index int) Iterateset
	ClearIterateset(index int)
	ValidateTransactionState(index int) (bool, []int)
	ValidateTransactionStateWithKeys(index int) (bool, []int, []string)
}

type WriteSet map[string][]byte
type ReadSet map[string][][]byte
type Iterateset []*iterationTracker

var _ MultiVersionStore = (*Store)(nil)

type Store struct {
	// map that stores the key string -> MultiVersionValue mapping for accessing from a given key
	multiVersionMap *sync.Map
	// TODO: do we need to support iterators as well similar to how cachekv does it - yes

	txWritesetKeys *sync.Map // map of tx index -> writeset keys []string
	txReadSets     *sync.Map // map of tx index -> readset ReadSet
	txIterateSets  *sync.Map // map of tx index -> iterateset Iterateset

	parentStore types.KVStore
}

func NewMultiVersionStore(parentStore types.KVStore) *Store { _ = "STUB: not implemented"; return nil }

// VersionedIndexedStore creates a new versioned index store for a given incarnation and transaction index
func (s *Store) VersionedIndexedStore(index int, incarnation int, abortChannel chan occ.Abort) *VersionIndexedStore {
	_ = "STUB: not implemented"
	return nil
}

// GetLatest implements MultiVersionStore.
func (s *Store) GetLatest(key []byte) (value MultiVersionValueItem) {
	_ = "STUB: not implemented"
	return *new(MultiVersionValueItem)
}

// if the key doesn't exist in the overall map, return nil

// this is possible IF there is are writeset that are then removed for that key

// GetLatestBeforeIndex implements MultiVersionStore.
func (s *Store) GetLatestBeforeIndex(index int, key []byte) (value MultiVersionValueItem) {
	_ = "STUB: not implemented"
	return *new(MultiVersionValueItem)
}

// if the key doesn't exist in the overall map, return nil

// otherwise, we may have found a value for that key, but its not written before the index passed in

// found a value prior to the passed in index, return that value (could be estimate OR deleted, but it is a definitive value)

// Has implements MultiVersionStore. It checks if the key exists in the multiversion store at or before the specified index.
func (s *Store) Has(index int, key []byte) bool { _ = "STUB: not implemented"; return false }

// if the key doesn't exist in the overall map, return nil

// this is okay because the caller of this will THEN need to access the parent store to verify that the key doesnt exist there

func (s *Store) removeOldWriteset(index int, newWriteSet WriteSet) {
	_ = "STUB: not implemented"
	return
}

// if non-nil writeset passed in, we can use that to optimize removals

// if there is already a writeset existing, we should remove that fully

// we need to delete all of the keys in the writeset from the multiversion store

// small optimization to check if the new writeset is going to write this key, if so, we can leave it behind

// we don't need to remove this key because it will be overwritten anyways - saves the operation of removing + rebalancing underlying btree

// remove from the appropriate item if present in multiVersionMap

// if the key doesn't exist in the overall map, return nil

// SetWriteset sets a writeset for a transaction index, and also writes all of the multiversion items in the writeset to the multiversion store.
// TODO: returns a list of NEW keys added
func (s *Store) SetWriteset(index int, incarnation int, writeset WriteSet) {
	_ = "STUB: not implemented"
	// TODO: add telemetry spans
	// remove old writeset if it exists
	return
}

// init if necessary

// delete if nil value
// TODO: sync map

// TODO: if we're sorting here anyways, maybe we just put it into a btree instead of a slice

// InvalidateWriteset iterates over the keys for the given index and incarnation writeset and replaces with ESTIMATEs
func (s *Store) InvalidateWriteset(index int, incarnation int) { _ = "STUB: not implemented"; return }

// invalidate all of the writeset items - is this suboptimal? - we could potentially do concurrently if slow because locking is on an item specific level

// we leave the writeset in place because we'll need it for key removal later if/when we replace with a new writeset

// SetEstimatedWriteset is used to directly write estimates instead of writing a writeset and later invalidating
func (s *Store) SetEstimatedWriteset(index int, incarnation int, writeset WriteSet) {
	_ = "STUB: not implemented"
	// remove old writeset if it exists
	return
}

// still need to save the writeset so we can remove the elements later:

// init if necessary

// GetAllWritesetKeys implements MultiVersionStore.
func (s *Store) GetAllWritesetKeys() map[int][]string { _ = "STUB: not implemented"; return nil }

// TODO: is this safe?

func (s *Store) SetReadset(index int, readset ReadSet) { _ = "STUB: not implemented"; return }

func (s *Store) GetReadset(index int) ReadSet { _ = "STUB: not implemented"; return *new(ReadSet) }

func (s *Store) SetIterateset(index int, iterateset Iterateset) { _ = "STUB: not implemented"; return }

func (s *Store) GetIterateset(index int) Iterateset {
	_ = "STUB: not implemented"
	return *new(Iterateset)
}

func (s *Store) ClearReadset(index int) { _ = "STUB: not implemented"; return }

func (s *Store) ClearIterateset(index int) { _ = "STUB: not implemented"; return }

// CollectIteratorItems implements MultiVersionStore. It will return a memDB containing all of the keys present in the multiversion store within the iteration range prior to (exclusive of) the index.
func (s *Store) CollectIteratorItems(index int) *db.MemDB { _ = "STUB: not implemented"; return nil }

// get all writeset keys prior to index

// TODO: do we want to exclude keys out of the range or just let the iterator handle it?

// TODO: inefficient because (logn) for each key + rebalancing? maybe theres a better way to add to a tree to reduce rebalancing overhead

func (s *Store) validateIterator(index int, tracker iterationTracker) bool {
	_ = "STUB: not implemented"
	// collect items from multiversion store
	return false
}

// add the iterationtracker writeset keys to the sorted items

// listen for abort while iterating

// create a new MVSMergeiterator

// if we have no more expected keys, then the iterator is invalid

// TODO: is this ok to not delete the key since we shouldnt have duplicate keys?

// if key isn't found

// remove from expected keys

// delete(expectedKeys, string(key))

// if our iterator key was the early stop, then we can break

// return whether we found the exact number of expected keys

// if we get an abort, then we know that the iterator is invalid

func (s *Store) checkIteratorAtIndex(index int) bool { _ = "STUB: not implemented"; return false }

// TODO: if the value of the key is nil maybe we need to exclude it? - actually it should

func (s *Store) checkReadsetAtIndex(index int) (bool, []int, []string) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

// iterate over readset and check if the value is the same as the latest value relateive to txIndex in the multiversion store

// get the latest value from the multiversion store

// this is possible if we previously read a value from a transaction write that was later reverted, so this time we read from parent store

// if estimate, mark as conflict index - but don't invalidate

// conflict
// TODO: would we want to return early?

// TODO: do we want to return bool + []int where bool indicates whether it was valid and then []int indicates only ones for which we need to wait due to estimates? - yes i think so?
func (s *Store) ValidateTransactionState(index int) (bool, []int) {
	_ = "STUB: not implemented"
	// defer telemetry.MeasureSince(time.Now(), "store", "mvs", "validate")
	return false, nil
}

// TODO: can we parallelize for all iterators?

func (s *Store) ValidateTransactionStateWithKeys(index int) (bool, []int, []string) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

func (s *Store) WriteLatestToStore() {
	_ = "STUB: not implemented"
	// sort the keys
	return
}

// this means that at some point, there was an estimate, but we have since removed it so there isn't anything writeable at the key, so we can skip

// we shouldn't have any ESTIMATE values when performing the write, because we read the latest non-estimate values only

// if the value is deleted, then delete it from the parent store

// We use []byte(key) instead of conv.UnsafeStrToBytes because we cannot
// be sure if the underlying store might do a save with the byteslice or
// not. Once we get confirmation that .Delete is guaranteed not to
// save the byteslice, then we can assume only a read-only copy is sufficient.
