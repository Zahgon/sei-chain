// Package benchmark provides transaction generation capabilities for benchmarking.
//
// The benchmark system operates in two phases:
//
//  1. Setup Phase: Deploys any contracts required by the configured scenarios.
//     During this phase, deployment transactions are generated and submitted.
//     After each block, receipts are checked to extract deployed contract addresses.
//
//  2. Load Phase: Once all contracts are deployed, the system transitions to
//     generating load transactions according to the configured scenario weights.
//
// Usage:
//
//	cfg, _ := benchmark.LoadConfig(configPath, evmChainID, seiChainID)
//	gen, _ := benchmark.NewGenerator(cfg, txConfig)
//	benchLogger := benchmark.NewLogger(logger)
//	proposalCh := gen.StartProposalChannel(ctx, benchLogger)
//
// The generator can be configured via JSON config files that follow the sei-load
// LoadConfig format. See benchmark/scenarios/ for example configurations.
package benchmark

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	evmtypes "github.com/sei-protocol/sei-chain/x/evm/types"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("app", "benchmark")

// Manager coordinates benchmark generation and logging.
type Manager struct {
	Generator  *Generator
	Logger     *Logger
	proposalCh <-chan [][]byte
}

// NewManager creates a new benchmark manager from configuration.
func NewManager(ctx context.Context, txConfig client.TxConfig, chainID string, evmChainID int64) (*Manager, error) {
	_ = "STUB: not implemented"
	// Defensive check: prevent benchmarking on live chains
	return nil, nil
}

// Load config from environment variable or use default

// ProposalChannel returns the channel of prepared proposals.
func (m *Manager) ProposalChannel() <-chan [][]byte { _ = "STUB: not implemented"; return nil }

// ProcessReceipts forwards receipts to the generator for deployment tracking.
func (m *Manager) ProcessReceipts(receipts map[common.Hash]*evmtypes.Receipt) {
	_ = "STUB: not implemented"
	return
}

// IsSetupPhase returns true if the benchmark is still in the setup phase.
func (m *Manager) IsSetupPhase() bool { _ = "STUB: not implemented"; return false }

// GetPendingDeployHashes returns the hashes of pending deployment transactions.
func (m *Manager) GetPendingDeployHashes() []common.Hash { _ = "STUB: not implemented"; return nil }
