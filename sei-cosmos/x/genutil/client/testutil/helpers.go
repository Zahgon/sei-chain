package testutil

import (
	tmcfg "github.com/sei-protocol/sei-chain/sei-tendermint/config"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
)

func ExecInitCmd(testMbm module.BasicManager, home string, cdc codec.Codec) error {
	_ = "STUB: not implemented"
	return nil
}

func CreateDefaultTendermintConfig(rootDir string) (*tmcfg.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
