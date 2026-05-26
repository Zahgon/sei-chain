package keeper

import (
	"encoding/json"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
)

const (
	QueryListContractByCode = "list-contracts-by-code"
	QueryGetContract        = "contract-info"
	QueryGetContractState   = "contract-state"
	QueryGetCode            = "code"
	QueryListCode           = "list-code"
	QueryContractHistory    = "contract-history"
)

const (
	QueryMethodContractStateSmart = "smart"
	QueryMethodContractStateAll   = "all"
	QueryMethodContractStateRaw   = "raw"
)

// NewLegacyQuerier creates a new querier
func NewLegacyQuerier(keeper types.ViewKeeper, gasLimit sdk.Gas) sdk.Querier {
	_ = "STUB: not implemented"
	return *new(sdk.Querier)
}

func queryContractState(ctx sdk.Context, bech, queryMethod string, data []byte, gasLimit sdk.Gas, keeper types.ViewKeeper) (json.RawMessage, error) {
	_ = "STUB: not implemented"
	return *new(json.RawMessage), nil
}

// this returns a serialized json object (which internally encoded binary fields properly)

// this returns the raw data from the state, base64-encoded

// we enforce a subjective gas limit on all queries to avoid infinite loops

// this returns raw bytes (must be base64-encoded)

func queryCodeList(ctx sdk.Context, keeper types.ViewKeeper) ([]types.CodeInfoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func queryContractHistory(ctx sdk.Context, contractAddr sdk.AccAddress, keeper types.ViewKeeper) ([]types.ContractCodeHistoryEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// redact response

func queryContractListByCode(ctx sdk.Context, codeID uint64, keeper types.ViewKeeper) []string {
	_ = "STUB: not implemented"
	return nil
}
