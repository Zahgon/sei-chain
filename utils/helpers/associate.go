package helpers

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/sei-protocol/sei-chain/precompiles/utils"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

type AssociationHelper struct {
	evmKeeper     evmKeeper
	bankKeeper    bankKeeper
	accountKeeper utils.AccountKeeper
}

type evmKeeper interface {
	SetAddressMapping(ctx sdk.Context, seiAddress sdk.AccAddress, evmAddress common.Address)
}

type bankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SendCoins(sdk.Context, sdk.AccAddress, sdk.AccAddress, sdk.Coins) error
	GetWeiBalance(ctx sdk.Context, addr sdk.AccAddress) sdk.Int
	SendCoinsAndWei(ctx sdk.Context, from sdk.AccAddress, to sdk.AccAddress, amt sdk.Int, wei sdk.Int) error
	LockedCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
}

func NewAssociationHelper(evmKeeper evmKeeper, bankKeeper bankKeeper, accountKeeper utils.AccountKeeper) *AssociationHelper {
	_ = "STUB: not implemented"
	return nil
}

func (p AssociationHelper) AssociateAddresses(ctx sdk.Context, seiAddr sdk.AccAddress, evmAddr common.Address, pubkey cryptotypes.PubKey, migrateUseiOnly bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (p AssociationHelper) MigrateBalance(ctx sdk.Context, evmAddr common.Address, seiAddr sdk.AccAddress, migrateUseiOnly bool) error {
	_ = "STUB: not implemented"
	return nil
}
