package sstest

import (
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
)

// Fills the db with multiple keys each with different versions
// TODO: Return just changeset so it can be altered after return
func FillData(db types.StateStore, numKeys int, versions int) error {
	_ = "STUB: not implemented"
	return nil
}

// Helper for creating the changeset and applying it to db
func DBApplyChangeset(db types.StateStore, version int64, storeKey string, key, val [][]byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Helper for creating the changeset and applying it to db
func DBApplyDeleteChangeset(db types.StateStore, version int64, storeKey string, key [][]byte) error {
	_ = "STUB: not implemented"
	return nil
}
