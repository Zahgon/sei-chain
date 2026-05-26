package types

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	cryptocodec "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/codec"
)

// RegisterLegacyAminoCodec registers the account interfaces and concrete types on the
// provided LegacyAmino codec. These types are used for Amino JSON serialization
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) { _ = "STUB: not implemented"; return }

// RegisterInterfaces associates protoName with AccountI interface
// and creates a registry of it's concrete implementations
func RegisterInterfaces(registry types.InterfaceRegistry) { _ = "STUB: not implemented"; return }

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
}
