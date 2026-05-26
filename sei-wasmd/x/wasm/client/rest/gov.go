package rest

import (
	"encoding/json"
	"net/http"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/rest"
	govrest "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/client/rest"
	govtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"

	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
)

type StoreCodeProposalJSONReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

	Title       string    `json:"title" yaml:"title"`
	Description string    `json:"description" yaml:"description"`
	Proposer    string    `json:"proposer" yaml:"proposer"`
	Deposit     sdk.Coins `json:"deposit" yaml:"deposit"`

	RunAs string `json:"run_as" yaml:"run_as"`
	// WASMByteCode can be raw or gzip compressed
	WASMByteCode []byte `json:"wasm_byte_code" yaml:"wasm_byte_code"`
	// InstantiatePermission to apply on contract creation, optional
	InstantiatePermission *types.AccessConfig `json:"instantiate_permission" yaml:"instantiate_permission"`
}

func (s StoreCodeProposalJSONReq) Content() govtypes.Content {
	_ = "STUB: not implemented"
	return *new(govtypes.Content)
}

func (s StoreCodeProposalJSONReq) GetProposer() string { _ = "STUB: not implemented"; return "" }

func (s StoreCodeProposalJSONReq) GetDeposit() sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

func (s StoreCodeProposalJSONReq) GetBaseReq() rest.BaseReq {
	_ = "STUB: not implemented"
	return *new(rest.BaseReq)
}

func StoreCodeProposalHandler(cliCtx client.Context) govrest.ProposalRESTHandler {
	_ = "STUB: not implemented"
	return *new(govrest.ProposalRESTHandler)
}

type InstantiateProposalJSONReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`

	Proposer string    `json:"proposer" yaml:"proposer"`
	Deposit  sdk.Coins `json:"deposit" yaml:"deposit"`

	RunAs string `json:"run_as" yaml:"run_as"`
	// Admin is an optional address that can execute migrations
	Admin string          `json:"admin,omitempty" yaml:"admin"`
	Code  uint64          `json:"code_id" yaml:"code_id"`
	Label string          `json:"label" yaml:"label"`
	Msg   json.RawMessage `json:"msg" yaml:"msg"`
	Funds sdk.Coins       `json:"funds" yaml:"funds"`
}

func (s InstantiateProposalJSONReq) Content() govtypes.Content {
	_ = "STUB: not implemented"
	return *new(govtypes.Content)
}

func (s InstantiateProposalJSONReq) GetProposer() string { _ = "STUB: not implemented"; return "" }

func (s InstantiateProposalJSONReq) GetDeposit() sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

func (s InstantiateProposalJSONReq) GetBaseReq() rest.BaseReq {
	_ = "STUB: not implemented"
	return *new(rest.BaseReq)
}

func InstantiateProposalHandler(cliCtx client.Context) govrest.ProposalRESTHandler {
	_ = "STUB: not implemented"
	return *new(govrest.ProposalRESTHandler)
}

type MigrateProposalJSONReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`

	Proposer string    `json:"proposer" yaml:"proposer"`
	Deposit  sdk.Coins `json:"deposit" yaml:"deposit"`

	Contract string          `json:"contract" yaml:"contract"`
	Code     uint64          `json:"code_id" yaml:"code_id"`
	Msg      json.RawMessage `json:"msg" yaml:"msg"`
}

func (s MigrateProposalJSONReq) Content() govtypes.Content {
	_ = "STUB: not implemented"
	return *new(govtypes.Content)
}

func (s MigrateProposalJSONReq) GetProposer() string { _ = "STUB: not implemented"; return "" }

func (s MigrateProposalJSONReq) GetDeposit() sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

func (s MigrateProposalJSONReq) GetBaseReq() rest.BaseReq {
	_ = "STUB: not implemented"
	return *new(rest.BaseReq)
}

func MigrateProposalHandler(cliCtx client.Context) govrest.ProposalRESTHandler {
	_ = "STUB: not implemented"
	return *new(govrest.ProposalRESTHandler)
}

type ExecuteProposalJSONReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`

	Proposer string    `json:"proposer" yaml:"proposer"`
	Deposit  sdk.Coins `json:"deposit" yaml:"deposit"`

	Contract string          `json:"contract" yaml:"contract"`
	Msg      json.RawMessage `json:"msg" yaml:"msg"`
	// RunAs is the role that is passed to the contract's environment
	RunAs string    `json:"run_as" yaml:"run_as"`
	Funds sdk.Coins `json:"funds" yaml:"funds"`
}

func (s ExecuteProposalJSONReq) Content() govtypes.Content {
	_ = "STUB: not implemented"
	return *new(govtypes.Content)
}

func (s ExecuteProposalJSONReq) GetProposer() string { _ = "STUB: not implemented"; return "" }

func (s ExecuteProposalJSONReq) GetDeposit() sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

func (s ExecuteProposalJSONReq) GetBaseReq() rest.BaseReq {
	_ = "STUB: not implemented"
	return *new(rest.BaseReq)
}

func ExecuteProposalHandler(cliCtx client.Context) govrest.ProposalRESTHandler {
	_ = "STUB: not implemented"
	return *new(govrest.ProposalRESTHandler)
}

type SudoProposalJSONReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`

	Proposer string    `json:"proposer" yaml:"proposer"`
	Deposit  sdk.Coins `json:"deposit" yaml:"deposit"`

	Contract string          `json:"contract" yaml:"contract"`
	Msg      json.RawMessage `json:"msg" yaml:"msg"`
}

func (s SudoProposalJSONReq) Content() govtypes.Content {
	_ = "STUB: not implemented"
	return *new(govtypes.Content)
}

func (s SudoProposalJSONReq) GetProposer() string { _ = "STUB: not implemented"; return "" }

func (s SudoProposalJSONReq) GetDeposit() sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

func (s SudoProposalJSONReq) GetBaseReq() rest.BaseReq {
	_ = "STUB: not implemented"
	return *new(rest.BaseReq)
}

func SudoProposalHandler(cliCtx client.Context) govrest.ProposalRESTHandler {
	_ = "STUB: not implemented"
	return *new(govrest.ProposalRESTHandler)
}

type UpdateAdminJSONReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`

	Proposer string    `json:"proposer" yaml:"proposer"`
	Deposit  sdk.Coins `json:"deposit" yaml:"deposit"`

	NewAdmin string `json:"new_admin" yaml:"new_admin"`
	Contract string `json:"contract" yaml:"contract"`
}

func (s UpdateAdminJSONReq) Content() govtypes.Content {
	_ = "STUB: not implemented"
	return *new(govtypes.Content)
}

func (s UpdateAdminJSONReq) GetProposer() string { _ = "STUB: not implemented"; return "" }

func (s UpdateAdminJSONReq) GetDeposit() sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

func (s UpdateAdminJSONReq) GetBaseReq() rest.BaseReq {
	_ = "STUB: not implemented"
	return *new(rest.BaseReq)
}

func UpdateContractAdminProposalHandler(cliCtx client.Context) govrest.ProposalRESTHandler {
	_ = "STUB: not implemented"
	return *new(govrest.ProposalRESTHandler)
}

type ClearAdminJSONReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`

	Proposer string    `json:"proposer" yaml:"proposer"`
	Deposit  sdk.Coins `json:"deposit" yaml:"deposit"`

	Contract string `json:"contract" yaml:"contract"`
}

func (s ClearAdminJSONReq) Content() govtypes.Content {
	_ = "STUB: not implemented"
	return *new(govtypes.Content)
}

func (s ClearAdminJSONReq) GetProposer() string { _ = "STUB: not implemented"; return "" }

func (s ClearAdminJSONReq) GetDeposit() sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

