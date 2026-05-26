package commands

import (
	"github.com/spf13/cobra"
)

// MakeGenAutobahnConfigCommand creates a cobra command that generates an autobahn JSON config file.
// Each node directory must contain validator_pubkey.txt, node_pubkey.txt,
// autobahn_address.txt, and evmrpc_url.txt.
func MakeGenAutobahnConfigCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

//nolint:gosec // G304: dir comes from command args; filepath.Join already calls Clean

//nolint:gosec // G304: dir comes from command args; filepath.Join already calls Clean

//nolint:gosec // G304: dir comes from command args; filepath.Join already calls Clean

//nolint:gosec // G304: dir comes from command args; filepath.Join already calls Clean

// The flag defaults to "data/autobahn" so persistence is on without
// operator action. node/setup.go rootifies the relative path against
// cfg.RootDir at load time. Passing --persistent-state-dir= (empty)
// disables persistence and runs both consensus and data layers
// in-memory only.
