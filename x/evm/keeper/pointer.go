package keeper

import (
	"github.com/ethereum/go-ethereum/common"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	sdkerrors "github.com/sei-protocol/sei-chain/sei-cosmos/types/errors"

	"github.com/sei-protocol/sei-chain/x/evm/types"
)

type PointerGetter func(sdk.Context, string) (common.Address, uint16, bool)
type PointerSetter func(sdk.Context, string, common.Address) error

var ErrorPointerToPointerNotAllowed = sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "cannot create a pointer to a pointer")

// ERC20 -> Native Token
func (k *Keeper) SetERC20NativePointer(ctx sdk.Context, token string, addr common.Address) error {
	_ = "STUB: not implemented"
	return nil
}

// ERC20 -> Native Token
func (k *Keeper) SetERC20NativePointerWithVersion(ctx sdk.Context, token string, addr common.Address, version uint16) error {
	_ = "STUB: not implemented"
	return nil
}

// ERC20 -> Native Token
func (k *Keeper) GetERC20NativePointer(ctx sdk.Context, token string) (addr common.Address, version uint16, exists bool) {
	_ = "STUB: not implemented"
	return *new(common.Address), 0, false
}

// ERC20 -> Native Token
func (k *Keeper) DeleteERC20NativePointer(ctx sdk.Context, token string, version uint16) {
	_ = "STUB: not implemented"
	return
}

// ERC20 -> CW20
func (k *Keeper) SetERC20CW20Pointer(ctx sdk.Context, cw20Address string, addr common.Address) error {
	_ = "STUB: not implemented"
	return nil
}

// ERC20 -> CW20
func (k *Keeper) SetERC20CW20PointerWithVersion(ctx sdk.Context, cw20Address string, addr common.Address, version uint16) error {
	_ = "STUB: not implemented"
	return nil
}

// ERC20 -> CW20
func (k *Keeper) GetERC20CW20Pointer(ctx sdk.Context, cw20Address string) (addr common.Address, version uint16, exists bool) {
	_ = "STUB: not implemented"
	return *new(common.Address), 0, false
}

// ERC20 -> CW20
func (k *Keeper) DeleteERC20CW20Pointer(ctx sdk.Context, cw20Address string, version uint16) {
	_ = "STUB: not implemented"
	return
}

// ERC721 -> CW721
func (k *Keeper) SetERC721CW721Pointer(ctx sdk.Context, cw721Address string, addr common.Address) error {
	_ = "STUB: not implemented"
	return nil
}

// ERC721 -> CW721
func (k *Keeper) SetERC721CW721PointerWithVersion(ctx sdk.Context, cw721Address string, addr common.Address, version uint16) error {
	_ = "STUB: not implemented"
	return nil
}

// ERC721 -> CW721
func (k *Keeper) GetERC721CW721Pointer(ctx sdk.Context, cw721Address string) (addr common.Address, version uint16, exists bool) {
	_ = "STUB: not implemented"
	return *new(common.Address), 0, false
}

// ERC721 -> CW721
func (k *Keeper) DeleteERC721CW721Pointer(ctx sdk.Context, cw721Address string, version uint16) {
	_ = "STUB: not implemented"
	return
}

// ERC1155 -> CW1155
func (k *Keeper) SetERC1155CW1155Pointer(ctx sdk.Context, cw1155Address string, addr common.Address) error {
	_ = "STUB: not implemented"
	return nil
}

// ERC1155 -> CW1155
func (k *Keeper) SetERC1155CW1155PointerWithVersion(ctx sdk.Context, cw1155Address string, addr common.Address, version uint16) error {
	_ = "STUB: not implemented"
	return nil
}

// ERC1155 -> CW1155
func (k *Keeper) GetERC1155CW1155Pointer(ctx sdk.Context, cw1155Address string) (addr common.Address, version uint16, exists bool) {
	_ = "STUB: not implemented"
	return *new(common.Address), 0, false
}

// ERC1155 -> CW1155
func (k *Keeper) DeleteERC1155CW1155Pointer(ctx sdk.Context, cw1155Address string, version uint16) {
	_ = "STUB: not implemented"
	return
}

// CW20 -> ERC20
func (k *Keeper) SetCW20ERC20Pointer(ctx sdk.Context, erc20Address common.Address, addr string) error {
	_ = "STUB: not implemented"
	return nil
}

// CW20 -> ERC20
func (k *Keeper) SetCW20ERC20PointerWithVersion(ctx sdk.Context, erc20Address common.Address, addr string, version uint16) error {
	_ = "STUB: not implemented"
	return nil
}

// CW20 -> ERC20
func (k *Keeper) GetCW20ERC20Pointer(ctx sdk.Context, erc20Address common.Address) (addr sdk.AccAddress, version uint16, exists bool) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), 0, false
}

// CW20 -> ERC20
func (k *Keeper) DeleteCW20ERC20Pointer(ctx sdk.Context, erc20Address common.Address, version uint16) {
	_ = "STUB: not implemented"
	return
}

