package blocktest

import (
	servertypes "github.com/sei-protocol/sei-chain/sei-cosmos/server/types"
)

type Config struct {
	Enabled      bool   `mapstructure:"eth_blocktest_enabled"`
	TestDataPath string `mapstructure:"eth_blocktest_test_data_path"`
}

var DefaultConfig = Config{
	Enabled:      false,
	TestDataPath: "~/testdata/",
}

const (
	flagEnabled      = "eth_blocktest.eth_blocktest_enabled"
	flagTestDataPath = "eth_blocktest.eth_blocktest_test_data_path"
)

func ReadConfig(opts servertypes.AppOptions) (Config, error) {
	_ = "STUB: not implemented"
	// copy
	return *new(Config), nil
}
