package apptesting

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/baseapp"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/tx/signing"
	authsigning "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/signing"
	stakingtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
	"github.com/stretchr/testify/suite"

	"github.com/sei-protocol/sei-chain/app"
)

type KeeperTestHelper struct {
	suite.Suite

	App         *app.App
	Ctx         sdk.Context
	QueryHelper *baseapp.QueryServiceTestHelper
	TestAccs    []sdk.AccAddress
}

// Setup sets up basic environment for suite (App, Ctx, and test accounts)
func (s *KeeperTestHelper) Setup() { _ = "STUB: not implemented"; return }

// CreateTestContext creates a test context.
func (s *KeeperTestHelper) CreateTestContext() sdk.Context {
	_ = "STUB: not implemented"
	return *new(sdk.Context)
}

// CreateTestContext creates a test context.
func (s *KeeperTestHelper) Commit() { _ = "STUB: not implemented"; return }

// FundAcc funds target address with specified amount.
func (s *KeeperTestHelper) FundAcc(acc sdk.AccAddress, amounts sdk.Coins) {
	_ = "STUB: not implemented"
	return
}

// SetupValidator sets up a validator and returns the ValAddress.
func (s *KeeperTestHelper) SetupValidator(bondStatus stakingtypes.BondStatus) sdk.ValAddress {
	_ = "STUB: not implemented"
	return *new(sdk.ValAddress)
}

// SetupTokenFactory sets up a token module account for the TokenFactoryKeeper.
func (s *KeeperTestHelper) SetupTokenFactory() { _ = "STUB: not implemented"; return }

// EndBlock ends the block.
func (s *KeeperTestHelper) EndBlock() { _ = "STUB: not implemented"; return }

// AllocateRewardsToValidator allocates reward tokens to a distribution module then allocates rewards to the validator address.
func (s *KeeperTestHelper) AllocateRewardsToValidator(valAddr sdk.ValAddress, rewardAmt sdk.Int) {
	_ = "STUB: not implemented"
	return
}

// allocate reward tokens to distribution module

// allocate rewards to validator

// BuildTx builds a transaction.
func (s *KeeperTestHelper) BuildTx(
	txBuilder client.TxBuilder,
	msgs []sdk.Msg,
	sigV2 signing.SignatureV2,
	memo string,
	txFee sdk.Coins,
	gasLimit uint64,
) authsigning.Tx {
	_ = "STUB: not implemented"
	return *new(authsigning.Tx)
}

// CreateRandomAccounts is a function return a list of randomly generated AccAddresses
func CreateRandomAccounts(numAccts int) []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

func GenerateTestAddrs() (string, string) { _ = "STUB: not implemented"; return "", "" }

type bankKeeper interface {
	MintCoins(ctx sdk.Context, moduleName string, amounts sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, moduleName string, recipient sdk.AccAddress, amounts sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule string, recipientModule string, amounts sdk.Coins) error
}

func FundAccount(bankKeeper bankKeeper, ctx sdk.Context, addr sdk.AccAddress, amounts sdk.Coins) error {
	_ = "STUB: not implemented"
	return nil
}

func FundModuleAccount(bankKeeper bankKeeper, ctx sdk.Context, recipientMod string, amounts sdk.Coins) error {
	_ = "STUB: not implemented"
	return nil
}
