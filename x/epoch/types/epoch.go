package types

// NewEpoch creates a new Epoch instance
func NewEpoch() Epoch {
	_ = "STUB: not implemented"

	// DefaultParams returns a default set of parameters
	return *new(Epoch)
}

func DefaultEpoch() Epoch { _ = "STUB: not implemented"; return *new(Epoch) }

func (e *Epoch) Validate() error { _ = "STUB: not implemented"; return nil }
