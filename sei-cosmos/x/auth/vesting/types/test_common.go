package types

import (
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/testutil/testdata"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// NewTestMsg generates a test message
func NewTestMsg(addrs ...sdk.AccAddress) *testdata.TestMsg { _ = "STUB: not implemented"; return nil }

// NewTestCoins coins to more than cover the fee
func NewTestCoins() sdk.Coins { _ = "STUB: not implemented"; return *new(sdk.Coins) }

// KeyTestPubAddr generates a test key pair
func KeyTestPubAddr() (cryptotypes.PrivKey, cryptotypes.PubKey, sdk.AccAddress) {
	_ = "STUB: not implemented"
	return *new(cryptotypes.PrivKey), *new(cryptotypes.PubKey), *new(sdk.AccAddress)
}
