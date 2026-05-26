// nolint
package testutils

import (
	"testing"

	"github.com/sei-protocol/sei-chain/x/oracle/utils"

	appparams "github.com/sei-protocol/sei-chain/app/params"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/keys/secp256k1"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth"
	authkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/keeper"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/bank"
	bankkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/keeper"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/distribution"
	distrkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/distribution/keeper"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/params"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/staking"
	stakingkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/keeper"
	stakingtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
	oraclekeeper "github.com/sei-protocol/sei-chain/x/oracle/keeper"
)

const faucetAccountName = "faucet"

// ModuleBasics nolint
var ModuleBasics = module.NewBasicManager(
	auth.AppModuleBasic{},
	bank.AppModuleBasic{},
	distribution.AppModuleBasic{},
	staking.AppModuleBasic{},
	params.AppModuleBasic{},
)

// MakeTestCodec nolint
func MakeTestCodec(t *testing.T) codec.Codec { _ = "STUB: not implemented"; return *new(codec.Codec) }

// MakeEncodingConfig nolint
func MakeEncodingConfig(_ *testing.T) appparams.EncodingConfig {
	_ = "STUB: not implemented"
	return *new(appparams.EncodingConfig)
}

// Test addresses
var (
	ValPubKeys = CreateTestPubKeys(7)

	pubKeys = []cryptotypes.PubKey{
		secp256k1.GenPrivKey().PubKey(),
		secp256k1.GenPrivKey().PubKey(),
		secp256k1.GenPrivKey().PubKey(),
		secp256k1.GenPrivKey().PubKey(),
		secp256k1.GenPrivKey().PubKey(),
		secp256k1.GenPrivKey().PubKey(),
		secp256k1.GenPrivKey().PubKey(),
	}

	Addrs = []sdk.AccAddress{
		sdk.AccAddress(pubKeys[0].Address()),
		sdk.AccAddress(pubKeys[1].Address()),
		sdk.AccAddress(pubKeys[2].Address()),
		sdk.AccAddress(pubKeys[3].Address()),
		sdk.AccAddress(pubKeys[4].Address()),
		sdk.AccAddress(pubKeys[5].Address()),
		sdk.AccAddress(pubKeys[6].Address()),
	}

	ValAddrs = []sdk.ValAddress{
		sdk.ValAddress(pubKeys[0].Address()),
		sdk.ValAddress(pubKeys[1].Address()),
		sdk.ValAddress(pubKeys[2].Address()),
		sdk.ValAddress(pubKeys[3].Address()),
		sdk.ValAddress(pubKeys[4].Address()),
		sdk.ValAddress(pubKeys[5].Address()),
		sdk.ValAddress(pubKeys[6].Address()),
	}

	InitTokens = sdk.TokensFromConsensusPower(200, sdk.DefaultPowerReduction)
	InitCoins  = sdk.NewCoins(sdk.NewCoin(utils.MicroSeiDenom, InitTokens))

	OracleDecPrecision = 8
)

// TestInput nolint
type TestInput struct {
	Ctx           sdk.Context
	Cdc           *codec.LegacyAmino
	AccountKeeper authkeeper.AccountKeeper
	BankKeeper    bankkeeper.Keeper
	OracleKeeper  oraclekeeper.Keeper
	StakingKeeper stakingkeeper.Keeper
	DistrKeeper   distrkeeper.Keeper
}

// CreateTestInput nolint
func CreateTestInput(t *testing.T) TestInput { _ = "STUB: not implemented"; return *new(TestInput) }

// NewTestMsgCreateValidator test msg creator
func NewTestMsgCreateValidator(address sdk.ValAddress, pubKey cryptotypes.PubKey, amt sdk.Int) *stakingtypes.MsgCreateValidator {
	_ = "STUB: not implemented"
	return nil
}

// FundAccount is a utility function that funds an account by minting and
// sending the coins to the address. This should be used for testing purposes
// only!
func FundAccount(input TestInput, addr sdk.AccAddress, amounts sdk.Coins) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateTestPubKeys returns a total of numPubKeys public keys in ascending order.
func CreateTestPubKeys(numPubKeys int) []cryptotypes.PubKey { _ = "STUB: not implemented"; return nil }

// start at 10 to avoid changing 1 to 01, 2 to 02, etc

// base pubkey string
// adding on final two digits to make pubkeys unique

// NewPubKeyFromHex returns a PubKey from a hex string.
func NewPubKeyFromHex(pk string) (res cryptotypes.PubKey) {
	_ = "STUB: not implemented"
	return *new(cryptotypes.PubKey)
}
