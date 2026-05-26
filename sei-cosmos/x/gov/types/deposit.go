package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// NewDeposit creates a new Deposit instance
func NewDeposit(proposalID uint64, depositor sdk.AccAddress, amount sdk.Coins) Deposit {
	_ = "STUB: not implemented"
	return *new(Deposit)
}

func (d Deposit) String() string { _ = "STUB: not implemented"; return "" }

// Deposits is a collection of Deposit objects
type Deposits []Deposit

// Equal returns true if two slices (order-dependant) of deposits are equal.
func (d Deposits) Equal(other Deposits) bool { _ = "STUB: not implemented"; return false }

func (d Deposits) String() string { _ = "STUB: not implemented"; return "" }

// Empty returns whether a deposit is empty.
func (d Deposit) Empty() bool { _ = "STUB: not implemented"; return false }
