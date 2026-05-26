package baseapp

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
)

func (app *BaseApp) Check(txEncoder sdk.TxEncoder, tx sdk.Tx) (sdk.GasInfo, *sdk.Result, error) {
	_ = "STUB: not implemented"
	// runTx expects tx bytes as argument, so we encode the tx argument into
	// bytes. Note that runTx will actually decode those bytes again. But since
	// this helper is only used in tests/simulation, it's fine.
	return *new(sdk.GasInfo), nil, nil
}

func (app *BaseApp) Deliver(txEncoder sdk.TxEncoder, tx sdk.Tx) (sdk.GasInfo, *sdk.Result, error) {
	_ = "STUB: not implemented"
	// See comment for Check().
	return *new(sdk.GasInfo), nil, nil
}

// Context with current {check, deliver}State of the app used by tests.
func (app *BaseApp) NewContext(isCheckTx bool, header tmproto.Header) sdk.Context {
	_ = "STUB: not implemented"
	return *new(sdk.Context)
}

func (app *BaseApp) NewUncachedContext(isCheckTx bool, header tmproto.Header) sdk.Context {
	_ = "STUB: not implemented"
	return *new(sdk.Context)
}

// DeliverContext returns the current deliverState context, or nil if not in a deliver block.
// Useful for reading uncommitted state (e.g. after InitChain before Commit).
func (app *BaseApp) DeliverContext() *sdk.Context { _ = "STUB: not implemented"; return nil }

func (app *BaseApp) GetContextForDeliverTx(txBytes []byte) sdk.Context {
	_ = "STUB: not implemented"
	return *new(sdk.Context)
}
