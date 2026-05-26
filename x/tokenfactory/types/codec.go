package types

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	cdctypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	// this line is used by starport scaffolding # 1
)

func RegisterCodec(cdc *codec.LegacyAmino) { _ = "STUB: not implemented"; return }

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) { _ = "STUB: not implemented"; return }

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterCodec(amino)
	sdk.RegisterLegacyAminoCodec(amino)

	amino.Seal()
}
