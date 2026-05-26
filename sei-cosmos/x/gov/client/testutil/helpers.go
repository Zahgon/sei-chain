package testutil

import (
	"fmt"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client/flags"
	"github.com/sei-protocol/sei-chain/sei-cosmos/testutil"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

var commonArgs = []string{
	fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
	fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
	fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))).String()),
}

// MsgSubmitProposal creates a tx for submit proposal
func MsgSubmitProposal(clientCtx client.Context, from, title, description, proposalType string, extraArgs ...string) (testutil.BufferWriter, error) {
	_ = "STUB: not implemented"
	return *new(testutil.BufferWriter), nil
}

// MsgVote votes for a proposal
func MsgVote(clientCtx client.Context, from, id, vote string, extraArgs ...string) (testutil.BufferWriter, error) {
	_ = "STUB: not implemented"
	return *new(testutil.BufferWriter), nil
}

func MsgDeposit(clientCtx client.Context, from, id, deposit string, extraArgs ...string) (testutil.BufferWriter, error) {
	_ = "STUB: not implemented"
	return *new(testutil.BufferWriter), nil
}
