package keeper

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/sei-protocol/sei-chain/precompiles/bank"
	"github.com/sei-protocol/sei-chain/precompiles/gov"
	"github.com/sei-protocol/sei-chain/precompiles/staking"
	"github.com/sei-protocol/sei-chain/precompiles/wasmd"
)

// add any payable precompiles here
// these will suppress transfer events to/from the precompile address
var payablePrecompiles = map[string]struct{}{
	bank.BankAddress:       {},
	staking.StakingAddress: {},
	gov.GovAddress:         {},
	wasmd.WasmdAddress:     {},
}

func IsPayablePrecompile(addr *common.Address) bool { _ = "STUB: not implemented"; return false }
