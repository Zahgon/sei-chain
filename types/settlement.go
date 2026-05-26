package types

type Settlement struct {
	To       string
	Quantity uint64
	Denom    string
}

func (s Settlement) String() string { _ = "STUB: not implemented"; return "" }
