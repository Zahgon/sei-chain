package sync

import (
	"sync"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

type GasWrapper struct {
	sdk.GasMeter
	mu *sync.Mutex
}

func NewGasWrapper(wrapped sdk.GasMeter) sdk.GasMeter {
	_ = "STUB: not implemented"
	return *new(sdk.GasMeter)
}

func (g GasWrapper) ConsumeGas(amount sdk.Gas, descriptor string) {
	_ = "STUB: not implemented"
	return
}

func (g GasWrapper) RefundGas(amount sdk.Gas, descriptor string) { _ = "STUB: not implemented"; return }
