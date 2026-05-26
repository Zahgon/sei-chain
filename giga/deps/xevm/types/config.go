package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/params"
)

var CancunTime int64 = 0
var PragueTime int64 = 0

/*
*
XXBlock/Time fields indicate upgrade heights/timestamps. For example, a BerlinBlock
of 123 means the chain upgraded to the Berlin version at height 123; a ShanghaiTime
of 42198537129 means the chain upgraded to the Shanghai version at timestamp 42198537129.
A value of 0 means the upgrade is included in the genesis of the EVM, which will be the
case on Sei for all versions up to Cancun. Still, we want to keep these fields in the
config for backward compatibility with the official EVM lib.
*/
func (cc ChainConfig) EthereumConfig(chainID *big.Int) *params.ChainConfig {
	_ = "STUB: not implemented"
	return nil
}

// fork of Sei is supported outside EVM

func (cc ChainConfig) EthereumConfigWithSstore(chainID *big.Int, sstoreSetGasEIP2200 *uint64) *params.ChainConfig {
	_ = "STUB: not implemented"
	return nil
}

func DefaultChainConfig() ChainConfig { _ = "STUB: not implemented"; return *new(ChainConfig) }

func getUpgradeTimestamp(i int64) *uint64 { _ = "STUB: not implemented"; return nil }

func (cc ChainConfig) Validate() error { _ = "STUB: not implemented"; return nil }
