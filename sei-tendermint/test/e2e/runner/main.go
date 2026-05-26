package main

import (
	"context"

	"github.com/spf13/cobra"

	e2e "github.com/sei-protocol/sei-chain/sei-tendermint/test/e2e/pkg"
)

const randomSeed = 2308084734268

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	NewCLI().Run(ctx)
}

// CLI is the Cobra-based command-line interface.
type CLI struct {
	root     *cobra.Command
	testnet  *e2e.Testnet
	preserve bool
}

// NewCLI sets up the CLI.
func NewCLI() *CLI { _ = "STUB: not implemented"; return nil }

// we'll output them ourselves in Run()

// nolint: gosec

// allow some txs to go through

// allow some txs to go through

// ensure chain progress

// to help make sure that we don't run into
// situations where 0 transactions have
// happened on quick cases, we make sure that
// it's been at least 10s before canceling the
// load generator.
//
// TODO allow the load generator to report
// successful transactions to avoid needing
// this sleep.

// wait for network to settle before tests

// nolint: gosec

// nolint: gosec

// nolint: gosec

// allow some txs to go through

// we benchmark performance over the next 100 blocks

// Run runs the CLI.
func (cli *CLI) Run(ctx context.Context) { _ = "STUB: not implemented"; return }
