package app

import (
	"context"
	"sync"

	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("tendermint", "test", "e2e", "app")

// Application is an ABCI application for use by end-to-end tests. It is a
// simple key/value store for strings, storing data in memory and persisting
// to disk as JSON, taking state sync snapshots if requested.
type Application struct {
	abci.BaseApplication
	mu              sync.Mutex
	state           *State
	snapshots       *SnapshotStore
	cfg             *Config
	restoreSnapshot *abci.Snapshot
	restoreChunks   [][]byte
}

// Config allows for the setting of high level parameters for running the e2e Application
// KeyType and ValidatorUpdates must be the same for all nodes running the same application.
type Config struct {
	// The directory with which state.json will be persisted in. Usually $HOME/.tendermint/data
	Dir string `toml:"dir"`

	// SnapshotInterval specifies the height interval at which the application
	// will take state sync snapshots. Defaults to 0 (disabled).
	SnapshotInterval uint64 `toml:"snapshot_interval"`

	// RetainBlocks specifies the number of recent blocks to retain. Defaults to
	// 0, which retains all blocks. Must be greater that PersistInterval,
	// SnapshotInterval and EvidenceAgeHeight.
	RetainBlocks uint64 `toml:"retain_blocks"`

	// KeyType sets the curve that will be used by validators.
	// Consensus keys are fixed to ed25519
	KeyType string `toml:"key_type"`

	// PersistInterval specifies the height interval at which the application
	// will persist state to disk. Defaults to 1 (every height), setting this to
	// 0 disables state persistence.
	PersistInterval uint64 `toml:"persist_interval"`

	// ValidatorUpdates is a map of heights to validator names and their power,
	// and will be returned by the ABCI application. For example, the following
	// changes the power of validator01 and validator02 at height 1000:
	//
	// [validator_update.1000]
	// validator01 = 20
	// validator02 = 10
	//
	// Specifying height 0 returns the validator update during InitChain. The
	// application returns the validator updates as-is, i.e. removing a
	// validator must be done by returning it with power 0, and any validators
	// not specified are not changed.
	//
	// height <-> pubkey <-> voting power
	ValidatorUpdates map[string]map[string]uint8 `toml:"validator_update"`

	// Add artificial delays to each of the main ABCI calls to mimic computation time
	// of the application
	PrepareProposalDelayMS uint64 `toml:"prepare_proposal_delay_ms"`
	ProcessProposalDelayMS uint64 `toml:"process_proposal_delay_ms"`
	CheckTxDelayMS         uint64 `toml:"check_tx_delay_ms"`
	FinalizeBlockDelayMS   uint64 `toml:"finalize_block_delay_ms"`
}

func DefaultConfig(dir string) *Config { _ = "STUB: not implemented"; return nil }

// NewApplication creates the application.
func NewApplication(cfg *Config) (*Application, error) { _ = "STUB: not implemented"; return nil, nil }

// Info implements ABCI.
func (app *Application) Info(_ context.Context, req *abci.RequestInfo) (*abci.ResponseInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // Height is a non-negative block height

// Info implements ABCI.
func (app *Application) InitChain(_ context.Context, req *abci.RequestInitChain) (*abci.ResponseInitChain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // InitialHeight is a validated non-negative value

// CheckTx implements ABCI.
func (app *Application) CheckTx(_ context.Context, req *abci.RequestCheckTxV2) *abci.ResponseCheckTxV2 {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // CheckTxDelayMS is a small test config value

// FinalizeBlock implements ABCI.
func (app *Application) FinalizeBlock(_ context.Context, req *abci.RequestFinalizeBlock) (*abci.ResponseFinalizeBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// shouldn't happen since we verified it in CheckTx

//nolint:gosec // Height is a non-negative block height

//nolint:gosec // FinalizeBlockDelayMS is a small config value

// Commit implements ABCI.
func (app *Application) Commit(_ context.Context) (*abci.ResponseCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // height > RetainBlocks when RetainBlocks > 0 in steady state

// Query implements ABCI.
func (app *Application) Query(_ context.Context, req *abci.RequestQuery) (*abci.ResponseQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // Height is a non-negative block height

// ListSnapshots implements ABCI.
func (app *Application) ListSnapshots(_ context.Context, req *abci.RequestListSnapshots) (*abci.ResponseListSnapshots, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LoadSnapshotChunk implements ABCI.
func (app *Application) LoadSnapshotChunk(_ context.Context, req *abci.RequestLoadSnapshotChunk) (*abci.ResponseLoadSnapshotChunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// OfferSnapshot implements ABCI.
func (app *Application) OfferSnapshot(_ context.Context, req *abci.RequestOfferSnapshot) (*abci.ResponseOfferSnapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ApplySnapshotChunk implements ABCI.
func (app *Application) ApplySnapshotChunk(_ context.Context, req *abci.RequestApplySnapshotChunk) (*abci.ResponseApplySnapshotChunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // Chunks is a small snapshot chunk count

// ProcessProposal implements part of the Application interface.
// It accepts any proposal that does not contain a malformed transaction.
func (app *Application) ProcessProposal(_ context.Context, req *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // ProcessProposalDelayMS is a small test config value

func (app *Application) CanRollback() bool { _ = "STUB: not implemented"; return false }

func (app *Application) Rollback() error { _ = "STUB: not implemented"; return nil }

// validatorUpdates generates a validator set update.
func (app *Application) validatorUpdates(height uint64) (abci.ValidatorUpdates, error) {
	_ = "STUB: not implemented"
	return *new(abci.ValidatorUpdates), nil
}

// the validator updates could be returned in arbitrary order,
// and that seems potentially bad. This orders the validator
// set.

// parseTx parses a tx in 'key=value' format into a key and value.
func parseTx(tx []byte) (string, string, error) { _ = "STUB: not implemented"; return "", "", nil }
