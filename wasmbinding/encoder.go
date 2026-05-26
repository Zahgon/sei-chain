package wasmbinding

import (
	"encoding/json"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	wasmtypes "github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
	wasmvmtypes "github.com/sei-protocol/sei-chain/sei-wasmvm/types"
)

type SeiWasmMessage struct {
	CreateDenom     json.RawMessage `json:"create_denom,omitempty"`
	MintTokens      json.RawMessage `json:"mint_tokens,omitempty"`
	BurnTokens      json.RawMessage `json:"burn_tokens,omitempty"`
	ChangeAdmin     json.RawMessage `json:"change_admin,omitempty"`
	SetMetadata     json.RawMessage `json:"set_metadata,omitempty"`
	CallEVM         json.RawMessage `json:"call_evm,omitempty"`
	DelegateCallEVM json.RawMessage `json:"delegate_call_evm,omitempty"`
}

func CustomEncoder(sender sdk.AccAddress, msg json.RawMessage, info wasmvmtypes.MessageInfo, codeInfo wasmtypes.CodeInfo) ([]sdk.Msg, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
