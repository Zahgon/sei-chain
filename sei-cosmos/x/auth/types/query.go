package types

import codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"

func (m *QueryAccountResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

var _ codectypes.UnpackInterfacesMessage = &QueryAccountResponse{}
