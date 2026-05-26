package precompiles

import (
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ecommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/sei-protocol/sei-chain/precompiles/bank"
	"github.com/sei-protocol/sei-chain/precompiles/utils"
)

var SetupMtx = &sync.Mutex{}
var Initialized = false

type PrecompileInfo struct {
	ABI     abi.ABI
	Address ecommon.Address
}

// PrecompileNamesToInfo is Populated by InitializePrecompiles
var PrecompileNamesToInfo = map[string]PrecompileInfo{}

type IPrecompile interface {
	vm.PrecompiledContract
	GetABI() abi.ABI
	GetName() string
	Address() ecommon.Address
}

func GetCustomPrecompiles(
	latestUpgrade string,
	keepers utils.Keepers,
) map[ecommon.Address]utils.VersionedPrecompiles {
	_ = "STUB: not implemented"
	return nil
}

func InitializePrecompiles(
	dryRun bool,
	keepers utils.Keepers,
) error {
	_ = "STUB: not implemented"
	return nil
}

func GetPrecompileInfo(name string) PrecompileInfo {
	_ = "STUB: not implemented"

	// Precompile Info does not require any keeper state
	return *new(PrecompileInfo)
}

// This function modifies global variable in `vm` module. It should only be called once
// per precompile during initialization
func addPrecompileToVM(p IPrecompile) { _ = "STUB: not implemented"; return }

var PrecompileLastUpgrade = map[string]int64{
	bank.BankAddress: 1,
}
