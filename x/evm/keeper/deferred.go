package keeper

import (
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/evm/types"
)

func (k *Keeper) GetAllEVMTxDeferredInfo(ctx sdk.Context) (res []*types.DeferredInfo) {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec

// this means the transaction got reverted during execution, either in ante handler
// or due to a panic in msg server

//nolint:gosec

// unable to unmarshal deferred info is serious, because it could cause
// balance surplus to be mishandled and thus affect total supply

func (k *Keeper) AppendToEvmTxDeferredInfo(ctx sdk.Context, bloom ethtypes.Bloom, txHash common.Hash, surplus sdk.Int) {
	_ = "STUB: not implemented"
	return
}

//nolint:gosec

//nolint:gosec

// unable to marshal deferred info is serious, because it could cause
// balance surplus to be mishandled and thus affect total supply

func (k *Keeper) SetNonceBumped(ctx sdk.Context) { _ = "STUB: not implemented"; return }

//nolint:gosec

func (k *Keeper) GetNonceBumped(ctx sdk.Context, txIndex uint32) bool {
	_ = "STUB: not implemented"
	return false
}

func (k *Keeper) GetEVMTxDeferredInfo(ctx sdk.Context) (*types.DeferredInfo, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

//nolint:gosec
