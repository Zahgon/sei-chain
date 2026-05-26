package keeper

import (
	"sync"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// SdkIntPool is a pool of *sdk.Int for reuse across balance reads from storage.
// sdk.Int.Unmarshal reuses the existing *big.Int when non-nil, so a pooled entry
// that has been used before skips the big.Int allocation on subsequent unmarshal
// calls. Callers must Put the pointer back after use and must not retain it.
type SdkIntPool struct {
	p sync.Pool
}

func newSdkIntPool() *SdkIntPool { _ = "STUB: not implemented"; return nil }

func (s *SdkIntPool) Get() *sdk.Int { _ = "STUB: not implemented"; return nil }

func (s *SdkIntPool) Put(i *sdk.Int) { _ = "STUB: not implemented"; return }
