package testutil

import (
	"testing"

	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/keyring"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

type TestAccount struct {
	Name    string
	Address types.AccAddress
}

func CreateKeyringAccounts(t *testing.T, kr keyring.Keyring, num int) []TestAccount {
	_ = "STUB: not implemented"
	return nil
}
