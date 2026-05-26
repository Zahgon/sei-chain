package wasm

import (
	"encoding/json"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	wasmtypes "github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
	wasmvmtypes "github.com/sei-protocol/sei-chain/sei-wasmvm/types"
)

func EncodeCallEVM(rawMsg json.RawMessage, sender sdk.AccAddress, info wasmvmtypes.MessageInfo) ([]sdk.Msg, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func EncodeDelegateCallEVM(rawMsg json.RawMessage, sender sdk.AccAddress, info wasmvmtypes.MessageInfo, codeInfo wasmtypes.CodeInfo) ([]sdk.Msg, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
