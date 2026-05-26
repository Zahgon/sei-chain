package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/query"
)

// Querier path constants
const (
	QueryBalance     = "balance"
	QueryAllBalances = "all_balances"
	QueryTotalSupply = "total_supply"
	QuerySupplyOf    = "supply_of"
)

// NewQueryBalanceRequest creates a new instance of QueryBalanceRequest.
func NewQueryBalanceRequest(addr sdk.AccAddress, denom string) *QueryBalanceRequest {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryAllBalancesRequest creates a new instance of QueryAllBalancesRequest.
func NewQueryAllBalancesRequest(addr sdk.AccAddress, req *query.PageRequest) *QueryAllBalancesRequest {
	_ = "STUB: not implemented"
	return nil
}

// NewQuerySpendableBalancesRequest creates a new instance of a
// QuerySpendableBalancesRequest.
func NewQuerySpendableBalancesRequest(addr sdk.AccAddress, req *query.PageRequest) *QuerySpendableBalancesRequest {
	_ = "STUB: not implemented"
	return nil
}

// QueryTotalSupplyParams defines the params for the following queries:
//
// - 'custom/bank/totalSupply'
type QueryTotalSupplyParams struct {
	Page, Limit int
}

// NewQueryTotalSupplyParams creates a new instance to query the total supply
func NewQueryTotalSupplyParams(page, limit int) QueryTotalSupplyParams {
	_ = "STUB: not implemented"
	return *new(QueryTotalSupplyParams)
}

// QuerySupplyOfParams defines the params for the following queries:
//
// - 'custom/bank/totalSupplyOf'
type QuerySupplyOfParams struct {
	Denom string
}

// NewQuerySupplyOfParams creates a new instance to query the total supply
// of a given denomination
func NewQuerySupplyOfParams(denom string) QuerySupplyOfParams {
	_ = "STUB: not implemented"
	return *new(QuerySupplyOfParams)
}
