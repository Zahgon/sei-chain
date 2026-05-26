package p2p

import (
	"context"
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

type ord[T any] interface {
	Less(T) bool
}

type withIdx[T any] struct {
	v      T
	minIdx int // index in byMin
	maxIdx int // index in byMax
}

func newWithIdx[T any](v T) *withIdx[T] { _ = "STUB: not implemented"; return nil }

// Heap returning minimal elements.
type byMin[T ord[T]] struct{ a []*withIdx[T] }

func newByMin[T ord[T]](capacity int) byMin[T] { _ = "STUB: not implemented"; return nil }
func (x *byMin[T]) Less(i, j int) bool         { _ = "STUB: not implemented"; return false }
func (x *byMin[T]) Len() int                   { _ = "STUB: not implemented"; return 0 }
func (x *byMin[T]) Swap(i, j int)              { _ = "STUB: not implemented"; return }

func (x *byMin[T]) Push(v any) { _ = "STUB: not implemented"; return }

func (x *byMin[T]) Pop() any { _ = "STUB: not implemented"; return *new(any) }

// Heap returning maximal elements.
type byMax[T ord[T]] struct{ a []*withIdx[T] }

func newByMax[T ord[T]](capacity int) byMax[T] { _ = "STUB: not implemented"; return nil }
func (x *byMax[T]) Less(i, j int) bool         { _ = "STUB: not implemented"; return false }
func (x *byMax[T]) Len() int                   { _ = "STUB: not implemented"; return 0 }
func (x *byMax[T]) Swap(i, j int)              { _ = "STUB: not implemented"; return }

func (x *byMax[T]) Push(v any) { _ = "STUB: not implemented"; return }

func (x *byMax[T]) Pop() any { _ = "STUB: not implemented"; return *new(any) }

// pqEnvelope defines a wrapper around an Envelope with priority to be inserted
// into a priority Queue used for Envelope scheduling.
type pqEnvelope[M any] struct {
	msg       M
	priority  int
	size      int
	timestamp time.Time
}

// true <=> a has higher priority than b
func (a *pqEnvelope[M]) Less(b *pqEnvelope[M]) bool {
	_ = "STUB: not implemented"
	// higher base priority wins
	return false
}

// newer timestamp wins

// larger first

type inner[M any] struct {
	capacity int
	byMin    byMin[*pqEnvelope[M]]
	byMax    byMax[*pqEnvelope[M]]
}

func newInner[M any](capacity int) *inner[M] { _ = "STUB: not implemented"; return nil }

// We prune the maximal elements whenever capacity is exceeded.
// Therefore to avoid reallocation we need the heaps to have capacity+1.

func (i *inner[M]) Len() int { _ = "STUB: not implemented"; return 0 }

func (i *inner[M]) Push(e *pqEnvelope[M]) utils.Option[M] { _ = "STUB: not implemented"; return nil }

func (i *inner[M]) Pop() *pqEnvelope[M] { _ = "STUB: not implemented"; return nil }

type Queue[M any] struct{ inner utils.Watch[*inner[M]] }

func NewQueue[M any](size int) *Queue[M] {
	_ = "STUB: not implemented"

	// prevent caller from shooting self in the foot.
	return nil
}

func (q *Queue[M]) Len() int { _ = "STUB: not implemented"; return 0 }

// Non-blocking send.
// Returns the pruned message if any.
func (q *Queue[M]) Send(msg M, size int, priority int) utils.Option[M] {
	_ = "STUB: not implemented"
	// We construct the pqEnvelope without holding the lock to avoid contention.
	return nil
}

// Blocking recv.
func (q *Queue[M]) Recv(ctx context.Context) (M, error) {
	_ = "STUB: not implemented"
	return *new(M), nil
}
