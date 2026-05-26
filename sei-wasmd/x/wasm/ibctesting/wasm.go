package ibctesting

import (

	//nolint
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
)

var wasmIdent = []byte("\x00\x61\x73\x6D")

// SeedNewContractInstance stores some wasm code and instantiates a new contract on this chain.
// This method can be called to prepare the store with some valid CodeInfo and ContractInfo. The returned
// Address is the contract address for this instance. Test should make use of this data and/or use NewIBCContractMockWasmer
// for using a contract mock in Go.
func (chain *TestChain) SeedNewContractInstance() sdk.AccAddress {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress)
}

func (chain *TestChain) StoreCodeFile(filename string) types.MsgStoreCodeResponse {
	_ = "STUB: not implemented"
	return *new(types.MsgStoreCodeResponse)
}

// compress for gas limit

func (chain *TestChain) StoreCode(byteCode []byte) types.MsgStoreCodeResponse {
	_ = "STUB: not implemented"
	return *new(types.MsgStoreCodeResponse)
}

// unmarshal protobuf response from data

func (chain *TestChain) InstantiateContract(codeID uint64, initMsg []byte) sdk.AccAddress {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress)
}

// SmartQuery This will serialize the query message and submit it to the contract.
// The response is parsed into the provided interface.
// Usage: SmartQuery(addr, QueryMsg{Foo: 1}, &response)
func (chain *TestChain) SmartQuery(contractAddr string, queryMsg interface{}, response interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: what is the query?

// unpack protobuf

// unpack json content

func (chain *TestChain) parseSDKResultData(r *sdk.Result) sdk.TxMsgData {
	_ = "STUB: not implemented"
	return *new(sdk.TxMsgData)
}

// ContractInfo is a helper function to returns the ContractInfo for the given contract address
func (chain *TestChain) ContractInfo(contractAddr sdk.AccAddress) *types.ContractInfo {
	_ = "STUB: not implemented"
	return nil
}
