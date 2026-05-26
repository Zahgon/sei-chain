package backend

import (
	"github.com/sei-protocol/sei-chain/sei-db/config"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
)

func openPebbleDB(dbHome string, cfg config.StateStoreConfig) (types.StateStore, error) {
	_ = "STUB: not implemented"
	return *new(types.StateStore), nil
}
