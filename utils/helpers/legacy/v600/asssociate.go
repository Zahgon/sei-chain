package v600

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/sei-protocol/sei-chain/precompiles/utils"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

type AssociationHelper struct {
	evmKeeper     utils.EVMKeeper
	bankKeeper    utils.BankKeeper
	accountKeeper utils.AccountKeeper
}

func NewAssociationHelper(evmKeeper utils.EVMKeeper, bankKeeper utils.BankKeeper, accountKeeper utils.AccountKeeper) *AssociationHelper {
	_ = "STUB: not implemented"
	return nil
}

func (p AssociationHelper) AssociateAddresses(ctx sdk.Context, seiAddr sdk.AccAddress, evmAddr common.Address, pubkey cryptotypes.PubKey) error {
	_ = "STUB: not implemented"
	return nil
}

func (p AssociationHelper) MigrateBalance(ctx sdk.Context, evmAddr common.Address, seiAddr sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}
