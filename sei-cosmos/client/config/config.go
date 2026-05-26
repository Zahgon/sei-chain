package config

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
)

// Default constants
const (
	chainID        = ""
	keyringBackend = "os"
	output         = "text"
	node           = "tcp://localhost:26657"
	broadcastMode  = "sync"
)

type ClientConfig struct {
	ChainID        string `mapstructure:"chain-id" json:"chain-id"`
	KeyringBackend string `mapstructure:"keyring-backend" json:"keyring-backend"`
	Output         string `mapstructure:"output" json:"output"`
	Node           string `mapstructure:"node" json:"node"`
	BroadcastMode  string `mapstructure:"broadcast-mode" json:"broadcast-mode"`
}

// defaultClientConfig returns the reference to ClientConfig with default values.
func defaultClientConfig() *ClientConfig { _ = "STUB: not implemented"; return nil }

func (c *ClientConfig) SetChainID(chainID string) { _ = "STUB: not implemented"; return }

func (c *ClientConfig) SetKeyringBackend(keyringBackend string) { _ = "STUB: not implemented"; return }

func (c *ClientConfig) SetOutput(output string) { _ = "STUB: not implemented"; return }

func (c *ClientConfig) SetNode(node string) { _ = "STUB: not implemented"; return }

func (c *ClientConfig) SetBroadcastMode(broadcastMode string) { _ = "STUB: not implemented"; return }

// ReadFromClientConfig reads values from client.toml file and updates them in client Context
func ReadFromClientConfig(ctx client.Context) (client.Context, error) {
	_ = "STUB: not implemented"
	return *new(client.Context), nil
}

// if config.toml file does not exist we create it and write default ClientConfig values into it.

// we need to update KeyringDir field on Client Context first cause it is used in NewKeyringFromBackend

// https://github.com/cosmos/cosmos-sdk/issues/8986
