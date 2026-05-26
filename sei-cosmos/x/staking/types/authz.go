package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/authz"
)

// TODO: Revisit this once we have propoer gas fee framework.
// Tracking issues https://github.com/cosmos/cosmos-sdk/issues/9054, https://github.com/cosmos/cosmos-sdk/discussions/9072
const gasCostPerIteration = uint64(10)

// Normalized Msg type URLs
var (
	_ authz.Authorization = &StakeAuthorization{}
)

// NewStakeAuthorization creates a new StakeAuthorization object.
func NewStakeAuthorization(allowed []sdk.ValAddress, denied []sdk.ValAddress, authzType AuthorizationType, amount *sdk.Coin) (*StakeAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MsgTypeURL implements Authorization.MsgTypeURL.
func (a StakeAuthorization) MsgTypeURL() string { _ = "STUB: not implemented"; return "" }

func (a StakeAuthorization) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// Accept implements Authorization.Accept.
func (a StakeAuthorization) Accept(ctx sdk.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
	_ = "STUB: not implemented"
	return *new(authz.AcceptResponse), nil
}

func validateAndBech32fy(allowed []sdk.ValAddress, denied []sdk.ValAddress) ([]string, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func normalizeAuthzType(authzType AuthorizationType) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
