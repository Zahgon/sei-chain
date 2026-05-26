package mint

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/mint/keeper"
	"github.com/sei-protocol/sei-chain/x/mint/types"
)

func HandleUpdateMinterProposal(ctx sdk.Context, k *keeper.Keeper, p *types.UpdateMinterProposal) error {
	_ = "STUB: not implemented"
	return nil
}
