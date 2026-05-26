package api

/*
#include "bindings.h"
*/
import "C"

// makeView creates a view into the given byte slice what allows Rust code to read it.
// The byte slice is managed by Go and will be garbage collected. Use runtime.KeepAlive
// to ensure the byte slice lives long enough.
func makeView(s []byte) C.ByteSliceView { _ = "STUB: not implemented"; return *new(C.ByteSliceView) }

// In Go, accessing the 0-th element of an empty array triggers a panic. That is why in the case
// of an empty `[]byte` we can't get the internal heap pointer to the underlying array as we do
// below with `&data[0]`. https://play.golang.org/p/xvDY3g9OqUk

// Creates a C.UnmanagedVector, which cannot be done in test files directly
func constructUnmanagedVector(is_none cbool, ptr cu8_ptr, len cusize, cap cusize) C.UnmanagedVector {
	_ = "STUB: not implemented"
	return *new(C.UnmanagedVector)
}

// uninitializedUnmanagedVector returns an invalid C.UnmanagedVector
// instance. Only use then after someone wrote an instance to it.
func uninitializedUnmanagedVector() C.UnmanagedVector {
	_ = "STUB: not implemented"
	return *new(C.UnmanagedVector)
}

func newUnmanagedVector(data []byte) C.UnmanagedVector {
	_ = "STUB: not implemented"
	return *new(C.UnmanagedVector)
}

// in Go, accessing the 0-th element of an empty array triggers a panic. That is why in the case
// of an empty `[]byte` we can't get the internal heap pointer to the underlying array as we do
// below with `&data[0]`.
// https://play.golang.org/p/xvDY3g9OqUk

// This will allocate a proper vector with content and return a description of it

func copyAndDestroyUnmanagedVector(v C.UnmanagedVector) []byte {
	_ = "STUB: not implemented"
	return nil
}

// There is no allocation we can copy

// C.GoBytes create a copy (https://stackoverflow.com/a/40950744/2013738)

// copyU8Slice copies the contents of an Option<&[u8]> that was allocated on the Rust side.
// Returns nil if and only if the source is None.
func copyU8Slice(view C.U8SliceView) []byte { _ = "STUB: not implemented"; return nil }

// In this case, we don't want to look into the ptr

// C.GoBytes create a copy (https://stackoverflow.com/a/40950744/2013738)
