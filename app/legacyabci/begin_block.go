package legacyabci

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	capabilitykeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/keeper"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"

	distrkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/distribution/keeper"

	evidencekeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/evidence/keeper"
	slashingkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/slashing/keeper"

	stakingkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/keeper"
	upgradekeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/upgrade/keeper"

	ibckeeper "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/keeper"
	epochmodulekeeper "github.com/sei-protocol/sei-chain/x/epoch/keeper"
	evmkeeper "github.com/sei-protocol/sei-chain/x/evm/keeper"
)

type BeginBlockKeepers struct {
	EpochKeeper      *epochmodulekeeper.Keeper
	UpgradeKeeper    *upgradekeeper.Keeper
	CapabilityKeeper *capabilitykeeper.Keeper
	DistrKeeper      *distrkeeper.Keeper
	SlashingKeeper   *slashingkeeper.Keeper
	EvidenceKeeper   *evidencekeeper.Keeper
	StakingKeeper    *stakingkeeper.Keeper
	IBCKeeper        *ibckeeper.Keeper
	EvmKeeper        *evmkeeper.Keeper
}

func BeginBlock(
	ctx sdk.Context,
	height int64,
	votes []abci.VoteInfo,
	byzantineValidators []abci.Misbehavior,
	keepers BeginBlockKeepers,
) {
	_ = "STUB: not implemented"
	return
}

// TODO(PLT-343): remove once begin_blocker_duration verified

// TODO(PLT-343): remove once ibc_begin_blocker_duration verified
