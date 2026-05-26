package types

// String implements fmt.Stringer interface
func (d Denom) String() string { _ = "STUB: not implemented"; return "" }

// Equal implements equal interface
func (d Denom) Equal(d1 *Denom) bool { _ = "STUB: not implemented"; return false }

// DenomList is array of Denom
type DenomList []Denom

func (dl DenomList) Contains(denom string) bool { _ = "STUB: not implemented"; return false }

// String implements fmt.Stringer interface
func (dl DenomList) String() (out string) { _ = "STUB: not implemented"; return "" }
