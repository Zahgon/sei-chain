package legacytm

import (
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
)

func ABCIToLegacyConsensusParams(params *tmproto.ConsensusParams) *abci.ConsensusParams {
	_ = "STUB: not implemented"
	return nil
}
