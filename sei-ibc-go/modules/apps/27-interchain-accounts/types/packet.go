package types

import (
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
)

// MaxMemoCharLength defines the maximum length for the InterchainAccountPacketData memo field
const MaxMemoCharLength = 256

// ValidateBasic performs basic validation of the interchain account packet data.
// The memo may be empty.
func (iapd InterchainAccountPacketData) ValidateBasic() error {
	_ = "STUB: not implemented"
	return nil
}

// GetBytes returns the JSON marshalled interchain account packet data.
func (iapd InterchainAccountPacketData) GetBytes() []byte { _ = "STUB: not implemented"; return nil }

// GetBytes returns the JSON marshalled interchain account CosmosTx.
func (ct CosmosTx) GetBytes() []byte { _ = "STUB: not implemented"; return nil }

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (ct CosmosTx) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}
