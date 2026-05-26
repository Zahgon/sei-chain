package v100

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	genutiltypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/genutil/types"
	tmtypes "github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

// MigrateGenesis accepts exported v1.0.0 IBC client genesis file and migrates it to:
//
// - Update solo machine client state protobuf definition (v1 to v2)
// - Remove all solo machine consensus states
// - Remove all expired tendermint consensus states
func MigrateGenesis(appState genutiltypes.AppMap, clientCtx client.Context, genDoc tmtypes.GenesisDoc, maxExpectedTimePerBlock uint64) (genutiltypes.AppMap, error) {
	_ = "STUB: not implemented"
	return *new(genutiltypes.AppMap), nil
}

// ensure legacy solo machines are registered

// unmarshal relative source genesis application state

// #nosec G115 --- checked above

// set max expected time per block

// delete old genesis state

// set new ibc genesis state
