package simulation

import (
	"math/rand"

	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// Account contains a privkey, pubkey, address tuple
// eventually more useful data can be placed in here.
// (e.g. number of coins)
type Account struct {
	PrivKey cryptotypes.PrivKey
	PubKey  cryptotypes.PubKey
	Address sdk.AccAddress
	ConsKey cryptotypes.PrivKey
}

// Equals returns true if two accounts are equal
func (acc Account) Equals(acc2 Account) bool { _ = "STUB: not implemented"; return false }

// RandomAcc picks and returns a random account from an array and returs its
// position in the array.
func RandomAcc(r *rand.Rand, accs []Account) (Account, int) {
	_ = "STUB: not implemented"
	return *new(Account), 0
}

// RandomAccounts generates n random accounts
func RandomAccounts(r *rand.Rand, n int) []Account { _ = "STUB: not implemented"; return nil }

// don't need that much entropy for simulation

// FindAccount iterates over all the simulation accounts to find the one that matches
// the given address
func FindAccount(accs []Account, address sdk.Address) (Account, bool) {
	_ = "STUB: not implemented"
	return *new(Account), false
}

// RandomFees returns a random fee by selecting a random coin denomination and
// amount from the account's available balance. If the user doesn't have enough
// funds for paying fees, it returns empty coins.
func RandomFees(r *rand.Rand, ctx sdk.Context, spendableCoins sdk.Coins) (sdk.Coins, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Coins), nil
}

// Create a random fee and verify the fees are within the account's spendable
// balance.
