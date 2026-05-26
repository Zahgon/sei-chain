package keeper

import (
	"github.com/ethereum/go-ethereum/common"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const CoinbaseSeedAddress = "0000000000000000000000000000000000000001"
const CoinbaseNonce = 42

func (k *Keeper) GetFeeCollectorAddress(ctx sdk.Context) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// we don't want to charge gas for this query, since it could cause non-determinism

// ok to write multiple times since it's idempotent

func GetCoinbaseAddress() common.Address { _ = "STUB: not implemented"; return *new(common.Address) }
