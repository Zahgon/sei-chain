package baseapp

import (
	"io"

	dbm "github.com/tendermint/tm-db"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/snapshots"
	"github.com/sei-protocol/sei-chain/sei-cosmos/store"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// File for storing in-package BaseApp optional functions,
// for options that need access to non-exported fields of the BaseApp

// SetPruning sets a pruning option on the multistore associated with the app
func SetPruning(opts sdk.PruningOptions) func(*BaseApp) { _ = "STUB: not implemented"; return nil }

// SetMinGasPrices returns an option that sets the minimum gas prices on the app.
func SetMinGasPrices(gasPricesStr string) func(*BaseApp) { _ = "STUB: not implemented"; return nil }

// SetHaltHeight returns a BaseApp option function that sets the halt block height.
func SetHaltHeight(blockHeight uint64) func(*BaseApp) { _ = "STUB: not implemented"; return nil }

// SetHaltTime returns a BaseApp option function that sets the halt block time.
func SetHaltTime(haltTime uint64) func(*BaseApp) { _ = "STUB: not implemented"; return nil }

// SetMinRetainBlocks returns a BaseApp option function that sets the minimum
// block retention height value when determining which heights to prune during
// ABCI Commit.
func SetMinRetainBlocks(minRetainBlocks uint64) func(*BaseApp) {
	_ = "STUB: not implemented"
	return nil
}

func SetCompactionInterval(compactionInterval uint64) func(*BaseApp) {
	_ = "STUB: not implemented"
	return nil
}

// SetTrace will turn on or off trace flag
func SetTrace(trace bool) func(*BaseApp) { _ = "STUB: not implemented"; return nil }

// SetIndexEvents provides a BaseApp option function that sets the events to index.
func SetIndexEvents(ie []string) func(*BaseApp) { _ = "STUB: not implemented"; return nil }

// SetInterBlockCache provides a BaseApp option function that sets the
// inter-block cache.
func SetInterBlockCache(cache sdk.MultiStorePersistentCache) func(*BaseApp) {
	_ = "STUB: not implemented"
	return nil
}

// SetSnapshotInterval sets the snapshot interval.
func SetSnapshotInterval(interval uint64) func(*BaseApp) { _ = "STUB: not implemented"; return nil }

func SetConcurrencyWorkers(workers int) func(*BaseApp) { _ = "STUB: not implemented"; return nil }

func SetOccEnabled(occEnabled bool) func(*BaseApp) { _ = "STUB: not implemented"; return nil }

// SetSnapshotKeepRecent sets the recent snapshots to keep.
func SetSnapshotKeepRecent(keepRecent uint32) func(*BaseApp) { _ = "STUB: not implemented"; return nil }

// SetSnapshotDirectory sets the snapshot directory.
func SetSnapshotDirectory(dir string) func(*BaseApp) { _ = "STUB: not implemented"; return nil }

// SetSnapshotStore sets the snapshot store.
func SetSnapshotStore(snapshotStore *snapshots.Store) func(*BaseApp) {
	_ = "STUB: not implemented"
	return nil
}

func (app *BaseApp) SetName(name string) { _ = "STUB: not implemented"; return }

// SetParamStore sets a parameter store on the BaseApp.
func (app *BaseApp) SetParamStore(ps ParamStore) { _ = "STUB: not implemented"; return }

// SetVersion sets the application's version string.
func (app *BaseApp) SetVersion(v string) { _ = "STUB: not implemented"; return }

// SetProtocolVersion sets the application's protocol version
func (app *BaseApp) SetProtocolVersion(v uint64) { _ = "STUB: not implemented"; return }

func (app *BaseApp) SetDB(db dbm.DB) { _ = "STUB: not implemented"; return }

func (app *BaseApp) SetCMS(cms store.CommitMultiStore) { _ = "STUB: not implemented"; return }

func (app *BaseApp) SetInitChainer(initChainer sdk.InitChainer) { _ = "STUB: not implemented"; return }

func (app *BaseApp) SetMidBlocker(midBlocker sdk.MidBlocker) { _ = "STUB: not implemented"; return }

func (app *BaseApp) SetEndBlocker(endBlocker sdk.EndBlocker) { _ = "STUB: not implemented"; return }

func (app *BaseApp) SetPreCommitHandler(preCommitHandler sdk.PreCommitHandler) {
	_ = "STUB: not implemented"
	return
}

func (app *BaseApp) SetCloseHandler(closeHandler sdk.CloseHandler) {
	_ = "STUB: not implemented"
	return
}

func (app *BaseApp) SetProcessProposalHandler(processProposalHandler sdk.ProcessProposalHandler) {
	_ = "STUB: not implemented"
	return
}

func (app *BaseApp) SetFinalizeBlocker(finalizeBlocker sdk.FinalizeBlocker) {
	_ = "STUB: not implemented"
	return
}

func (app *BaseApp) SetLoadVersionHandler(loadVersionHandler sdk.LoadVersionHandler) {
	_ = "STUB: not implemented"
	return
}

func (app *BaseApp) SetInplaceTestnetInitializer(inplaceTestnetInitializer sdk.InplaceTestnetInitializer) {
	_ = "STUB: not implemented"
	return
}

func (app *BaseApp) SetAnteHandler(ah sdk.AnteHandler) { _ = "STUB: not implemented"; return }

func (app *BaseApp) SetFauxMerkleMode() { _ = "STUB: not implemented"; return }

// SetCommitMultiStoreTracer sets the store tracer on the BaseApp's underlying
// CommitMultiStore.
func (app *BaseApp) SetCommitMultiStoreTracer(w io.Writer) { _ = "STUB: not implemented"; return }

// SetStoreLoader allows us to customize the rootMultiStore initialization.
func (app *BaseApp) SetStoreLoader(loader StoreLoader) { _ = "STUB: not implemented"; return }

// SetRouter allows us to customize the router.
func (app *BaseApp) SetRouter(router sdk.Router) { _ = "STUB: not implemented"; return }

// SetSnapshotStore sets the snapshot store.
func (app *BaseApp) SetSnapshotStore(snapshotStore *snapshots.Store) {
	_ = "STUB: not implemented"
	return
}

// SetSnapshotInterval sets the snapshot interval.
func (app *BaseApp) SetSnapshotInterval(snapshotInterval uint64) { _ = "STUB: not implemented"; return }

func (app *BaseApp) SetConcurrencyWorkers(workers int) { _ = "STUB: not implemented"; return }

func (app *BaseApp) SetOccEnabled(occEnabled bool) { _ = "STUB: not implemented"; return }

// SetSnapshotKeepRecent sets the number of recent snapshots to keep.
func (app *BaseApp) SetSnapshotKeepRecent(snapshotKeepRecent uint32) {
	_ = "STUB: not implemented"
	return
}

// SetSnapshotDirectory sets the snapshot directory.
func (app *BaseApp) SetSnapshotDirectory(dir string) { _ = "STUB: not implemented"; return }

// SetInterfaceRegistry sets the InterfaceRegistry.
func (app *BaseApp) SetInterfaceRegistry(registry types.InterfaceRegistry) {
	_ = "STUB: not implemented"
	return
}

// SetQueryMultiStore set a alternative MultiStore implementation to support online migration fallback read.
func (app *BaseApp) SetQueryMultiStore(ms sdk.CommitMultiStore) {
	_ = "STUB: not implemented"

	// SetMigrationHeight set the migration height for online migration so that query below this height will still be served from IAVL.
	return
}

func (app *BaseApp) SetMigrationHeight(height int64) { _ = "STUB: not implemented"; return }

// SetTxPrioritizer sets the transaction prioritizer for the BaseApp. If unset,
// calls to GetTxPriorityHint for all valid transactions will return 0.
func (app *BaseApp) SetTxPrioritizer(prioritizer sdk.TxPrioritizer) {
	_ = "STUB: not implemented"
	return
}
