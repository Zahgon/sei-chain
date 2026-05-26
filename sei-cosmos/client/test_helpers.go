package client

import (
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

var (
	_ AccountRetriever = TestAccountRetriever{}
	_ Account          = TestAccount{}
)

// TestAccount represents a client Account that can be used in unit tests
type TestAccount struct {
	Address sdk.AccAddress
	Num     uint64
	Seq     uint64
}

// GetAddress implements client Account.GetAddress
func (t TestAccount) GetAddress() sdk.AccAddress {
	_ = "STUB: not implemented"

	// GetPubKey implements client Account.GetPubKey
	return *new(sdk.AccAddress)
}

func (t TestAccount) GetPubKey() cryptotypes.PubKey {
	_ = "STUB: not implemented"

	// GetAccountNumber implements client Account.GetAccountNumber
	return *new(cryptotypes.PubKey)
}

func (t TestAccount) GetAccountNumber() uint64 {
	_ = "STUB: not implemented"

	// GetSequence implements client Account.GetSequence
	return 0
}

func (t TestAccount) GetSequence() uint64 {
	_ = "STUB: not implemented"

	// TestAccountRetriever is an AccountRetriever that can be used in unit tests
	return 0
}

type TestAccountRetriever struct {
	Accounts map[string]TestAccount
}

// GetAccount implements AccountRetriever.GetAccount
func (t TestAccountRetriever) GetAccount(_ Context, addr sdk.AccAddress) (Account, error) {
	_ = "STUB: not implemented"
	return *new(Account), nil
}

// GetAccountWithHeight implements AccountRetriever.GetAccountWithHeight
func (t TestAccountRetriever) GetAccountWithHeight(clientCtx Context, addr sdk.AccAddress) (Account, int64, error) {
	_ = "STUB: not implemented"
	return *new(Account), 0, nil
}

// EnsureExists implements AccountRetriever.EnsureExists
func (t TestAccountRetriever) EnsureExists(_ Context, addr sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

// GetAccountNumberSequence implements AccountRetriever.GetAccountNumberSequence
func (t TestAccountRetriever) GetAccountNumberSequence(_ Context, addr sdk.AccAddress) (accNum uint64, accSeq uint64, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}
