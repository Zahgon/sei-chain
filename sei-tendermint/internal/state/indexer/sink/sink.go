package sink

import (
	"github.com/sei-protocol/sei-chain/sei-tendermint/config"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/state/indexer"
)

// EventSinksFromConfig constructs a slice of indexer.EventSink using the provided
// configuration.
func EventSinksFromConfig(cfg *config.Config, dbProvider config.DBProvider, chainID string) ([]indexer.EventSink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check for duplicated sinks

// When we see null in the config, the eventsinks will be reset with the
// nullEventSink.
