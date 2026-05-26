package cli

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// Transaction flags for the x/distribution module
var (
	FlagCommission       = "commission"
	FlagMaxMessagesPerTx = "max-msgs"
)

const (
	MaxMessagesPerTxDefault = 0
)

// NewTxCmd returns a root CLI command handler for all x/distribution transaction commands.
func NewTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

type newGenerateOrBroadcastFunc func(context.Context, client.Context, *pflag.FlagSet, ...sdk.Msg) error

func NewSplitAndApply(
	ctx context.Context,
	genOrBroadcastFn newGenerateOrBroadcastFunc, clientCtx client.Context,
	fs *pflag.FlagSet, msgs []sdk.Msg, chunkSize int,
) error {
	_ = "STUB: not implemented"
	return nil
}

// split messages into slices of length chunkSize

func NewWithdrawRewardsCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func NewWithdrawAllRewardsCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// The transaction cannot be generated offline since it requires a query
// to get all the validators.

// build multi-message transaction

func NewSetWithdrawAddrCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func NewFundCommunityPoolCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdSubmitProposal implements the command to submit a community-pool-spend proposal
func GetCmdSubmitProposal() *cobra.Command { _ = "STUB: not implemented"; return nil }
