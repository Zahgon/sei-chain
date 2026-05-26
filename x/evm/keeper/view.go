package keeper

import (
	"github.com/ethereum/go-ethereum/common"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func (k *Keeper) QueryERCSingleOutput(ctx sdk.Context, typ string, addr common.Address, query string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
