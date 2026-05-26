package cosmovisor

import (
	"bufio"
	"io"
	"os/exec"
	"sync"
)

// LaunchProcess runs a subprocess and returns when the subprocess exits,
// either when it dies, or *after* a successful upgrade.
func LaunchProcess(cfg *Config, args []string, stdout, stderr io.Writer) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// set scanner's buffer size to cfg.LogBufferSize, and ensure larger than bufio.MaxScanTokenSize otherwise fallback to bufio.MaxScanTokenSize

// three ways to exit - command ends, find regexp in scanOut, find regexp in scanErr

// WaitResult is used to wrap feedback on cmd state with some mutex logic.
// This is needed as multiple go-routines can affect this - two read pipes that can trigger upgrade
// As well as the command, which can fail
type WaitResult struct {
	// both err and info may be updated from several go-routines
	// access is wrapped by mutex and should only be done through methods
	err   error
	info  *UpgradeInfo
	mutex sync.Mutex
}

// AsResult reads the data protected by mutex to avoid race conditions
func (u *WaitResult) AsResult() (*UpgradeInfo, error) { _ = "STUB: not implemented"; return nil, nil }

// SetError will set with the first error using a mutex
// don't set it once info is set, that means we chose to kill the process
func (u *WaitResult) SetError(myErr error) { _ = "STUB: not implemented"; return }

// SetUpgrade sets first non-nil upgrade info, ensure error is then nil
// pass in a command to shutdown on successful upgrade
func (u *WaitResult) SetUpgrade(up *UpgradeInfo) { _ = "STUB: not implemented"; return }

// WaitForUpgradeOrExit listens to both output streams of the process, as well as the process state itself
// When it returns, the process is finished and all streams have closed.
//
// It returns (info, nil) if an upgrade should be initiated (and we killed the process)
// It returns (nil, err) if the process died by itself, or there was an issue reading the pipes
// It returns (nil, nil) if the process exited normally without triggering an upgrade. This is very unlikely
// to happened with "start" but may happened with short-lived commands like `gaiad export ...`
func WaitForUpgradeOrExit(cmd *exec.Cmd, scanOut, scanErr *bufio.Scanner) (*UpgradeInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// now we need to kill the process

// wait for the scanners, which can trigger upgrade and kill cmd

// if the command exits normally (eg. short command like `gaiad version`), just return (nil, nil)
// we often get broken read pipes if it runs too fast.
// if we had upgrade info, we would have killed it, and thus got a non-nil error code

// this will set the error code if it wasn't killed due to upgrade
