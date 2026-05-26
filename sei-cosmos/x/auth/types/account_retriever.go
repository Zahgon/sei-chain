package types

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

var (
	_ client.Account          = AccountI(nil)
	_ client.AccountRetriever = AccountRetriever{}
)

// AccountRetriever defines the properties of a type that can be used to
// retrieve accounts.
type AccountRetriever struct{}

// GetAccount queries for an account given an address and a block height. An
// error is returned if the query or decoding fails.
func (ar AccountRetriever) GetAccount(clientCtx client.Context, addr sdk.AccAddress) (client.Account, error) {
	_ = "STUB: not implemented"
	return *new(client.Account), nil
}

// GetAccountWithHeight queries for an account given an address. Returns the
// height of the query with the account. An error is returned if the query
// or decoding fails.
func (ar AccountRetriever) GetAccountWithHeight(clientCtx client.Context, addr sdk.AccAddress) (client.Account, int64, error) {
	_ = "STUB: not implemented"
	return *new(client.Account), 0, nil
}

// EnsureExists returns an error if no account exists for the given address else nil.
func (ar AccountRetriever) EnsureExists(clientCtx client.Context, addr sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

// GetAccountNumberSequence returns sequence and account number for the given address.
// It returns an error if the account couldn't be retrieved from the state.
func (ar AccountRetriever) GetAccountNumberSequence(clientCtx client.Context, addr sdk.AccAddress) (uint64, uint64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}
