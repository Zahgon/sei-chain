package simapp

import (
	"testing"

	"github.com/sei-protocol/sei-chain/sei-cosmos/baseapp"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	bankkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/keeper"
	capabilitykeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/keeper"
	distrkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/distribution/keeper"
	evidencekeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/evidence/keeper"
	slashingkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/slashing/keeper"
	stakingkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/keeper"
	ibctransferkeeper "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/transfer/keeper"
	ibckeeper "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/keeper"
)

type TestSupport struct {
	t   testing.TB
	app *SimApp
}

func NewTestSupport(t testing.TB, app *SimApp) *TestSupport { _ = "STUB: not implemented"; return nil }

func (s TestSupport) IBCKeeper() *ibckeeper.Keeper { _ = "STUB: not implemented"; return nil }

func (s TestSupport) AppCodec() codec.Codec { _ = "STUB: not implemented"; return *new(codec.Codec) }

func (s TestSupport) StakingKeeper() stakingkeeper.Keeper {
	_ = "STUB: not implemented"
	return *new(stakingkeeper.Keeper)
}

func (s TestSupport) BankKeeper() bankkeeper.Keeper {
	_ = "STUB: not implemented"
	return *new(bankkeeper.Keeper)
}

func (s TestSupport) TransferKeeper() ibctransferkeeper.Keeper {
	_ = "STUB: not implemented"
	return *new(ibctransferkeeper.Keeper)
}

func (s TestSupport) CapabilityKeeper() *capabilitykeeper.Keeper {
	_ = "STUB: not implemented"
	return nil
}

func (s TestSupport) DistrKeeper() *distrkeeper.Keeper { _ = "STUB: not implemented"; return nil }

func (s TestSupport) SlashingKeeper() *slashingkeeper.Keeper { _ = "STUB: not implemented"; return nil }

func (s TestSupport) EvidenceKeeper() *evidencekeeper.Keeper { _ = "STUB: not implemented"; return nil }

func (s TestSupport) GetBaseApp() *baseapp.BaseApp { _ = "STUB: not implemented"; return nil }

func (s TestSupport) GetTxConfig() client.TxConfig {
	_ = "STUB: not implemented"
	return *new(client.TxConfig)
}
