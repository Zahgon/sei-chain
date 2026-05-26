package cli

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
)

func buildCommissionRates(rateStr, maxRateStr, maxChangeRateStr string) (commission types.CommissionRates, err error) {
	_ = "STUB: not implemented"
	return *new(types.CommissionRates), nil
}