func (k *Keeper) evmAddressIsPointer(ctx sdk.Context, addr common.Address) bool {
	_ = "STUB: not implemented"
	return false
}

func (k *Keeper) cwAddressIsPointer(ctx sdk.Context, addr string) bool {
	_ = "STUB: not implemented"
	return false
}

// CW721 -> ERC721
func (k *Keeper) SetCW721ERC721Pointer(ctx sdk.Context, erc721Address common.Address, addr string) error {
	_ = "STUB: not implemented"
	return nil
}

// CW721 -> ERC721
func (k *Keeper) SetCW721ERC721PointerWithVersion(ctx sdk.Context, erc721Address common.Address, addr string, version uint16) error {
	_ = "STUB: not implemented"
	return nil
}

// CW721 -> ERC721
func (k *Keeper) GetCW721ERC721Pointer(ctx sdk.Context, erc721Address common.Address) (addr sdk.AccAddress, version uint16, exists bool) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), 0, false
}

// CW721 -> ERC721
func (k *Keeper) DeleteCW721ERC721Pointer(ctx sdk.Context, erc721Address common.Address, version uint16) {
	_ = "STUB: not implemented"
	return
}

// CW1155 -> ERC1155
func (k *Keeper) SetCW1155ERC1155Pointer(ctx sdk.Context, erc1155Address common.Address, addr string) error {
	_ = "STUB: not implemented"
	return nil
}

// CW1155 -> ERC1155
func (k *Keeper) SetCW1155ERC1155PointerWithVersion(ctx sdk.Context, erc1155Address common.Address, addr string, version uint16) error {
	_ = "STUB: not implemented"
	return nil
}

// CW1155 -> ERC1155
func (k *Keeper) GetCW1155ERC1155Pointer(ctx sdk.Context, erc1155Address common.Address) (addr sdk.AccAddress, version uint16, exists bool) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), 0, false
}

// CW1155 -> ERC1155
func (k *Keeper) DeleteCW1155ERC1155Pointer(ctx sdk.Context, erc1155Address common.Address, version uint16) {
	_ = "STUB: not implemented"
	return
}

func (k *Keeper) GetPointerInfo(ctx sdk.Context, pref []byte, maxVersion uint16) (addr []byte, version uint16, exists bool) {
	_ = "STUB: not implemented"
	return nil, 0, false
}

//nolint:gosec

func (k *Keeper) GetAnyPointeeInfo(ctx sdk.Context, cwAddress string) (common.Address, uint16, bool) {
	_ = "STUB: not implemented"
	return *new(common.Address), 0, false
}

func (k *Keeper) GetAnyPointerInfo(ctx sdk.Context, pref []byte) (addr []byte, version uint16, exists bool) {
	_ = "STUB: not implemented"
	return nil, 0, false
}

//nolint:gosec

func (k *Keeper) setPointerInfo(ctx sdk.Context, pref []byte, addr []byte, version uint16) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Keeper) deletePointerInfo(ctx sdk.Context, pref []byte, version uint16) {
	_ = "STUB: not implemented"
	return
}

func maxCurrentPointerVersion(ctx sdk.Context) uint16 { _ = "STUB: not implemented"; return 0 }

func (k *Keeper) GetStoredPointerCodeID(ctx sdk.Context, pointerType types.PointerType) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func (k *Keeper) GetCW20Pointee(ctx sdk.Context, erc20Address common.Address) (cw20Address string, version uint16, exists bool) {
	_ = "STUB: not implemented"
	return "", 0, false
}

func (k *Keeper) GetCW721Pointee(ctx sdk.Context, erc721Address common.Address) (cw721Address string, version uint16, exists bool) {
	_ = "STUB: not implemented"
	return "", 0, false
}

func (k *Keeper) GetCW1155Pointee(ctx sdk.Context, erc1155Address common.Address) (cw1155Address string, version uint16, exists bool) {
	_ = "STUB: not implemented"
	return "", 0, false
}

func (k *Keeper) GetERC20Pointee(ctx sdk.Context, cw20Address string) (erc20Address common.Address, version uint16, exists bool) {
	_ = "STUB: not implemented"
	return *new(common.Address), 0, false
}

func (k *Keeper) GetERC721Pointee(ctx sdk.Context, cw721Address string) (erc721Address common.Address, version uint16, exists bool) {
	_ = "STUB: not implemented"
	return *new(common.Address), 0, false
}

func (k *Keeper) GetERC1155Pointee(ctx sdk.Context, cw1155Address string) (erc1155Address common.Address, version uint16, exists bool) {
	_ = "STUB: not implemented"
	return *new(common.Address), 0, false
}

func (k *Keeper) GetNativePointee(ctx sdk.Context, erc20Address string) (token string, version uint16, exists bool) {
	_ = "STUB: not implemented"
	// Ensure the key matches how it was set in SetERC20NativePointer
	return "", 0, false
}
