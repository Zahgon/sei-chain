package clist

/*

The purpose of CList is to provide a goroutine-safe linked-list.
This list can be traversed concurrently by any number of goroutines.
However, removed CElements cannot be added back.
NOTE: Not all methods of container/list are (yet) implemented.
NOTE: Removed elements need to DetachPrev or DetachNext consistently
to ensure garbage collection of removed elements.

*/

import (
	"context"
	"errors"
	"sync"
)

/*
CElement is an element of a linked-list
Traversal from a CElement is goroutine-safe.

We can't avoid using WaitGroups or for-loops given the documentation
spec without re-implementing the primitives that already exist in
golang/sync. Notice that WaitGroup allows many go-routines to be
simultaneously released, which is what we want. Mutex doesn't do
this. RWMutex does this, but it's clumsy to use in the way that a
WaitGroup would be used -- and we'd end up having two RWMutex's for
prev/next each, which is doubly confusing.

sync.Cond would be sort-of useful, but we don't need a write-lock in
the for-loop. Use sync.Cond when you need serial access to the
"condition". In our case our condition is if `next != nil || removed`,
and there's no reason to serialize that condition for goroutines
waiting on NextWait() (since it's just a read operation).
*/
type CElement[T any] struct {
	mtx        sync.RWMutex
	prev       *CElement[T]
	next       *CElement[T]
	nextWaitCh chan struct{}
	removed    bool

	value T // immutable
}

var ErrRemoved = errors.New("element was removed")

// Blocking implementation of Next().
// May return ErrRemoved iff CElement was tail and got removed.
func (e *CElement[T]) NextWait(ctx context.Context) (*CElement[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// e.next doesn't necessarily exist here.
// That's why we need to continue a for-loop.

// Nonblocking, may return nil if at the end.
func (e *CElement[T]) Next() *CElement[T] { _ = "STUB: not implemented"; return nil }

// Nonblocking, may return nil if at the end.
func (e *CElement[T]) Prev() *CElement[T] { _ = "STUB: not implemented"; return nil }

func (e *CElement[T]) Removed() bool { _ = "STUB: not implemented"; return false }

func (e *CElement[T]) Value() T { _ = "STUB: not implemented"; return *new(T) }

func (e *CElement[T]) detachNext() { _ = "STUB: not implemented"; return }

func (e *CElement[T]) DetachPrev() { _ = "STUB: not implemented"; return }

// NOTE: This function needs to be safe for
// concurrent goroutines waiting on nextWg.
func (e *CElement[T]) setNext(newNext *CElement[T]) { _ = "STUB: not implemented"; return }

// See https://golang.org/pkg/sync/:
//
// If a WaitGroup is reused to wait for several independent sets of
// events, new Add calls must happen after all previous Wait calls have
// returned.

// NOTE: This function needs to be safe for
// concurrent goroutines waiting on prevWg
func (e *CElement[T]) setPrev(newPrev *CElement[T]) { _ = "STUB: not implemented"; return }

func (e *CElement[T]) setRemoved() { _ = "STUB: not implemented"; return }

// This wakes up anyone waiting.

//--------------------------------------------------------------------------------

// CList represents a linked list.
// The zero value for CList is an empty list ready to use.
// Operations are goroutine-safe.
type CList[T any] struct {
	mtx    sync.RWMutex
	waitCh chan struct{}
	head   *CElement[T] // first element
	tail   *CElement[T] // last element
	len    int          // list length
}

func New[T any]() *CList[T] { _ = "STUB: not implemented"; return nil }

func (l *CList[T]) Len() int { _ = "STUB: not implemented"; return 0 }

func (l *CList[T]) Front() *CElement[T] { _ = "STUB: not implemented"; return nil }

func (l *CList[T]) WaitFront(ctx context.Context) (*CElement[T], error) {
	_ = "STUB: not implemented"
	// Loop until the head is non-nil else wait and try again
	return nil, nil
}

// NOTE: If you think l.head exists here, think harder.

func (l *CList[T]) Back() *CElement[T] { _ = "STUB: not implemented"; return nil }

// Panics if list grows beyond its max length.
func (l *CList[T]) PushBack(v T) *CElement[T] {
	_ = "STUB: not implemented"

	// Construct a new element
	return nil
}

// Release waiters on FrontWait/BackWait maybe

// Modify the tail

// We must init e first.
// This will make e accessible.
// Update the list.

// CONTRACT: Caller must call e.DetachPrev() and/or e.DetachNext() to avoid memory leaks.
// NOTE: As per the contract of CList, removed elements cannot be added back.
func (l *CList[T]) Remove(e *CElement[T]) T { _ = "STUB: not implemented"; return *new(T) }

// If we're removing the only item, make CList FrontWait/BackWait wait.

// Update l.len

// Connect next/prev and set head/tail

// Set .Done() on e, otherwise waiters will wait forever.
