package testutil

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
)

// TestEncodingConfig defines an encoding configuration that is used for testing
// purposes. Note, MakeTestEncodingConfig takes a series of AppModuleBasic types
// which should only contain the relevant module being tested and any potential
// dependencies.
type TestEncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Codec             codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

func MakeTestEncodingConfig(modules ...module.AppModuleBasic) TestEncodingConfig {
	_ = "STUB: not implemented"
	return *new(TestEncodingConfig)
}

func MakeTestTxConfig() client.TxConfig { _ = "STUB: not implemented"; return *new(client.TxConfig) }

type TestBuilderTxConfig struct {
	client.TxConfig
	TxBuilder *TestTxBuilder
}

func MakeBuilderTestTxConfig() TestBuilderTxConfig {
	_ = "STUB: not implemented"
	return *new(TestBuilderTxConfig)
}

func (cfg TestBuilderTxConfig) NewTxBuilder() client.TxBuilder {
	_ = "STUB: not implemented"
	return *new(client.TxBuilder)
}

type TestTxBuilder struct {
	client.TxBuilder
	ExtOptions []*types.Any
}

func (b *TestTxBuilder) SetExtensionOptions(extOpts ...*types.Any) {
	_ = "STUB: not implemented"
	return
}
