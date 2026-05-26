package cli

import (
	"github.com/spf13/cobra"
)

// GetTxCmd returns a CLI command that has all the native evidence module tx
// commands mounted. In addition, it mounts all childCmds, implemented by outside
// modules, under a sub-command. This allows external modules to implement custom
// Evidence types and Handlers while having the ability to create and sign txs
// containing them all from a single root command.
func GetTxCmd(childCmds []*cobra.Command) *cobra.Command { _ = "STUB: not implemented"; return nil }

// TODO: Add tx commands.

// SubmitEvidenceCmd returns the top-level evidence submission command handler.
// All concrete evidence submission child command handlers should be registered
// under this command.
func SubmitEvidenceCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }
