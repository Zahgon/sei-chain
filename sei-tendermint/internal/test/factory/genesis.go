package factory

import (
	"time"

	cfg "github.com/sei-protocol/sei-chain/sei-tendermint/config"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

func GenesisDoc(
	config *cfg.Config,
	time time.Time,
	validators []*types.Validator,
	consensusParams *types.ConsensusParams,
) *types.GenesisDoc {
	_ = "STUB: not implemented"
	return nil
}
