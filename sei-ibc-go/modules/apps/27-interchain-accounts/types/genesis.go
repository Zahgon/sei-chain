package types

import (
	controllertypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/27-interchain-accounts/controller/types"
	hosttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/27-interchain-accounts/host/types"
)

// DefaultGenesis creates and returns the interchain accounts GenesisState
func DefaultGenesis() *GenesisState { _ = "STUB: not implemented"; return nil }

// NewGenesisState creates and returns a new GenesisState instance from the provided controller and host genesis state types
func NewGenesisState(controllerGenesisState ControllerGenesisState, hostGenesisState HostGenesisState) *GenesisState {
	_ = "STUB: not implemented"
	return nil
}

// Validate performs basic validation of the interchain accounts GenesisState
func (gs GenesisState) Validate() error { _ = "STUB: not implemented"; return nil }

// DefaultControllerGenesis creates and returns the default interchain accounts ControllerGenesisState
func DefaultControllerGenesis() ControllerGenesisState {
	_ = "STUB: not implemented"
	return *new(ControllerGenesisState)
}

// NewControllerGenesisState creates a returns a new ControllerGenesisState instance
func NewControllerGenesisState(channels []ActiveChannel, accounts []RegisteredInterchainAccount, ports []string, controllerParams controllertypes.Params) ControllerGenesisState {
	_ = "STUB: not implemented"
	return *new(ControllerGenesisState)
}

// Validate performs basic validation of the ControllerGenesisState
func (gs ControllerGenesisState) Validate() error { _ = "STUB: not implemented"; return nil }

// DefaultHostGenesis creates and returns the default interchain accounts HostGenesisState
func DefaultHostGenesis() HostGenesisState {
	_ = "STUB: not implemented"
	return *new(HostGenesisState)
}

// NewHostGenesisState creates a returns a new HostGenesisState instance
func NewHostGenesisState(channels []ActiveChannel, accounts []RegisteredInterchainAccount, port string, hostParams hosttypes.Params) HostGenesisState {
	_ = "STUB: not implemented"
	return *new(HostGenesisState)
}

// Validate performs basic validation of the HostGenesisState
func (gs HostGenesisState) Validate() error { _ = "STUB: not implemented"; return nil }
