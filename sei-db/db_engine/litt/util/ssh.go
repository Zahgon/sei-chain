package util

import (
	"log/slog"

	"golang.org/x/crypto/ssh"
)

// SSHSession encapsulates an SSH session with a remote host.
type SSHSession struct {
	logger         *slog.Logger
	client         *ssh.Client
	user           string
	host           string
	port           uint64
	keyPath        string
	knownHostsPath string
	verbose        bool
}

// Create a new SSH session to a remote host.
//
// If the knownHosts parameter is provided, it will be used to verify the host's key. If it is absent or empty,
// the host key verification will be skipped.
func NewSSHSession(
	logger *slog.Logger,
	user string,
	host string,
	port uint64,
	keyPath string,
	knownHosts string,
	verbose bool,
) (*SSHSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // overridden when knownHosts provided

//nolint:gosec // caller-supplied key path

// Close the SSH session.
func (s *SSHSession) Close() error { _ = "STUB: not implemented"; return nil }

// Search for all files matching a regex inside a file tree at the specified root path.
func (s *SSHSession) FindFiles(root string, extensions []string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// There are no files since the directory does not exist.

// Skip empty lines

// Stop checking other extensions once a match is found

// Mkdirs creates the specified directory on the remote machine, including any necessary parent directories.
func (s *SSHSession) Mkdirs(path string) error { _ = "STUB: not implemented"; return nil }

// Directory already exists, no error needed

// Rsync transfers files from the local machine to the remote machine using rsync. The throttle is ignored
// if less than or equal to 0.
func (s *SSHSession) Rsync(sourceFile string, destFile string, throttleMB float64) error {
	_ = "STUB: not implemented"
	return nil
}

// If the source file is a symlink, we actually want to send the thing the symlink points to.

// Resolve the symlink to get the actual file it points to

// rsync interprets --bwlimit in KB/s, so we convert MB to KB

//nolint:gosec // arguments built from caller-trusted config

// Exec executes a command on the remote machine and returns the output. Returns the result of stdout and stderr.
func (s *SSHSession) Exec(command string) (stdout string, stderr string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}
