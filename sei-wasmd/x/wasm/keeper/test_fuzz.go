package keeper

import (
	fuzz "github.com/google/gofuzz"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
)

var ModelFuzzers = []interface{}{FuzzAddr, FuzzAddrString, FuzzAbsoluteTxPosition, FuzzContractInfo, FuzzStateModel, FuzzAccessType, FuzzAccessConfig, FuzzContractCodeHistory}

func FuzzAddr(m *sdk.AccAddress, c fuzz.Continue) { _ = "STUB: not implemented"; return }

func FuzzAddrString(m *string, c fuzz.Continue) { _ = "STUB: not implemented"; return }

func FuzzAbsoluteTxPosition(m *types.AbsoluteTxPosition, c fuzz.Continue) {
	_ = "STUB: not implemented"
	return
}

func FuzzContractInfo(m *types.ContractInfo, c fuzz.Continue) { _ = "STUB: not implemented"; return }

func FuzzContractCodeHistory(m *types.ContractCodeHistoryEntry, c fuzz.Continue) {
	_ = "STUB: not implemented"
	return
}

func FuzzStateModel(m *types.Model, c fuzz.Continue) { _ = "STUB: not implemented"; return }

func FuzzAccessType(m *types.AccessType, c fuzz.Continue) { _ = "STUB: not implemented"; return }

func FuzzAccessConfig(m *types.AccessConfig, c fuzz.Continue) { _ = "STUB: not implemented"; return }
