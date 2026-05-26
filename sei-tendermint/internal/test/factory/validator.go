package factory

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

func Validator(ctx context.Context, votingPower int64) (*types.Validator, types.PrivValidator, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.PrivValidator), nil
}

func ValidatorSet(ctx context.Context, numValidators int, votingPower int64) (*types.ValidatorSet, []types.PrivValidator) {
	_ = "STUB: not implemented"
	return nil, nil
}
