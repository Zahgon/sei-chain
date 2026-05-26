package cli

import (
	"github.com/spf13/cobra"
)

const (
	flagLatestHeight = "latest-height"
)

// GetCmdQueryClientStates defines the command to query all the light clients
// that this chain mantains.
func GetCmdQueryClientStates() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryClientState defines the command to query the state of a client with
// a given id as defined in https://github.com/cosmos/ibc/tree/master/spec/core/ics-002-client-semantics#query
func GetCmdQueryClientState() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryClientStatus defines the command to query the status of a client with a given id
func GetCmdQueryClientStatus() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryConsensusStates defines the command to query all the consensus states from a given
// client state.
func GetCmdQueryConsensusStates() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryConsensusStateHeights defines the command to query the heights of all client consensus states associated with the
// provided client ID.
func GetCmdQueryConsensusStateHeights() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryConsensusState defines the command to query the consensus state of
// the chain as defined in https://github.com/cosmos/ibc/tree/master/spec/core/ics-002-client-semantics#query
func GetCmdQueryConsensusState() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryHeader defines the command to query the latest header on the chain
func GetCmdQueryHeader() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdSelfConsensusState defines the command to query the self consensus state of a chain
func GetCmdSelfConsensusState() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdParams returns the command handler for ibc client parameter querying.
func GetCmdParams() *cobra.Command { _ = "STUB: not implemented"; return nil }
