package wasmbinding

import (
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	authkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/keeper"
	stakingkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/keeper"
	wasmkeeper "github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/keeper"
	wasmtypes "github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
	epochkeeper "github.com/sei-protocol/sei-chain/x/epoch/keeper"
	evmkeeper "github.com/sei-protocol/sei-chain/x/evm/keeper"
	oraclekeeper "github.com/sei-protocol/sei-chain/x/oracle/keeper"
	tokenfactorykeeper "github.com/sei-protocol/sei-chain/x/tokenfactory/keeper"
)

func RegisterCustomPlugins(
	oracle *oraclekeeper.Keeper,
	epoch *epochkeeper.Keeper,
	tokenfactory *tokenfactorykeeper.Keeper,
	_ *authkeeper.AccountKeeper,
	router wasmkeeper.MessageRouter,
	channelKeeper wasmtypes.ChannelKeeper,
	capabilityKeeper wasmtypes.CapabilityKeeper,
	bankKeeper wasmtypes.Burner,
	unpacker codectypes.AnyUnpacker,
	portSource wasmtypes.ICS20TransferPortSource,
	evmKeeper *evmkeeper.Keeper,
	stakingKeeper stakingkeeper.Keeper,
) []wasmkeeper.Option {
	_ = "STUB: not implemented"
	return nil
}
