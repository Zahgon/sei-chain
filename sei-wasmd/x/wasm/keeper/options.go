package keeper

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
)

type optsFn func(*Keeper)

func (f optsFn) apply(keeper *Keeper) {
	_ = "STUB: not implemented"

	// WithWasmEngine is an optional constructor parameter to replace the default wasmVM engine with the
	// given one.
	return
}

func WithWasmEngine(x types.WasmerEngine) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMessageHandler is an optional constructor parameter to set a custom handler for wasmVM messages.
// This option should not be combined with Option `WithMessageEncoders` or `WithMessageHandlerDecorator`
func WithMessageHandler(x Messenger) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMessageHandlerDecorator is an optional constructor parameter to decorate the wasm handler for wasmVM messages.
// This option should not be combined with Option `WithMessageEncoders` or `WithMessageHandler`
func WithMessageHandlerDecorator(d func(old Messenger) Messenger) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithQueryHandler is an optional constructor parameter to set custom query handler for wasmVM requests.
// This option should not be combined with Option `WithQueryPlugins` or `WithQueryHandlerDecorator`
func WithQueryHandler(x WasmVMQueryHandler) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithQueryHandlerDecorator is an optional constructor parameter to decorate the default wasm query handler for wasmVM requests.
// This option should not be combined with Option `WithQueryPlugins` or `WithQueryHandler`
func WithQueryHandlerDecorator(d func(old WasmVMQueryHandler) WasmVMQueryHandler) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithQueryPlugins is an optional constructor parameter to pass custom query plugins for wasmVM requests.
// This option expects the default `QueryHandler` set an should not be combined with Option `WithQueryHandler` or `WithQueryHandlerDecorator`.
func WithQueryPlugins(x *QueryPlugins) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMessageEncoders is an optional constructor parameter to pass custom message encoder to the default wasm message handler.
// This option expects the `DefaultMessageHandler` set and should not be combined with Option `WithMessageHandler` or `WithMessageHandlerDecorator`.
func WithMessageEncoders(x *MessageEncoders) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCoinTransferrer is an optional constructor parameter to set a custom coin transferrer
func WithCoinTransferrer(x CoinTransferrer) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithVMCacheMetrics(r prometheus.Registerer) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithGasRegister set a new gas register to implement custom gas costs.
// When the "gas multiplier" for wasmvm gas conversion is modified inside the new register,
// make sure to also use `WithApiCosts` option for non default values
func WithGasRegister(x GasRegister) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAPICosts sets custom api costs. Amounts are in cosmwasm gas Not SDK gas.
func WithAPICosts(human, canonical uint64) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMaxQueryStackSize overwrites the default limit for maximum query stacks
func WithMaxQueryStackSize(m uint32) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMaxCallDepth(m uint32) Option { _ = "STUB: not implemented"; return *new(Option) }
