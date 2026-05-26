package replay

import (
	servertypes "github.com/sei-protocol/sei-chain/sei-cosmos/server/types"
)

type Config struct {
	Enabled             bool   `mapstructure:"eth_replay_enabled"`
	EthRPC              string `mapstructure:"eth_rpc"`
	EthDataDir          string `mapstructure:"eth_data_dir"`
	ContractStateChecks bool   `mapstructure:"contract_state_checks"`
}

var DefaultConfig = Config{
	Enabled:             false,
	EthRPC:              "http://44.234.105.54:18545",
	EthDataDir:          "/root/.ethereum/chaindata",
	ContractStateChecks: false,
}

const (
	flagEnabled             = "eth_replay.eth_replay_enabled"
	flagEthRPC              = "eth_replay.eth_rpc"
	flagEthDataDir          = "eth_replay.eth_data_dir"
	flagContractStateChecks = "eth_replay.contract_state_checks"
)

func ReadConfig(opts servertypes.AppOptions) (Config, error) {
	_ = "STUB: not implemented"
	// copy
	return *new(Config), nil
}
