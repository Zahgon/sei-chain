package cli

import (
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryFeeGrant returns cmd to query for a grant between granter and grantee.
func GetCmdQueryFeeGrant() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryFeeGrantsByGrantee returns cmd to query for all grants for a grantee.
func GetCmdQueryFeeGrantsByGrantee() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryFeeGrantsByGranter returns cmd to query for all grants by a granter.
func GetCmdQueryFeeGrantsByGranter() *cobra.Command { _ = "STUB: not implemented"; return nil }
