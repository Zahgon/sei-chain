package simulation

import (
	"math/rand"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	simtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"

	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
)

func ParamChanges(r *rand.Rand, cdc codec.Codec) []simtypes.ParamChange {
	_ = "STUB: not implemented"
	return nil
}

func RandomParams(r *rand.Rand) types.Params { _ = "STUB: not implemented"; return *new(types.Params) }

// #nosec G115
