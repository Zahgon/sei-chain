package cli

import (
	"github.com/spf13/cobra"
)

// GetCmdQueryConnections defines the command to query all the connection ends
// that this chain mantains.
func GetCmdQueryConnections() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryConnection defines the command to query a connection end
func GetCmdQueryConnection() *cobra.Command { _ = "STUB: not implemented"; return nil }

// #nosec G115 -- revision height is bounds checked above

// GetCmdQueryClientConnections defines the command to query a client connections
func GetCmdQueryClientConnections() *cobra.Command { _ = "STUB: not implemented"; return nil }

// #nosec G115 -- revision height is bounds checked above
