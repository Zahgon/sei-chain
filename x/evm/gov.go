package evm

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/evm/keeper"
	"github.com/sei-protocol/sei-chain/x/evm/types"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("x", "evm")

func HandleAddERCNativePointerProposalV2(ctx sdk.Context, k *keeper.Keeper, p *types.AddERCNativePointerProposalV2) error {
	_ = "STUB: not implemented"
	return nil
}

// should always be the case given validation
//nolint:gosec

func logNativeV2Error(ctx sdk.Context, p *types.AddERCNativePointerProposalV2, step string, err string) {
	_ = "STUB: not implemented"
	return
}

func HandleAddERCNativePointerProposal(ctx sdk.Context, k *keeper.Keeper, p *types.AddERCNativePointerProposal) error {
	_ = "STUB: not implemented"
	return nil
}

func HandleAddERCCW20PointerProposal(ctx sdk.Context, k *keeper.Keeper, p *types.AddERCCW20PointerProposal) error {
	_ = "STUB: not implemented"
	return nil
}

func HandleAddERCCW721PointerProposal(ctx sdk.Context, k *keeper.Keeper, p *types.AddERCCW721PointerProposal) error {
	_ = "STUB: not implemented"
	return nil
}

func HandleAddERCCW1155PointerProposal(ctx sdk.Context, k *keeper.Keeper, p *types.AddERCCW1155PointerProposal) error {
	_ = "STUB: not implemented"
	return nil
}

func HandleAddCWERC20PointerProposal(ctx sdk.Context, k *keeper.Keeper, p *types.AddCWERC20PointerProposal) error {
	_ = "STUB: not implemented"
	return nil
}

func HandleAddCWERC721PointerProposal(ctx sdk.Context, k *keeper.Keeper, p *types.AddCWERC721PointerProposal) error {
	_ = "STUB: not implemented"
	return nil
}

func HandleAddCWERC1155PointerProposal(ctx sdk.Context, k *keeper.Keeper, p *types.AddCWERC1155PointerProposal) error {
	_ = "STUB: not implemented"
	return nil
}
