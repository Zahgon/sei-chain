package privval

import (
	"context"

	privvalproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/privval"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

func DefaultValidationRequestHandler(
	ctx context.Context,
	privVal types.PrivValidator,
	req privvalproto.Message,
	chainID string,
) (privvalproto.Message, error) {
	_ = "STUB: not implemented"
	return *new(privvalproto.Message), nil
}
