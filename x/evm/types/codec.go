package types

//nolint:typecheck
import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	cryptocodec "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/codec"

	"github.com/sei-protocol/sei-chain/x/evm/types/ethtx"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

func GetAmino() *codec.LegacyAmino { _ = "STUB: not implemented"; return nil }

func RegisterCodec(cdc *codec.LegacyAmino) { _ = "STUB: not implemented"; return }

func RegisterInterfaces(registry codectypes.InterfaceRegistry) { _ = "STUB: not implemented"; return }

func PackTxData(txData ethtx.TxData) (*codectypes.Any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func UnpackTxData(any *codectypes.Any) (ethtx.TxData, error) {
	_ = "STUB: not implemented"
	return *new(ethtx.TxData), nil
}

// value is a legacy tx

// value is a accesslist tx

// value is a dynamic fee tx

// value is a blob tx

// value is an associate tx

// value is a set code tx
