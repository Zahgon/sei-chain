package kv

// This code was copied from golang.org/pkg/container/list, but specially adapted
// for use with kv.Pair to avoid the type assertion CPU expense of using Value with
// an interface, per https://github.com/cosmos/cosmos-sdk/issues/8810
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Element is an element of a linked list.
type Element struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	next, prev *Element

	// The list to which this element belongs.
	list *List

	// The value stored with this element.
	Value *Pair
}

// Next returns the next list element or nil.
func (e *Element) Next() *Element { _ = "STUB: not implemented"; return nil }

// Prev returns the previous list element or nil.
func (e *Element) Prev() *Element { _ = "STUB: not implemented"; return nil }

// List represents a doubly linked list.
// The zero value for List is an empty list ready to use.
type List struct {
	root Element // sentinel list element, only &root, root.prev, and root.next are used
	len  int     // current list length excluding (this) sentinel element
}

// Init initializes or clears list l.
func (l *List) Init() *List { _ = "STUB: not implemented"; return nil }

// NewList returns an initialized list.
func NewList() *List { _ = "STUB: not implemented"; return nil }

// Len returns the number of elements of list l.
// The complexity is O(1).
func (l *List) Len() int {
	_ = "STUB: not implemented"

	// Front returns the first element of list l or nil if the list is empty.
	return 0
}

func (l *List) Front() *Element { _ = "STUB: not implemented"; return nil }

// Back returns the last element of list l or nil if the list is empty.
func (l *List) Back() *Element { _ = "STUB: not implemented"; return nil }

// lazyInit lazily initializes a zero List value.
func (l *List) lazyInit() { _ = "STUB: not implemented"; return }

// insert inserts e after at, increments l.len, and returns e.
func (l *List) insert(e, at *Element) *Element { _ = "STUB: not implemented"; return nil }

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *List) insertValue(v *Pair, at *Element) *Element { _ = "STUB: not implemented"; return nil }

// remove removes e from its list, decrements l.len, and returns e.
func (l *List) remove(e *Element) *Element { _ = "STUB: not implemented"; return nil }

// avoid memory leaks
// avoid memory leaks

// move moves e to next to at and returns e.
// nolint: unparam
func (l *List) move(e, at *Element) *Element { _ = "STUB: not implemented"; return nil }

// Remove removes e from l if e is an element of list l.
// It returns the element value e.Value.
// The element must not be nil.
func (l *List) Remove(e *Element) *Pair {
	_ = "STUB: not implemented"

	// if e.list == l, l must have been initialized when e was inserted
	// in l or l == nil (e is a zero Element) and l.remove will crash
	return nil
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *List) PushFront(v *Pair) *Element { _ = "STUB: not implemented"; return nil }

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *List) PushBack(v *Pair) *Element { _ = "STUB: not implemented"; return nil }

// InsertBefore inserts a new element e with value v immediately before mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *List) InsertBefore(v *Pair, mark *Element) *Element { _ = "STUB: not implemented"; return nil }

// see comment in List.Remove about initialization of l

// InsertAfter inserts a new element e with value v immediately after mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *List) InsertAfter(v *Pair, mark *Element) *Element { _ = "STUB: not implemented"; return nil }

// see comment in List.Remove about initialization of l

// MoveToFront moves element e to the front of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *List) MoveToFront(e *Element) { _ = "STUB: not implemented"; return }

// see comment in List.Remove about initialization of l

// MoveToBack moves element e to the back of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *List) MoveToBack(e *Element) { _ = "STUB: not implemented"; return }

// see comment in List.Remove about initialization of l

// MoveBefore moves element e to its new position before mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *List) MoveBefore(e, mark *Element) { _ = "STUB: not implemented"; return }

// MoveAfter moves element e to its new position after mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *List) MoveAfter(e, mark *Element) { _ = "STUB: not implemented"; return }

// PushBackList inserts a copy of another list at the back of list l.
// The lists l and other may be the same. They must not be nil.
func (l *List) PushBackList(other *List) { _ = "STUB: not implemented"; return }

// PushFrontList inserts a copy of another list at the front of list l.
// The lists l and other may be the same. They must not be nil.
func (l *List) PushFrontList(other *List) { _ = "STUB: not implemented"; return }
