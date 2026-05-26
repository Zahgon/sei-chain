package wasm

import (
	"encoding/json"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func EncodeTokenFactoryCreateDenom(rawMsg json.RawMessage, sender sdk.AccAddress) ([]sdk.Msg, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func EncodeTokenFactoryMint(rawMsg json.RawMessage, sender sdk.AccAddress) ([]sdk.Msg, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func EncodeTokenFactoryBurn(rawMsg json.RawMessage, sender sdk.AccAddress) ([]sdk.Msg, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func EncodeTokenFactoryChangeAdmin(rawMsg json.RawMessage, sender sdk.AccAddress) ([]sdk.Msg, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func EncodeTokenFactorySetMetadata(rawMsg json.RawMessage, sender sdk.AccAddress) ([]sdk.Msg, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
