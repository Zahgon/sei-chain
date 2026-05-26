package server

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/server/types"
	storetypes "github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
)

// GetPruningOptionsFromFlags parses command flags and returns the correct
// PruningOptions. If a pruning strategy is provided, that will be parsed and
// returned, otherwise, it is assumed custom pruning options are provided.
func GetPruningOptionsFromFlags(appOpts types.AppOptions) (storetypes.PruningOptions, error) {
	_ = "STUB: not implemented"
	return *new(storetypes.PruningOptions), nil
}
