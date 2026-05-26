package wasmtesting

import sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

type MockCoinTransferrer struct {
	TransferCoinsFn func(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
}

func (m *MockCoinTransferrer) TransferCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error {
	_ = "STUB: not implemented"
	return nil
}
