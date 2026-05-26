package types

import (
	"iter"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// SortedSet is an immutable set of elements.
// It supports iterating over elements in order,
// and O(1) access to elements by index.
type SortedSet[T Compare[T]] struct {
	sorted []T
	hashed map[T]struct{}
}

// Compare is an interface that defines a method for comparing two elements.
type Compare[T any] interface {
	comparable
	Compare(b T) int
}

// NewSortedSet creates a new SortedSet from a slice of elements.
func NewSortedSet[T Compare[T]](vs []T) SortedSet[T] { _ = "STUB: not implemented"; return nil }

// All returns an iterator over all elements in order.
func (s SortedSet[T]) All() iter.Seq[T] { _ = "STUB: not implemented"; return nil }

// Len returns the number of elements in the set.
func (s SortedSet[T]) Len() int { _ = "STUB: not implemented"; return 0 }

// Has checks if the set contains the given value.
func (s SortedSet[T]) Has(val T) bool { _ = "STUB: not implemented"; return false }

// At returns the element at the given index.
func (s SortedSet[T]) At(i int) T {
	_ = "STUB: not implemented"

	// Committee represents the consensus committee.
	return *new(T)
}

type Committee struct {
	replicas SortedSet[PublicKey]
	// Number of the first block of the chain.
	// TODO: firstBlock is not really a part of the committee,
	// but it does belong to a chain spec (or epoch spec/genesis/etc.),
	// which should be passed around to verify autobahn messages.
	// Once we introduce the chain spec it should wrap Committee and firstBlock.
	firstBlock GlobalBlockNumber
	// timestamp at genesis. All blocks need to have a timestamp later than genesis.
	genesisTimestamp time.Time
}

// Lanes is the list of nodes which are eligible to produce blocks.
func (c *Committee) Lanes() SortedSet[LaneID] {
	_ = "STUB: not implemented"

	// Replicas is the list of nodes which are eligible to participate in the consensus.
	return nil
}

func (c *Committee) Replicas() SortedSet[PublicKey] {
	_ = "STUB: not implemented"

	// FirstBlock is the index of the first global block finalized by this committee.
	return nil
}

func (c *Committee) FirstBlock() GlobalBlockNumber {
	_ = "STUB: not implemented"
	return *

	// GenesisTimestamp is the timestamp at genesis.
	new(GlobalBlockNumber)
}

func (c *Committee) GenesisTimestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (c *Committee) EvmShard(addr common.Address) PublicKey {
	_ = "STUB: not implemented"
	return *new(PublicKey)
}

// Leader for the consensus round with the given index.
func (c *Committee) Leader(view View) PublicKey { _ = "STUB: not implemented"; return *new(PublicKey) }

// Faulty is the number of faulty replicas that consensus can tolerate.
func (c *Committee) Faulty() int {
	_ = "STUB: not implemented"
	// 3f < N
	return 0
}

// CommitQuorum is the size of the quorum required for CommitQC.
func (c *Committee) CommitQuorum() int { _ = "STUB: not implemented"; return 0 }

// AppQuorum is the size of the quorum required for AppQC.
func (c *Committee) AppQuorum() int {
	_ = "STUB: not implemented"
	// This needs to be in range (c.Faulty(), c.CommitQuorum()]
	return 0
}

// PrepareQuorum is the size of the quorum required for PrepareQC.
func (c *Committee) PrepareQuorum() int { _ = "STUB: not implemented"; return 0 }

// TimeoutQuorum is the size of the quorum required for TimeoutQC.
func (c *Committee) TimeoutQuorum() int { _ = "STUB: not implemented"; return 0 }

// LaneQuorum is the size of the quorum required for LaneQC.
func (c *Committee) LaneQuorum() int { _ = "STUB: not implemented"; return 0 }

// NewRoundRobinElection creates a Committee with round robin election starting at firstBlock.
func NewRoundRobinElection(replicas []PublicKey, firstBlock GlobalBlockNumber, genesisTimestamp time.Time) (*Committee, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
