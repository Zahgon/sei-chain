package debug

import (
	"github.com/spf13/cobra"
)

func getKillCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Create a temporary directory which will contain all the state dumps and
// relevant files and directories that will be compressed into a file.

// killProc attempts to kill the Tendermint process with a given PID with an
// ABORT signal which should result in a goroutine stacktrace. The PID's STDERR
// is tailed and piped to a file under the directory dir. An error is returned
// if the output file cannot be created or the tail command cannot be started.
// An error is not returned if any subsequent syscall fails.
func killProc(pid int, dir string) error {
	_ = "STUB: not implemented"
	// pipe STDERR output from tailing the Tendermint process to a file
	//
	// NOTE: This will only work on UNIX systems.
	return nil
}

// nolint: gosec

// kill the underlying Tendermint process and subsequent tailing process

// Killing the Tendermint process with the '-ABRT|-6' signal will result in
// a goroutine stacktrace.

// allow some time to allow the Tendermint process to be killed
//
// TODO: We should 'wait' for a kill to succeed (e.g. poll for PID until it
// cannot be found). Regardless, this should be ample time.

// only return an error not invoked by a manual kill
