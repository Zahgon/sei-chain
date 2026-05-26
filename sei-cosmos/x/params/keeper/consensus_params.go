package keeper

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
)

// ConsensusParamsKeyTable returns an x/params module keyTable to be used in
// the BaseApp's ParamStore. The KeyTable registers the types along with the
// standard validation functions. Applications can choose to adopt this KeyTable
// or provider their own when the existing validation functions do not suite their
// needs.
func ConsensusParamsKeyTable() types.KeyTable {
	_ = "STUB: not implemented"
	return *new(types.KeyTable)
}
