package query

import (
	txtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/tx"
)

// GetTxsEvent query the detailed transaction data, same as `seid q txs --events`
func GetTxsEvent(blockHeight int64) (*txtypes.GetTxsEventResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetTxByHash query the transaction by TX hash, same as `seid q tx --hash`
func GetTxByHash(txHash string) (*txtypes.GetTxResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
