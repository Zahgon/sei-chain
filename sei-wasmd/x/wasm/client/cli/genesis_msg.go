package cli

import (
	"encoding/json"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	tmtypes "github.com/sei-protocol/sei-chain/sei-tendermint/types"
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
)

// GenesisReader reads genesis data. Extension point for custom genesis state readers.
type GenesisReader interface {
	ReadWasmGenesis(cmd *cobra.Command) (*GenesisData, error)
}

// GenesisMutator extension point to modify the wasm module genesis state.
// This gives flexibility to customize the data structure in the genesis file a bit.
type GenesisMutator interface {
	// AlterWasmModuleState loads the genesis from the default or set home dir,
	// unmarshalls the wasm module section into the object representation
	// calls the callback function to modify it
	// and marshals the modified state back into the genesis file
	AlterWasmModuleState(cmd *cobra.Command, callback func(state *types.GenesisState, appState map[string]json.RawMessage) error) error
}

// GenesisStoreCodeCmd cli command to add a `MsgStoreCode` to the wasm section of the genesis
// that is executed on block 0.
func GenesisStoreCodeCmd(defaultNodeHome string, genesisMutator GenesisMutator) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

// GenesisInstantiateContractCmd cli command to add a `MsgInstantiateContract` to the wasm section of the genesis
// that is executed on block 0.
func GenesisInstantiateContractCmd(defaultNodeHome string, genesisMutator GenesisMutator) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

// simple sanity check that sender has some balance although it may be consumed by appState previous message already

//  does code id exists?

// permissions correct?

// GenesisExecuteContractCmd cli command to add a `MsgExecuteContract` to the wasm section of the genesis
// that is executed on block 0.
func GenesisExecuteContractCmd(defaultNodeHome string, genesisMutator GenesisMutator) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

// simple sanity check that sender has some balance although it may be consumed by appState previous message already

// - does contract address exists?

// GenesisListCodesCmd cli command to list all codes stored in the genesis wasm.code section
// as well as from messages that are queued in the wasm.genMsgs section.
func GenesisListCodesCmd(defaultNodeHome string, genReader GenesisReader) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

// GenesisListContractsCmd cli command to list all contracts stored in the genesis wasm.contract section
// as well as from messages that are queued in the wasm.genMsgs section.
func GenesisListContractsCmd(defaultNodeHome string, genReader GenesisReader) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

// clientCtx marshaller works only with proto or bytes so we marshal the output ourself
func printJSONOutput(cmd *cobra.Command, obj interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

type CodeMeta struct {
	CodeID uint64         `json:"code_id"`
	Info   types.CodeInfo `json:"info"`
}

func GetAllCodes(state *types.GenesisState) ([]CodeMeta, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// add inflight

// default

type ContractMeta struct {
	ContractAddress string             `json:"contract_address"`
	Info            types.ContractInfo `json:"info"`
}

func GetAllContracts(state *types.GenesisState) []ContractMeta {
	_ = "STUB: not implemented"
	return nil
}

// add inflight

func hasAccountBalance(cmd *cobra.Command, appState map[string]json.RawMessage, sender sdk.AccAddress, coins sdk.Coins) (bool, error) {
	_ = "STUB: not implemented"
	// no coins needed, no account needed
	return false, nil
}

func hasContract(state *types.GenesisState, contractAddr string) bool {
	_ = "STUB: not implemented"
	return false
}

// GenesisData contains raw and unmarshalled data from the genesis file
type GenesisData struct {
	GenesisFile     string
	GenDoc          *tmtypes.GenesisDoc
	AppState        map[string]json.RawMessage
	WasmModuleState *types.GenesisState
}

func NewGenesisData(genesisFile string, genDoc *tmtypes.GenesisDoc, appState map[string]json.RawMessage, wasmModuleState *types.GenesisState) *GenesisData {
	_ = "STUB: not implemented"
	return nil
}

type DefaultGenesisReader struct{}

func (d DefaultGenesisReader) ReadWasmGenesis(cmd *cobra.Command) (*GenesisData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var (
	_ GenesisReader  = DefaultGenesisIO{}
	_ GenesisMutator = DefaultGenesisIO{}
)

// DefaultGenesisIO implements both interfaces to read and modify the genesis state for this module.
// This implementation uses the default data structure that is used by the module.go genesis import/ export.
type DefaultGenesisIO struct {
	DefaultGenesisReader
}

// NewDefaultGenesisIO constructor to create a new instance
func NewDefaultGenesisIO() *DefaultGenesisIO { _ = "STUB: not implemented"; return nil }

// AlterWasmModuleState loads the genesis from the default or set home dir,
// unmarshalls the wasm module section into the object representation
// calls the callback function to modify it
// and marshals the modified state back into the genesis file
func (x DefaultGenesisIO) AlterWasmModuleState(cmd *cobra.Command, callback func(state *types.GenesisState, appState map[string]json.RawMessage) error) error {
	_ = "STUB: not implemented"
	return nil
}

// and store update

// contractSeqValue reads the contract sequence from the genesis or
// returns default start value used in the keeper
func contractSeqValue(state *types.GenesisState) uint64 { _ = "STUB: not implemented"; return 0 }

// codeSeqValue reads the code sequence from the genesis or
// returns default start value used in the keeper
func codeSeqValue(state *types.GenesisState) uint64 { _ = "STUB: not implemented"; return 0 }

// getActorAddress returns the account address for the `--run-as` flag.
// The flag value can either be an address already or a key name where the
// address is read from the keyring instead.
func getActorAddress(cmd *cobra.Command) (sdk.AccAddress, error) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), nil
}

// attempt to lookup address from Keybase if no address was provided
