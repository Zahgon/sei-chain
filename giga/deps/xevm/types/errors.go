package types

type AssociationMissingErr struct {
	Address string
}

func NewAssociationMissingErr(address string) AssociationMissingErr {
	_ = "STUB: not implemented"
	return *new(AssociationMissingErr)
}

func (e AssociationMissingErr) Error() string { _ = "STUB: not implemented"; return "" }

func (e AssociationMissingErr) AddressType() string { _ = "STUB: not implemented"; return "" }
