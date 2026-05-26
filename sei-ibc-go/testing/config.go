package ibctesting

import (
	"time"

	connectiontypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/03-connection/types"
	channeltypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/04-channel/types"
	ibctmtypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/light-clients/07-tendermint/types"
)

type ClientConfig interface {
	GetClientType() string
}

type TendermintConfig struct {
	TrustLevel                   ibctmtypes.Fraction
	TrustingPeriod               time.Duration
	UnbondingPeriod              time.Duration
	MaxClockDrift                time.Duration
	AllowUpdateAfterExpiry       bool
	AllowUpdateAfterMisbehaviour bool
}

func NewTendermintConfig() *TendermintConfig { _ = "STUB: not implemented"; return nil }

func (tmcfg *TendermintConfig) GetClientType() string { _ = "STUB: not implemented"; return "" }

type ConnectionConfig struct {
	DelayPeriod uint64
	Version     *connectiontypes.Version
}

func NewConnectionConfig() *ConnectionConfig { _ = "STUB: not implemented"; return nil }

type ChannelConfig struct {
	PortID  string
	Version string
	Order   channeltypes.Order
}

func NewChannelConfig() *ChannelConfig { _ = "STUB: not implemented"; return nil }
