package rootmulti

import (
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"

	"github.com/sei-protocol/sei-chain/sei-cosmos/store/dbadapter"
	"github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
)

var commithash = []byte("FAKE_HASH")

//----------------------------------------
// commitDBStoreWrapper should only be used for simulation/debugging,
// as it doesn't compute any commit hash, and it cannot load older state.

// Wrapper type for dbm.Db with implementation of KVStore
type commitDBStoreAdapter struct {
	dbadapter.Store
}

func (cdsa commitDBStoreAdapter) Commit(_ bool) types.CommitID {
	_ = "STUB: not implemented"
	return *new(types.CommitID)
}

func (cdsa commitDBStoreAdapter) LastCommitID() types.CommitID {
	_ = "STUB: not implemented"
	return *new(types.CommitID)
}

func (cdsa commitDBStoreAdapter) SetPruning(_ types.PruningOptions) {
	_ = "STUB: not implemented"

	// GetPruning is a no-op as pruning options cannot be directly set on this store.
	// They must be set on the root commit multi-store.
	return
}

func (cdsa commitDBStoreAdapter) GetPruning() types.PruningOptions {
	_ = "STUB: not implemented"
	return *new(types.PruningOptions)
}

func (cdsa commitDBStoreAdapter) Query(req abci.RequestQuery) abci.ResponseQuery {
	_ = "STUB: not implemented"
	return *new(abci.ResponseQuery)
}
