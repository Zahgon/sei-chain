package app

import (
	putils "github.com/sei-protocol/sei-chain/precompiles/utils"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
)

type PrecompileKeepers struct {
	putils.BankKeeper
	putils.BankMsgServer
	putils.EVMKeeper
	putils.AccountKeeper
	putils.OracleKeeper
	putils.WasmdKeeper
	putils.WasmdViewKeeper
	putils.StakingKeeper
	putils.StakingQuerier
	putils.GovKeeper
	putils.GovMsgServer
	putils.DistributionKeeper
	putils.TransferKeeper
	putils.ClientKeeper
	putils.ConnectionKeeper
	putils.ChannelKeeper
	txConf client.TxConfig
}

func NewPrecompileKeepers(a *App) *PrecompileKeepers { _ = "STUB: not implemented"; return nil }

func (pk *PrecompileKeepers) BankK() putils.BankKeeper {
	_ = "STUB: not implemented"
	return *new(putils.BankKeeper)
}
func (pk *PrecompileKeepers) BankMS() putils.BankMsgServer {
	_ = "STUB: not implemented"
	return *new(putils.BankMsgServer)
}
func (pk *PrecompileKeepers) EVMK() putils.EVMKeeper {
	_ = "STUB: not implemented"
	return *new(putils.EVMKeeper)
}
func (pk *PrecompileKeepers) AccountK() putils.AccountKeeper {
	_ = "STUB: not implemented"
	return *new(putils.AccountKeeper)
}
func (pk *PrecompileKeepers) OracleK() putils.OracleKeeper {
	_ = "STUB: not implemented"
	return *new(putils.OracleKeeper)
}
func (pk *PrecompileKeepers) WasmdK() putils.WasmdKeeper {
	_ = "STUB: not implemented"
	return *new(putils.WasmdKeeper)
}
func (pk *PrecompileKeepers) WasmdVK() putils.WasmdViewKeeper {
	_ = "STUB: not implemented"
	return *new(putils.WasmdViewKeeper)
}
func (pk *PrecompileKeepers) StakingK() putils.StakingKeeper {
	_ = "STUB: not implemented"
	return *new(putils.StakingKeeper)
}
func (pk *PrecompileKeepers) StakingQ() putils.StakingQuerier {
	_ = "STUB: not implemented"
	return *new(putils.StakingQuerier)
}
func (pk *PrecompileKeepers) GovK() putils.GovKeeper {
	_ = "STUB: not implemented"
	return *new(putils.GovKeeper)
}
func (pk *PrecompileKeepers) GovMS() putils.GovMsgServer {
	_ = "STUB: not implemented"
	return *new(putils.GovMsgServer)
}
func (pk *PrecompileKeepers) DistributionK() putils.DistributionKeeper {
	_ = "STUB: not implemented"
	return *new(putils.DistributionKeeper)
}
func (pk *PrecompileKeepers) TransferK() putils.TransferKeeper {
	_ = "STUB: not implemented"
	return *new(putils.TransferKeeper)
}
func (pk *PrecompileKeepers) ClientK() putils.ClientKeeper {
	_ = "STUB: not implemented"
	return *new(putils.ClientKeeper)
}
func (pk *PrecompileKeepers) ConnectionK() putils.ConnectionKeeper {
	_ = "STUB: not implemented"
	return *new(putils.ConnectionKeeper)
}
func (pk *PrecompileKeepers) ChannelK() putils.ChannelKeeper {
	_ = "STUB: not implemented"
	return *new(putils.ChannelKeeper)
}
func (pk *PrecompileKeepers) TxConfig() client.TxConfig {
	_ = "STUB: not implemented"
	return *new(client.TxConfig)
}
