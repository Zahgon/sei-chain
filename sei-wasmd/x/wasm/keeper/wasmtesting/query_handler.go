package wasmtesting

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	wasmvmtypes "github.com/sei-protocol/sei-chain/sei-wasmvm/types"
)

type MockQueryHandler struct {
	HandleQueryFn func(ctx sdk.Context, request wasmvmtypes.QueryRequest, caller sdk.AccAddress) ([]byte, error)
}

func (m *MockQueryHandler) HandleQuery(ctx sdk.Context, caller sdk.AccAddress, request wasmvmtypes.QueryRequest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
