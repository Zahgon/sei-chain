package types

import (
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func (s Sequence) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (s GenesisState) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (s GenesisState) ValidateBasicStream(dataCh chan GenesisState) error {
	_ = "STUB: not implemented"
	return nil
}

func (c Code) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (c Contract) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// AsMsg returns the underlying cosmos-sdk message instance. Null when can not be mapped to a known type.
func (m GenesisState_GenMsgs) AsMsg() sdk.Msg { _ = "STUB: not implemented"; return *new(sdk.Msg) }

func (m GenesisState_GenMsgs) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// ValidateGenesis performs basic validation of supply genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data GenesisState) error { _ = "STUB: not implemented"; return nil }

// ValidateGenesisStream performs basic validation of wasm genesis data over a stream
// of wasm genesis states. It needs to pass the params validation at least one chunk and
// other checks on every chunk of the stream.
func ValidateGenesisStream(genesisStateCh <-chan GenesisState) error {
	_ = "STUB: not implemented"
	return nil
}

var _ codectypes.UnpackInterfacesMessage = GenesisState{}

// UnpackInterfaces implements codectypes.UnpackInterfaces
func (s GenesisState) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

var _ codectypes.UnpackInterfacesMessage = &Contract{}

// UnpackInterfaces implements codectypes.UnpackInterfaces
func (c *Contract) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}
