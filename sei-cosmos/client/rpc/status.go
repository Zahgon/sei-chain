package rpc

import (
	"net/http"

	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/bytes"
	ctypes "github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/types"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/version"
)

// ValidatorInfo is info about the node's validator, same as Tendermint,
// except that we use our own PubKey.
type validatorInfo struct {
	Address     bytes.HexBytes
	PubKey      cryptotypes.PubKey
	VotingPower int64
}

// ResultStatus is node's info, same as Tendermint, except that we use our own
// PubKey.
type resultStatus struct {
	NodeInfo      tmproto.NodeInfo
	SyncInfo      ctypes.SyncInfo
	ValidatorInfo validatorInfo
}

// StatusCommand returns the command to return the status of the network.
func StatusCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// `status` has TM pubkeys, we need to convert them to our pubkeys.

// NodeInfoResponse defines a response type that contains node status and version
// information.
type NodeInfoResponse struct {
	tmproto.NodeInfo `json:"node_info"`

	ApplicationVersion version.Info `json:"application_version"`
}

// REST handler for node info
func NodeInfoRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// SyncingResponse defines a response type that contains node syncing information.
type SyncingResponse struct {
	Syncing bool `json:"syncing"`
}

// REST handler for node syncing
func NodeSyncingRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}
