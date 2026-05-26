package keeper

import (
	"encoding/json"

	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	ibcclienttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/types"
	wasmvmtypes "github.com/sei-protocol/sei-chain/sei-wasmvm/types"

	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
)

type (
	BankEncoder         func(sender sdk.AccAddress, msg *wasmvmtypes.BankMsg) ([]sdk.Msg, error)
	CustomEncoder       func(sender sdk.AccAddress, msg json.RawMessage, info wasmvmtypes.MessageInfo, codeInfo types.CodeInfo) ([]sdk.Msg, error)
	DistributionEncoder func(sender sdk.AccAddress, msg *wasmvmtypes.DistributionMsg) ([]sdk.Msg, error)
	StakingEncoder      func(sender sdk.AccAddress, msg *wasmvmtypes.StakingMsg) ([]sdk.Msg, error)
	StargateEncoder     func(sender sdk.AccAddress, msg *wasmvmtypes.StargateMsg) ([]sdk.Msg, error)
	WasmEncoder         func(sender sdk.AccAddress, msg *wasmvmtypes.WasmMsg) ([]sdk.Msg, error)
	IBCEncoder          func(ctx sdk.Context, sender sdk.AccAddress, contractIBCPortID string, msg *wasmvmtypes.IBCMsg) ([]sdk.Msg, error)
)

type MessageEncoders struct {
	Bank         func(sender sdk.AccAddress, msg *wasmvmtypes.BankMsg) ([]sdk.Msg, error)
	Custom       func(sender sdk.AccAddress, msg json.RawMessage, info wasmvmtypes.MessageInfo, codeInfo types.CodeInfo) ([]sdk.Msg, error)
	Distribution func(sender sdk.AccAddress, msg *wasmvmtypes.DistributionMsg) ([]sdk.Msg, error)
	IBC          func(ctx sdk.Context, sender sdk.AccAddress, contractIBCPortID string, msg *wasmvmtypes.IBCMsg) ([]sdk.Msg, error)
	Staking      func(sender sdk.AccAddress, msg *wasmvmtypes.StakingMsg) ([]sdk.Msg, error)
	Stargate     func(sender sdk.AccAddress, msg *wasmvmtypes.StargateMsg) ([]sdk.Msg, error)
	Wasm         func(sender sdk.AccAddress, msg *wasmvmtypes.WasmMsg) ([]sdk.Msg, error)
	Gov          func(sender sdk.AccAddress, msg *wasmvmtypes.GovMsg) ([]sdk.Msg, error)
}

func DefaultEncoders(unpacker codectypes.AnyUnpacker, portSource types.ICS20TransferPortSource) MessageEncoders {
	_ = "STUB: not implemented"
	return *new(MessageEncoders)
}

func (e MessageEncoders) Merge(o *MessageEncoders) MessageEncoders {
	_ = "STUB: not implemented"
	return *new(MessageEncoders)
}

func (e MessageEncoders) Encode(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg, info wasmvmtypes.MessageInfo, codeInfo types.CodeInfo) ([]sdk.Msg, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func EncodeBankMsg(sender sdk.AccAddress, msg *wasmvmtypes.BankMsg) ([]sdk.Msg, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NoCustomMsg(sender sdk.AccAddress, msg json.RawMessage, info wasmvmtypes.MessageInfo, _ types.CodeInfo) ([]sdk.Msg, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func EncodeDistributionMsg(sender sdk.AccAddress, msg *wasmvmtypes.DistributionMsg) ([]sdk.Msg, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func EncodeStakingMsg(sender sdk.AccAddress, msg *wasmvmtypes.StakingMsg) ([]sdk.Msg, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func EncodeStargateMsg(unpacker codectypes.AnyUnpacker) StargateEncoder {
	_ = "STUB: not implemented"
	return *new(StargateEncoder)
}

func EncodeWasmMsg(sender sdk.AccAddress, msg *wasmvmtypes.WasmMsg) ([]sdk.Msg, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func EncodeIBCMsg(portSource types.ICS20TransferPortSource) func(ctx sdk.Context, sender sdk.AccAddress, contractIBCPortID string, msg *wasmvmtypes.IBCMsg) ([]sdk.Msg, error) {
	_ = "STUB: not implemented"
	return nil
}

func EncodeGovMsg(sender sdk.AccAddress, msg *wasmvmtypes.GovMsg) ([]sdk.Msg, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConvertWasmIBCTimeoutHeightToCosmosHeight converts a wasmvm type ibc timeout height to ibc module type height
func ConvertWasmIBCTimeoutHeightToCosmosHeight(ibcTimeoutBlock *wasmvmtypes.IBCTimeoutBlock) ibcclienttypes.Height {
	_ = "STUB: not implemented"
	return *new(ibcclienttypes.Height)
}

// ConvertWasmCoinsToSdkCoins converts the wasm vm type coins to sdk type coins
func ConvertWasmCoinsToSdkCoins(coins []wasmvmtypes.Coin) (sdk.Coins, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Coins), nil
}

// ConvertWasmCoinToSdkCoin converts a wasm vm type coin to sdk type coin
func ConvertWasmCoinToSdkCoin(coin wasmvmtypes.Coin) (sdk.Coin, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Coin), nil
}
