package wasmbinding

import (
	"encoding/json"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const (
	OracleRoute       = "oracle"
	EpochRoute        = "epoch"
	TokenFactoryRoute = "tokenfactory"
	EVMRoute          = "evm"
	StakingExtRoute   = "stakingext"
)

type SeiQueryWrapper struct {
	// specifies which module handler should handle the query
	Route string `json:"route,omitempty"`
	// The query data that should be parsed into the module query
	QueryData json.RawMessage `json:"query_data,omitempty"`
}

func CustomQuerier(qp *QueryPlugin) func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil
}
