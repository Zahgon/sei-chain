package teststaking

import (
	tmcrypto "github.com/sei-protocol/sei-chain/sei-tendermint/crypto"
	tmtypes "github.com/sei-protocol/sei-chain/sei-tendermint/types"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
)

// GetTmConsPubKey gets the validator's public key as a tmcrypto.PubKey.
func GetTmConsPubKey(v types.Validator) (tmcrypto.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(tmcrypto.PubKey), nil
}

// ToTmValidator casts an SDK validator to a tendermint type Validator.
func ToTmValidator(v types.Validator, r sdk.Int) (*tmtypes.Validator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ToTmValidators casts all validators to the corresponding tendermint type.
func ToTmValidators(v types.Validators, r sdk.Int) ([]*tmtypes.Validator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
