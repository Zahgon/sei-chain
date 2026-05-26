package utils

// Option type inspired https://pkg.go.dev/github.com/samber/mo.
type Option[T any] struct {
	ReadOnly
	isPresent bool
	value     T
}

// Some creates an Option with a value.
func Some[T any](value T) Option[T] { _ = "STUB: not implemented"; return nil }

// None creates an Option without a value.
func None[T any]() (zero Option[T]) {
	_ = "STUB: not implemented"

	// Get unpacks the value from the Option, returning true if it was present.
	return nil
}

func (o Option[T]) Get() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

// IsPresent checks if the Option contains a value.
func (o Option[T]) IsPresent() bool {
	_ = "STUB: not implemented"

	// Or returns the value if present, otherwise returns the default value.
	return false
}

func (o *Option[T]) Or(def T) T { _ = "STUB: not implemented"; return *new(T) }

func (o Option[T]) OrPanic(msg string) T { _ = "STUB: not implemented"; return *new(T) }

// MapOpt applies a function to the value if present, returning a new Option.
func MapOpt[T, R any](o Option[T], f func(T) R) Option[R] { _ = "STUB: not implemented"; return nil }

// MarshalJSON implements the json.Marshaler interface.
// Note that it is defined on value, not pointer, because
// json.Marshal cannot call pointer methods on fields
// (i.e. it is broken by design).
func (o Option[T]) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implements the json.Unmarshaler interface.
func (o *Option[T]) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
