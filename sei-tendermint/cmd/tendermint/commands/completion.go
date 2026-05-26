package commands

import (
	"github.com/spf13/cobra"
)

// NewCompletionCmd returns a cobra.Command that generates bash and zsh
// completion scripts for the given root command. If hidden is true, the
// command will not show up in the root command's list of available commands.
func NewCompletionCmd(rootCmd *cobra.Command, hidden bool) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}