func (s ClearAdminJSONReq) GetBaseReq() rest.BaseReq {
	_ = "STUB: not implemented"
	return *new(rest.BaseReq)
}

func ClearContractAdminProposalHandler(cliCtx client.Context) govrest.ProposalRESTHandler {
	_ = "STUB: not implemented"
	return *new(govrest.ProposalRESTHandler)
}

type PinCodeJSONReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`

	Proposer string    `json:"proposer" yaml:"proposer"`
	Deposit  sdk.Coins `json:"deposit" yaml:"deposit"`

	CodeIDs []uint64 `json:"code_ids" yaml:"code_ids"`
}

func (s PinCodeJSONReq) Content() govtypes.Content {
	_ = "STUB: not implemented"
	return *new(govtypes.Content)
}

func (s PinCodeJSONReq) GetProposer() string { _ = "STUB: not implemented"; return "" }

func (s PinCodeJSONReq) GetDeposit() sdk.Coins { _ = "STUB: not implemented"; return *new(sdk.Coins) }

func (s PinCodeJSONReq) GetBaseReq() rest.BaseReq {
	_ = "STUB: not implemented"
	return *new(rest.BaseReq)
}

func PinCodeProposalHandler(cliCtx client.Context) govrest.ProposalRESTHandler {
	_ = "STUB: not implemented"
	return *new(govrest.ProposalRESTHandler)
}

type UnpinCodeJSONReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`

	Proposer string    `json:"proposer" yaml:"proposer"`
	Deposit  sdk.Coins `json:"deposit" yaml:"deposit"`

	CodeIDs []uint64 `json:"code_ids" yaml:"code_ids"`
}

func (s UnpinCodeJSONReq) Content() govtypes.Content {
	_ = "STUB: not implemented"
	return *new(govtypes.Content)
}

func (s UnpinCodeJSONReq) GetProposer() string { _ = "STUB: not implemented"; return "" }

func (s UnpinCodeJSONReq) GetDeposit() sdk.Coins { _ = "STUB: not implemented"; return *new(sdk.Coins) }

func (s UnpinCodeJSONReq) GetBaseReq() rest.BaseReq {
	_ = "STUB: not implemented"
	return *new(rest.BaseReq)
}

func UnpinCodeProposalHandler(cliCtx client.Context) govrest.ProposalRESTHandler {
	_ = "STUB: not implemented"
	return *new(govrest.ProposalRESTHandler)
}

type UpdateInstantiateConfigProposalJSONReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

	Title               string                     `json:"title" yaml:"title"`
	Description         string                     `json:"description" yaml:"description"`
	Proposer            string                     `json:"proposer" yaml:"proposer"`
	Deposit             sdk.Coins                  `json:"deposit" yaml:"deposit"`
	AccessConfigUpdates []types.AccessConfigUpdate `json:"access_config_updates" yaml:"access_config_updates"`
}

func (s UpdateInstantiateConfigProposalJSONReq) Content() govtypes.Content {
	_ = "STUB: not implemented"
	return *new(govtypes.Content)
}

func (s UpdateInstantiateConfigProposalJSONReq) GetProposer() string {
	_ = "STUB: not implemented"
	return ""
}

func (s UpdateInstantiateConfigProposalJSONReq) GetDeposit() sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

func (s UpdateInstantiateConfigProposalJSONReq) GetBaseReq() rest.BaseReq {
	_ = "STUB: not implemented"
	return *new(rest.BaseReq)
}

func UpdateInstantiateConfigProposalHandler(cliCtx client.Context) govrest.ProposalRESTHandler {
	_ = "STUB: not implemented"
	return *new(govrest.ProposalRESTHandler)
}

type wasmProposalData interface {
	Content() govtypes.Content
	GetProposer() string
	GetDeposit() sdk.Coins
	GetBaseReq() rest.BaseReq
}

func toStdTxResponse(cliCtx client.Context, w http.ResponseWriter, data wasmProposalData) {
	_ = "STUB: not implemented"
	return
}
