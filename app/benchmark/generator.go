package benchmark

import (
	"context"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	evmtypes "github.com/sei-protocol/sei-chain/x/evm/types"
	"github.com/sei-protocol/sei-load/config"
	"github.com/sei-protocol/sei-load/generator"
	"github.com/sei-protocol/sei-load/generator/scenarios"
	loadtypes "github.com/sei-protocol/sei-load/types"
)

// Phase represents the current phase of the benchmark generator.
type Phase int

const (
	// PhaseWarmup is the initial phase to let the chain fully initialize.
	PhaseWarmup Phase = iota
	// PhaseSetup is the phase where contracts are deployed.
	PhaseSetup
	// PhaseLoad is the main phase where load transactions are generated.
	PhaseLoad
)

// WarmupBlocks is the number of blocks to wait before starting setup.
// This allows the chain to fully initialize (EVM genesis, fee collector address, etc.)
const WarmupBlocks = 3

// scenarioState tracks the state of a scenario instance.
type scenarioState struct {
	config       config.Scenario
	scenario     scenarios.TxGenerator
	accounts     loadtypes.AccountPool
	deployed     bool
	address      common.Address
	deployTx     *ethtypes.Transaction
	deployTxHash common.Hash
}

// Generator manages the benchmark transaction generation with setup and load phases.
type Generator struct {
	cfg      *config.LoadConfig
	txConfig client.TxConfig
	chainID  *big.Int

	scenarios      []*scenarioState
	deployer       *loadtypes.Account
	sharedAccounts loadtypes.AccountPool
	accountPools   []loadtypes.AccountPool

	phase          Phase
	warmupCounter  int // counts blocks during warmup phase
	pendingDeploys map[common.Hash]*scenarioState

	loadGenerator generator.Generator
	txsPerBatch   int

	mu sync.RWMutex
}

// NewGenerator creates a new benchmark generator from a config.
func NewGenerator(cfg *config.LoadConfig, txConfig client.TxConfig) (*Generator, error) {
	_ = "STUB: not implemented"
	// Read number of transactions per batch from environment variable, default to 1000
	return nil, nil
}

// Start with warmup to let chain initialize

// Create shared account pool

// Create scenario instances

// Determine account pool to use

// Scenario defines its own account settings - create separate pool

// Phase returns the current phase of the generator.
func (g *Generator) Phase() Phase { _ = "STUB: not implemented"; return *new(Phase) }

// IsSetupPhase returns true if the generator is in the setup phase.
func (g *Generator) IsSetupPhase() bool { _ = "STUB: not implemented"; return false }

// createDeploymentTx creates a deployment transaction for a scenario.
// Returns nil if the scenario doesn't need deployment.
func (g *Generator) createDeploymentTx(state *scenarioState) *ethtypes.Transaction {
	_ = "STUB: not implemented"
	// Create a temporary TxScenario for deployment
	return nil
}

// Not used for deployment

// Try to get a deployment transaction by checking if scenario supports it
// We use the scenario's Deploy method which creates deployment transactions
// For scenarios that don't need deployment (like EVMTransfer), this returns zero address

// Check if this is a contract scenario by trying to generate a deploy-style tx
// We do this by calling the scenario's internal deployment logic

// For now, we identify contract scenarios by name

// These don't need deployment

// For contract scenarios, create a deployment transaction
// We need to craft this manually since we're not using RPC

// craftDeploymentTx creates deployment bytecode transaction for contract scenarios.
func (g *Generator) craftDeploymentTx(state *scenarioState, txScenario *loadtypes.TxScenario) *ethtypes.Transaction {
	_ = "STUB: not implemented"
	return nil

	// Get deployment bytecode based on scenario type
}

// Create deployment transaction with chain ID set

// nil To address indicates contract creation

// 5M gas limit for deployment (increased)
// 1 gwei
// 1000 gwei (high to ensure inclusion)

// Sign the transaction (use CancunSigner to match sei-load scenarios)

// generateSetupBlock creates deployment transactions for undeployed scenarios.
// This is called on every PrepareProposal during setup phase, but we only
// create a deployment transaction ONCE per scenario.
func (g *Generator) generateSetupBlock() [][]byte { _ = "STUB: not implemented"; return nil }

// Skip if already deployed

// Skip if deployment transaction already created and pending
// (deployTxHash is set when we create the deploy tx)

// Already have a pending deploy tx for this scenario

// Create deployment transaction (only happens once per scenario)

// Scenario doesn't need deployment (e.g., EVMTransfer)

// Convert to Cosmos SDK tx

// Fast-path: if no scenarios need contract deployment (e.g., all EVMTransfer),
// we can transition to load phase immediately since all are marked deployed.
// For contract scenarios, transition happens in ProcessReceipts() after
// deployment transactions are confirmed.

// allScenariosDeployed returns true if all scenarios are marked as deployed.
func (g *Generator) allScenariosDeployed() bool { _ = "STUB: not implemented"; return false }

// generateLoadBlock generates load transactions.
func (g *Generator) generateLoadBlock() [][]byte { _ = "STUB: not implemented"; return nil }

// ethTxToTx converts an Ethereum transaction to an encoded Cosmos SDK tx.
func (g *Generator) ethTxToTx(ethTx *ethtypes.Transaction) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ProcessReceipts handles receipts from FinalizeBlock to extract deployed addresses.
func (g *Generator) ProcessReceipts(receipts map[common.Hash]*evmtypes.Receipt) {
	_ = "STUB: not implemented"
	return
}

// Attach the deployed address to the scenario

// Transition to load phase once all deployments are confirmed

// transitionToLoadPhase switches from setup to load generation.
func (g *Generator) transitionToLoadPhase() { _ = "STUB: not implemented"; return }

// Create weighted generator from deployed scenarios

// Generate returns the next batch of encoded txs.
func (g *Generator) Generate() [][]byte { _ = "STUB: not implemented"; return nil }

// Handle warmup phase - just count blocks and transition to setup

// Return empty during warmup

// GetPendingDeployHashes returns the transaction hashes of pending deployments.
func (g *Generator) GetPendingDeployHashes() []common.Hash { _ = "STUB: not implemented"; return nil }

// StartProposalChannel creates a channel that generates raw tx batches.
func (g *Generator) StartProposalChannel(ctx context.Context, logger *Logger) <-chan [][]byte {
	_ = "STUB: not implemented"
	return nil
}
