package keeper

import (
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
)

// creates a querier for staking REST endpoints
func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	_ = "STUB: not implemented"
	return *new(sdk.Querier)
}

func queryValidators(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func queryValidator(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func queryValidatorDelegations(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func queryValidatorUnbondingDelegations(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func queryDelegatorDelegations(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func queryDelegatorUnbondingDelegations(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func queryDelegatorValidators(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func queryDelegatorValidator(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func queryDelegation(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func queryUnbondingDelegation(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func queryRedelegations(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func queryHistoricalInfo(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func queryPool(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func queryParameters(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// util

func DelegationToDelegationResponse(ctx sdk.Context, k Keeper, del types.Delegation) (types.DelegationResponse, error) {
	_ = "STUB: not implemented"
	return *new(types.DelegationResponse), nil
}

func DelegationsToDelegationResponses(
	ctx sdk.Context, k Keeper, delegations types.Delegations,
) (types.DelegationResponses, error) {
	_ = "STUB: not implemented"
	return *new(types.DelegationResponses), nil
}

func RedelegationsToRedelegationResponses(
	ctx sdk.Context, k Keeper, redels types.Redelegations,
) (types.RedelegationResponses, error) {
	_ = "STUB: not implemented"
	return *new(types.RedelegationResponses), nil
}
