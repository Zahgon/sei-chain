package querier

import (
	servertypes "github.com/sei-protocol/sei-chain/sei-cosmos/server/types"
)

type Config struct {
	GasLimit uint64 `mapstructure:"evm_query_gas_limit"`
}

var DefaultConfig = Config{
	GasLimit: 300000,
}

const (
	flagGasLimit = "evm_query.evm_query_gas_limit"
)

func ReadConfig(opts servertypes.AppOptions) (Config, error) {
	_ = "STUB: not implemented"
	// copy
	return *new(Config), nil
}
