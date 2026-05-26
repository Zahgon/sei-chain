package utils

import (
	"encoding/json"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/rest"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types/proposal"
)

type (
	// ParamChangesJSON defines a slice of ParamChangeJSON objects which can be
	// converted to a slice of ParamChange objects.
	ParamChangesJSON []ParamChangeJSON

	// ParamChangeJSON defines a parameter change used in JSON input. This
	// allows values to be specified in raw JSON instead of being string encoded.
	ParamChangeJSON struct {
		Subspace string          `json:"subspace" yaml:"subspace"`
		Key      string          `json:"key" yaml:"key"`
		Value    json.RawMessage `json:"value" yaml:"value"`
	}

	// ParamChangeProposalJSON defines a ParameterChangeProposal with a deposit used
	// to parse parameter change proposals from a JSON file.
	ParamChangeProposalJSON struct {
		Title       string           `json:"title" yaml:"title"`
		Description string           `json:"description" yaml:"description"`
		IsExpedited bool             `json:"is_expedited" yaml:"is_expedited"`
		Changes     ParamChangesJSON `json:"changes" yaml:"changes"`
		Deposit     string           `json:"deposit" yaml:"deposit"`
	}

	// ParamChangeProposalReq defines a parameter change proposal request body.
	ParamChangeProposalReq struct {
		BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

		Title       string           `json:"title" yaml:"title"`
		Description string           `json:"description" yaml:"description"`
		IsExpedited bool             `json:"is_expedited" yaml:"is_expedited"`
		Changes     ParamChangesJSON `json:"changes" yaml:"changes"`
		Proposer    sdk.AccAddress   `json:"proposer" yaml:"proposer"`
		Deposit     sdk.Coins        `json:"deposit" yaml:"deposit"`
	}
)

func NewParamChangeJSON(subspace, key string, value json.RawMessage) ParamChangeJSON {
	_ = "STUB: not implemented"
	return *new(ParamChangeJSON)
}

// ToParamChange converts a ParamChangeJSON object to ParamChange.
func (pcj ParamChangeJSON) ToParamChange() proposal.ParamChange {
	_ = "STUB: not implemented"
	return *new(proposal.ParamChange)
}

// ToParamChanges converts a slice of ParamChangeJSON objects to a slice of
// ParamChange.
func (pcj ParamChangesJSON) ToParamChanges() []proposal.ParamChange {
	_ = "STUB: not implemented"
	return nil
}

// ParseParamChangeProposalJSON reads and parses a ParamChangeProposalJSON from
// file.
func ParseParamChangeProposalJSON(cdc *codec.LegacyAmino, proposalFile string) (ParamChangeProposalJSON, error) {
	_ = "STUB: not implemented"
	return *new(ParamChangeProposalJSON), nil
}
