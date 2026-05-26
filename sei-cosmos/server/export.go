package server

// DONTCOVER

import (
	"time"

	tmbytes "github.com/sei-protocol/sei-chain/sei-tendermint/libs/bytes"
	tmtypes "github.com/sei-protocol/sei-chain/sei-tendermint/types"
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-cosmos/server/types"
)

const (
	FlagIsStreaming      = "streaming"
	FlagStreamingFile    = "streaming-file"
	FlagHeight           = "height"
	FlagForZeroHeight    = "for-zero-height"
	FlagJailAllowedAddrs = "jail-allowed-addrs"
)

type GenesisDocNoAppState struct {
	GenesisTime     time.Time                  `json:"genesis_time"`
	ChainID         string                     `json:"chain_id"`
	InitialHeight   int64                      `json:"initial_height,string"`
	ConsensusParams *tmtypes.ConsensusParams   `json:"consensus_params,omitempty"`
	Validators      []tmtypes.GenesisValidator `json:"validators,omitempty"`
	AppHash         tmbytes.HexBytes           `json:"app_hash"`
}

// ExportCmd dumps app state to JSON.
func ExportCmd(appExporter types.AppExporter, defaultNodeHome string) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: Tendermint uses a custom JSON decoder for GenesisDoc
// (except for stuff inside AppState). Inside AppState, we're free
// to encode as protobuf or amino.

// NOTE: Tendermint uses a custom JSON decoder for GenesisDoc
// (except for stuff inside AppState). Inside AppState, we're free
// to encode as protobuf or amino.
