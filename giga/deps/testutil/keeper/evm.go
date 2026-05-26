package keeper

import (
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/app"
	evmkeeper "github.com/sei-protocol/sei-chain/giga/deps/xevm/keeper"
	evmtypes "github.com/sei-protocol/sei-chain/giga/deps/xevm/types"
)

var mockKeeper *evmkeeper.Keeper
var mockCtx sdk.Context
var mtx = &sync.Mutex{}

func MockApp(t *testing.T) (*app.App, sdk.Context) {
	_ = "STUB: not implemented"
	return nil, *new(sdk.Context)
}

func MockEVMKeeper(t *testing.T) (*evmkeeper.Keeper, sdk.Context) {
	_ = "STUB: not implemented"
	return nil, *new(sdk.Context)
}

// mint some coins to a sei address

func MockEVMKeeperPrecompiles(t *testing.T) (*evmkeeper.Keeper, sdk.Context) {
	_ = "STUB: not implemented"
	return nil, *new(sdk.Context)
}

// mint some coins to a sei address

func MockAddressPair() (sdk.AccAddress, common.Address) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), *new(common.Address)
}

func MockPrivateKey() cryptotypes.PrivKey {
	_ = "STUB: not implemented"
	// Generate a new Sei private key
	return *new(cryptotypes.PrivKey)
}

func PrivateKeyToAddresses(privKey cryptotypes.PrivKey) (sdk.AccAddress, common.Address) {
	_ = "STUB: not implemented"
	// Encode the private key to hex (i.e. what wallets do behind the scene when users reveal private keys)
	return *new(sdk.AccAddress), *new(common.Address)
}

// Sign an Ethereum transaction with the hex private key

// Recover the public keys from the Ethereum signature

func UseiCoins(amount int64) sdk.Coins { _ = "STUB: not implemented"; return *new(sdk.Coins) }

func WaitForReceipt(t *testing.T, k *evmkeeper.Keeper, ctx sdk.Context, txHash common.Hash) *evmtypes.Receipt {
	_ = "STUB: not implemented"
	return nil
}

func WaitForReceiptFromStore(t *testing.T, k *evmkeeper.Keeper, ctx sdk.Context, txHash common.Hash) *evmtypes.Receipt {
	_ = "STUB: not implemented"
	return nil
}
