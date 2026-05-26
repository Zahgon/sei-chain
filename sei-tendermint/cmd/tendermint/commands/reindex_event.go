package commands

import (
	"github.com/spf13/cobra"

	tmcfg "github.com/sei-protocol/sei-chain/sei-tendermint/config"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/state"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/state/indexer"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/store"
)

const (
	reindexFailed = "event re-index failed: "
)

// MakeReindexEventCommand constructs a command to re-index events in a block height interval.
func MakeReindexEventCommand(conf *tmcfg.Config) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

func loadEventSinks(cfg *tmcfg.Config) ([]indexer.EventSink, error) {
	_ = "STUB: not implemented"
	// Check duplicated sinks.
	return nil, nil
}

func loadStateAndBlockStore(cfg *tmcfg.Config) (*store.BlockStore, state.Store, error) {
	_ = "STUB: not implemented"
	return nil, *new(state.Store), nil
}

type eventReIndexArgs struct {
	startHeight int64
	endHeight   int64
	sinks       []indexer.EventSink
	blockStore  state.BlockStore
	stateStore  state.Store
}

func eventReIndex(cmd *cobra.Command, args eventReIndexArgs) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // int to int64 is always safe; len() is non-negative

//nolint:gosec // j is bounded by len(b.Txs) which is bounded by block tx limits

type checkValidHeightArgs struct {
	startHeight int64
	endHeight   int64
}

func checkValidHeight(bs state.BlockStore, args checkValidHeightArgs) error {
	_ = "STUB: not implemented"
	return nil
}
