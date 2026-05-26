package utils

func FilterUInt64Slice(slice []uint64, item uint64) []uint64 { _ = "STUB: not implemented"; return nil }

func Map[I any, O any](input []I, lambda func(i I) O) []O { _ = "STUB: not implemented"; return nil }

func SliceCopy[T any](slice []T) []T { _ = "STUB: not implemented"; return nil }

func Reduce[I, O any](input []I, reducer func(I, O) O, initial O) O {
	_ = "STUB: not implemented"
	return *new(O)
}

func Filter[T any](slice []T, lambda func(t T) bool) []T { _ = "STUB: not implemented"; return nil }

func Copy[T any](slice []T) []T { _ = "STUB: not implemented"; return nil }
