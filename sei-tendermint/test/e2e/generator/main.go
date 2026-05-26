// nolint: gosec
package main

import (
	"context"
	stdlog "log"

	"github.com/sei-protocol/seilog"
	"github.com/spf13/cobra"
)

var logger = seilog.NewLogger("tendermint", "test", "e2e", "generator")

const (
	randomSeed int64 = 4827085738
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cli, err := NewCLI()
	if err != nil {
		stdlog.Fatal(err)
	}

	cli.Run(ctx)
}

// CLI is the Cobra-based command-line interface.
type CLI struct {
	root *cobra.Command
	opts Options
}

// NewCLI sets up the CLI.
func NewCLI() (*CLI, error) { _ = "STUB: not implemented"; return nil, nil }

// we'll output them ourselves in Run()

// generate generates manifests in a directory.
func (cli *CLI) generate() error { _ = "STUB: not implemented"; return nil }

// Run runs the CLI.
func (cli *CLI) Run(ctx context.Context) { _ = "STUB: not implemented"; return }
