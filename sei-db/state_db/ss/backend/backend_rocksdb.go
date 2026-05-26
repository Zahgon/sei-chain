//go:build rocksdbBackend

package backend

import (
	"github.com/sei-protocol/sei-chain/sei-db/config"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
)

func openRocksDB(dbHome string, cfg config.StateStoreConfig) (types.StateStore, error) {
	_ = "STUB: not implemented"
	// RocksDB's CreateIfMissing only creates the leaf directory
	return *new(types.StateStore), nil
}
