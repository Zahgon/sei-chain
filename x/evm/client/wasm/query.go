package wasm

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/evm/keeper"
)

type EVMQueryHandler struct {
	k *keeper.Keeper
}

func NewEVMQueryHandler(k *keeper.Keeper) *EVMQueryHandler { _ = "STUB: not implemented"; return nil }

func (h *EVMQueryHandler) HandleStaticCall(ctx sdk.Context, from string, to string, data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *EVMQueryHandler) HandleERC20TransferPayload(ctx sdk.Context, recipient string, amount *sdk.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *EVMQueryHandler) HandleERC20TokenInfo(ctx sdk.Context, contractAddress string, caller string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *EVMQueryHandler) HandleERC20Balance(ctx sdk.Context, contractAddress string, account string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *EVMQueryHandler) HandleERC721Owner(ctx sdk.Context, caller string, contractAddress string, tokenId string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *EVMQueryHandler) HandleERC721TransferPayload(ctx sdk.Context, from string, recipient string, tokenId string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *EVMQueryHandler) HandleERC721ApprovePayload(ctx sdk.Context, spender string, tokenId string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// empty address if approval should be revoked (i.e. spender string is empty)

func (h *EVMQueryHandler) HandleERC721SetApprovalAllPayload(ctx sdk.Context, to string, approved bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *EVMQueryHandler) HandleERC20TransferFromPayload(ctx sdk.Context, owner string, recipient string, amount *sdk.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *EVMQueryHandler) HandleERC20ApprovePayload(ctx sdk.Context, spender string, amount *sdk.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *EVMQueryHandler) HandleERC20Allowance(ctx sdk.Context, contractAddress string, owner string, spender string) ([]byte, error) {
	_ = "STUB: not implemented"
	// Get the evm address of the owner
	return nil, nil
}

// Get the evm address of spender

// Fetch the contract ABI

// Make the query to allowance(owner, spender)

// Parse the response (Should be of type uint256 if successful)

func (h *EVMQueryHandler) HandleERC721Approved(ctx sdk.Context, caller string, contractAddress string, tokenId string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *EVMQueryHandler) HandleERC721IsApprovedForAll(ctx sdk.Context, caller string, contractAddress string, owner string, operator string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *EVMQueryHandler) HandleERC721TotalSupply(ctx sdk.Context, caller string, contractAddress string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *EVMQueryHandler) HandleERC721NameSymbol(ctx sdk.Context, caller string, contractAddress string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *EVMQueryHandler) HandleERC721Uri(ctx sdk.Context, caller string, contractAddress string, tokenId string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *EVMQueryHandler) HandleERC721RoyaltyInfo(ctx sdk.Context, caller string, contractAddress string, tokenId string, salePrice *sdk.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *EVMQueryHandler) HandleERC1155TransferPayload(ctx sdk.Context, from string, recipient string, tokenId string, amount *sdk.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *EVMQueryHandler) HandleERC1155BatchTransferPayload(ctx sdk.Context, from string, recipient string, tokenIds []string, amounts []*sdk.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *EVMQueryHandler) HandleERC1155SetApprovalAllPayload(ctx sdk.Context, to string, approved bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *EVMQueryHandler) HandleERC1155IsApprovedForAll(ctx sdk.Context, caller string, contractAddress string, owner string, operator string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *EVMQueryHandler) HandleERC1155BalanceOf(ctx sdk.Context, caller string, contractAddress string, account string, tokenId string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *EVMQueryHandler) HandleERC1155BalanceOfBatch(ctx sdk.Context, caller string, contractAddress string, accounts []string, tokenIds []string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *EVMQueryHandler) HandleERC1155Uri(ctx sdk.Context, caller string, contractAddress string, tokenId string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *EVMQueryHandler) HandleERC1155TotalSupply(ctx sdk.Context, caller string, contractAddress string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *EVMQueryHandler) HandleERC1155TotalSupplyForToken(ctx sdk.Context, caller string, contractAddress string, tokenId string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *EVMQueryHandler) HandleERC1155TokenExists(ctx sdk.Context, caller string, contractAddress string, tokenId string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *EVMQueryHandler) HandleERC1155NameSymbol(ctx sdk.Context, caller string, contractAddress string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *EVMQueryHandler) HandleERC1155RoyaltyInfo(ctx sdk.Context, caller string, contractAddress string, tokenId string, salePrice *sdk.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *EVMQueryHandler) HandleSupportsInterface(ctx sdk.Context, caller string, id string, contractAddress string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *EVMQueryHandler) HandleGetEvmAddress(ctx sdk.Context, seiAddr string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *EVMQueryHandler) HandleGetSeiAddress(ctx sdk.Context, evmAddr string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
