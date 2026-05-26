package cli

import (
	"context"

	"github.com/spf13/cobra"
)

// RunWithArgs executes the given command with the specified command line args
// and environmental variables set. It returns any error returned from cmd.Execute()
//
// This is only used in testing.
func RunWithArgs(ctx context.Context, cmd *cobra.Command, args []string, env map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// defer returns the environment back to normal

// set the args and env how we want them

// backup old value if there, to restore at end

// and finally run the command

func RunWithTrace(ctx context.Context, cmd *cobra.Command) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteConfigVals writes a toml file with the given values.
// It returns an error if writing was impossible.
func WriteConfigVals(dir string, vals map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// NewCompletionCmd returns a cobra.Command that generates bash and zsh
// completion scripts for the given root command. If hidden is true, the
// command will not show up in the root command's list of available commands.
func NewCompletionCmd(rootCmd *cobra.Command, hidden bool) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}
