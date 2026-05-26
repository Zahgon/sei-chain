package rpc

import (
	"context"
	"net/http"

	tmtypes "github.com/sei-protocol/sei-chain/sei-tendermint/types"
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// TODO these next two functions feel kinda hacky based on their placement

// ValidatorCommand returns the validator set for a given height
func ValidatorCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// optional height

// get the node

// Validator output
type ValidatorOutput struct {
	Address          sdk.ConsAddress    `json:"address"`
	PubKey           cryptotypes.PubKey `json:"pub_key"`
	ProposerPriority int64              `json:"proposer_priority"`
	VotingPower      int64              `json:"voting_power"`
}

// Validators at a certain height output in bech32 format
type ResultValidatorsOutput struct {
	BlockHeight int64             `json:"block_height"`
	Validators  []ValidatorOutput `json:"validators"`
	Total       uint64            `json:"total"`
}

func (rvo ResultValidatorsOutput) String() string { _ = "STUB: not implemented"; return "" }

func validatorOutput(validator *tmtypes.Validator) (ValidatorOutput, error) {
	_ = "STUB: not implemented"
	return *new(ValidatorOutput), nil
}

// GetValidators from client
func GetValidators(ctx context.Context, node client.Client, height *int64, page, limit *int) (ResultValidatorsOutput, error) {
	_ = "STUB: not implemented"
	return *new(ResultValidatorsOutput), nil
}

//nolint:gosec // total is guaranteed non-negative by the check above

// REST

// Validator Set at a height REST handler
func ValidatorSetRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// Latest Validator Set REST handler
func LatestValidatorSetRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}
