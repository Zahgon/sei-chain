package wasmbinding

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/baseapp"
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	wasmkeeper "github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/keeper"
	wasmtypes "github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
	evmkeeper "github.com/sei-protocol/sei-chain/x/evm/keeper"
)

type CustomRouter struct {
	wasmkeeper.MessageRouter

	evmKeeper *evmkeeper.Keeper
}

func (r *CustomRouter) Handler(msg sdk.Msg) baseapp.MsgServiceHandler {
	_ = "STUB: not implemented"
	return *new(baseapp.MsgServiceHandler)
}

// forked from wasm
func CustomMessageHandler(
	router wasmkeeper.MessageRouter,
	channelKeeper wasmtypes.ChannelKeeper,
	capabilityKeeper wasmtypes.CapabilityKeeper,
	bankKeeper wasmtypes.Burner,
	evmKeeper *evmkeeper.Keeper,
	unpacker codectypes.AnyUnpacker,
	portSource wasmtypes.ICS20TransferPortSource,
) wasmkeeper.Messenger {
	_ = "STUB: not implemented"
	return *new(wasmkeeper.Messenger)
}
