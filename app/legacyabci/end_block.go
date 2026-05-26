package legacyabci

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	crisiskeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/crisis/keeper"
	govkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/keeper"
	stakingkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/keeper"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	evmkeeper "github.com/sei-protocol/sei-chain/x/evm/keeper"
	oraclekeeper "github.com/sei-protocol/sei-chain/x/oracle/keeper"
)

type EndBlockKeepers struct {
	CrisisKeeper  *crisiskeeper.Keeper
	GovKeeper     *govkeeper.Keeper
	StakingKeeper *stakingkeeper.Keeper
	OracleKeeper  *oraclekeeper.Keeper
	EvmKeeper     *evmkeeper.Keeper
}

func EndBlock(ctx sdk.Context, height int64, blockGasUsed int64, keepers EndBlockKeepers) []abci.ValidatorUpdate {
	_ = "STUB: not implemented"
	return nil
}
