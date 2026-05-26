package types

import (
	"reflect"
)

type attribute struct {
	ty  reflect.Type
	vfn ValueValidatorFn
}

// KeyTable subspaces appropriate type for each parameter key
type KeyTable struct {
	m map[string]attribute
}

func NewKeyTable(pairs ...ParamSetPair) KeyTable { _ = "STUB: not implemented"; return *new(KeyTable) }

// RegisterType registers a single ParamSetPair (key-type pair) in a KeyTable.
func (t KeyTable) RegisterType(psp ParamSetPair) KeyTable {
	_ = "STUB: not implemented"
	return *new(KeyTable)
}

// indirect rty if it is a pointer

// RegisterParamSet registers multiple ParamSetPairs from a ParamSet in a KeyTable.
func (t KeyTable) RegisterParamSet(ps ParamSet) KeyTable {
	_ = "STUB: not implemented"
	return *new(KeyTable)
}

func (t KeyTable) maxKeyLength() (res int) { _ = "STUB: not implemented"; return 0 }
