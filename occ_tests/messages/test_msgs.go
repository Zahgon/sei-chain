package messages

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/sei-protocol/sei-chain/occ_tests/utils"
)

const instantiateMsg = `{"whitelist": ["sei1h9yjz89tl0dl6zu65dpxcqnxfhq60wxx8s5kag"],
    "use_whitelist":false,"admin":"sei1h9yjz89tl0dl6zu65dpxcqnxfhq60wxx8s5kag",
	"limit_order_fee":{"decimal":"0.0001","negative":false},
	"market_order_fee":{"decimal":"0.0001","negative":false},
	"liquidation_order_fee":{"decimal":"0.0001","negative":false},
	"margin_ratio":{"decimal":"0.0625","negative":false},
	"max_leverage":{"decimal":"4","negative":false},
	"default_base":"USDC",
	"native_token":"USDC","denoms": ["SEI","ATOM","USDC","SOL","ETH","OSMO","AVAX","BTC"],
	"full_denom_mapping": [["usei","SEI","0.000001"],["uatom","ATOM","0.000001"],["uusdc","USDC","0.000001"]],
	"funding_payment_lookback":3600,"spot_market_contract":"sei1h9yjz89tl0dl6zu65dpxcqnxfhq60wxx8s5kag",
	"supported_collateral_denoms": ["USDC"],
	"supported_multicollateral_denoms": ["ATOM"],
	"oracle_denom_mapping": [["usei","SEI","1"],["uatom","ATOM","1"],["uusdc","USDC","1"],["ueth","ETH","1"]],
	"multicollateral_whitelist": ["sei1h9yjz89tl0dl6zu65dpxcqnxfhq60wxx8s5kag"],
	"multicollateral_whitelist_enable": true,
	"funding_payment_pairs": [["USDC","ETH"]],
	"default_margin_ratios":{
		"initial":"0.3",
		"partial":"0.25",
		"maintenance":"0.06"
	}}`

func WasmInstantiate(tCtx *utils.TestContext, count int) []*utils.TestMessage {
	_ = "STUB: not implemented"
	return nil
}

// EVMTransferNonConflicting generates a list of EVM transfer messages that do not conflict with each other
// each message will have a brand new address
func EVMTransferNonConflicting(tCtx *utils.TestContext, count int) []*utils.TestMessage {
	_ = "STUB: not implemented"
	return nil
}

// EVMTransferConflicting generates a list of EVM transfer messages to the same address
func EVMTransferConflicting(tCtx *utils.TestContext, count int) []*utils.TestMessage {
	_ = "STUB: not implemented"
	return nil
}

func evmTransfer(testAcct utils.TestAcct, to common.Address, scenario string) *utils.TestMessage {
	_ = "STUB: not implemented"
	return nil
}

func BankTransfer(tCtx *utils.TestContext, count int) []*utils.TestMessage {
	_ = "STUB: not implemented"
	return nil
}

func GovernanceSubmitProposal(tCtx *utils.TestContext, count int) []*utils.TestMessage {
	_ = "STUB: not implemented"
	return nil
}

// ERC20toCWAssets generates messages that register EVM pointers to CW20 assets
// This creates ERC20 pointers to previously deployed CW20 tokens using an EVM transaction to the precompile
func ERC20toCWAssets(tCtx *utils.TestContext, count int) []*utils.TestMessage {
	_ = "STUB: not implemented"
	return nil
}

// Get the pointer precompile information

// Generate EVM transactions to register CW20 pointers

// Get the payload for calling the precompile's addCW20Pointer method

// Pack the method call with the CW20 contract address

// Create the EVM transaction

// Create and sign the transaction

// Create the MsgEVMTransaction
