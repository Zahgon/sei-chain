package types

import (
	"github.com/gogo/protobuf/jsonpb"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
)

var (
	ParamStoreKeyUploadAccess      = []byte("uploadAccess")
	ParamStoreKeyInstantiateAccess = []byte("instantiateAccess")
)

var AllAccessTypes = []AccessType{
	AccessTypeNobody,
	AccessTypeOnlyAddress,
	AccessTypeEverybody,
}

func (a AccessType) With(addr sdk.AccAddress) AccessConfig {
	_ = "STUB: not implemented"
	return *new(AccessConfig)
}

func (a AccessType) String() string { _ = "STUB: not implemented"; return "" }

func (a *AccessType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (a AccessType) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (a *AccessType) MarshalJSONPB(_ *jsonpb.Marshaler) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AccessType) UnmarshalJSONPB(_ *jsonpb.Unmarshaler, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (a AccessConfig) Equals(o AccessConfig) bool { _ = "STUB: not implemented"; return false }

var (
	DefaultUploadAccess = AllowEverybody
	AllowEverybody      = AccessConfig{Permission: AccessTypeEverybody}
	AllowNobody         = AccessConfig{Permission: AccessTypeNobody}
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	_ = "STUB: not implemented"
	return *new(paramtypes.KeyTable)
}

// DefaultParams returns default wasm parameters
func DefaultParams() Params { _ = "STUB: not implemented"; return *new(Params) }

func (p Params) String() string { _ = "STUB: not implemented"; return "" }

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	_ = "STUB: not implemented"
	return *new(paramtypes.ParamSetPairs)
}

// ValidateBasic performs basic validation on wasm parameters
func (p Params) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func validateAccessConfig(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateAccessType(i interface{}) error { _ = "STUB: not implemented"; return nil }

func (a AccessConfig) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (a AccessConfig) Allowed(actor sdk.AccAddress) bool { _ = "STUB: not implemented"; return false }
