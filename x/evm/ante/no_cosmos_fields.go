package ante

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	txtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/tx"
)

// EVMNoCosmosFieldsDecorator ensures all Cosmos tx fields are empty for EVM txs.
type EVMNoCosmosFieldsDecorator struct{}

func NewEVMNoCosmosFieldsDecorator() EVMNoCosmosFieldsDecorator {
	_ = "STUB: not implemented"
	return *new(EVMNoCosmosFieldsDecorator)
}

func (d EVMNoCosmosFieldsDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}

type protoTxProvider interface {
	GetProtoTx() *txtypes.Tx
}

// ValidateNoCosmosTxFields rejects Cosmos wrapper fields that EVM txs must not use.
func ValidateNoCosmosTxFields(tx sdk.Tx) error { _ = "STUB: not implemented"; return nil }
