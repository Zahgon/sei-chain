package keeper

import (
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/sei-protocol/sei-chain/giga/deps/xevm/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func (k *Keeper) GetBlockBloom(ctx sdk.Context) (res ethtypes.Bloom) {
	_ = "STUB: not implemented"
	return *new(ethtypes.Bloom)
}

func (k *Keeper) GetEvmOnlyBlockBloom(ctx sdk.Context) (res ethtypes.Bloom) {
	_ = "STUB: not implemented"
	return *new(ethtypes.Bloom)
}

func (k *Keeper) GetLegacyBlockBloom(ctx sdk.Context, height int64) (res ethtypes.Bloom) {
	_ = "STUB: not implemented"
	return *new(ethtypes.Bloom)
}

func (k *Keeper) SetEvmOnlyBlockBloom(ctx sdk.Context, blooms []ethtypes.Bloom) {
	_ = "STUB: not implemented"
	return
}

func (k *Keeper) SetBlockBloom(ctx sdk.Context, blooms []ethtypes.Bloom) {
	_ = "STUB: not implemented"
	return
}

func BloomsToBytes(blooms []ethtypes.Bloom) []byte { _ = "STUB: not implemented"; return nil }

func (k *Keeper) SetLegacyBlockBloomCutoffHeight(ctx sdk.Context) {
	_ = "STUB: not implemented"
	return
}

//nolint:gosec

func (k *Keeper) GetLegacyBlockBloomCutoffHeight(ctx sdk.Context) int64 {
	_ = "STUB: not implemented"
	return 0
}

//nolint:gosec

func GetLogsForTx(receipt *types.Receipt, logStartIndex uint) []*ethtypes.Log {
	_ = "STUB: not implemented"
	return nil
}

func convertLog(l *types.Log, receipt *types.Receipt, logStartIndex uint) *ethtypes.Log {
	_ = "STUB: not implemented"
	return nil
}

func ConvertEthLog(l *ethtypes.Log) *types.Log { _ = "STUB: not implemented"; return nil }

//nolint:gosec

func ConvertSyntheticEthLog(l *ethtypes.Log) *types.Log { _ = "STUB: not implemented"; return nil }
