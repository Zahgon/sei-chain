package types

import (
	"time"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const (
	// ModuleName is the name of the staking module
	ModuleName = "staking"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the staking module
	QuerierRoute = ModuleName

	// RouterKey is the msg router key for the staking module
	RouterKey = ModuleName
)

var (
	// Keys for store prefixes
	// Last* values are constant during a block.
	LastValidatorPowerKey = []byte{0x11} // prefix for each key to a validator index, for bonded validators
	LastTotalPowerKey     = []byte{0x12} // prefix for the total power

	ValidatorsKey             = []byte{0x21} // prefix for each key to a validator
	ValidatorsByConsAddrKey   = []byte{0x22} // prefix for each key to a validator index, by pubkey
	ValidatorsByPowerIndexKey = []byte{0x23} // prefix for each key to a validator index, sorted by power

	DelegationKey                    = []byte{0x31} // key for a delegation
	UnbondingDelegationKey           = []byte{0x32} // key for an unbonding-delegation
	UnbondingDelegationByValIndexKey = []byte{0x33} // prefix for each key for an unbonding-delegation, by validator operator
	RedelegationKey                  = []byte{0x34} // key for a redelegation
	RedelegationByValSrcIndexKey     = []byte{0x35} // prefix for each key for an redelegation, by source validator operator
	RedelegationByValDstIndexKey     = []byte{0x36} // prefix for each key for an redelegation, by destination validator operator

	UnbondingQueueKey    = []byte{0x41} // prefix for the timestamps in unbonding queue
	RedelegationQueueKey = []byte{0x42} // prefix for the timestamps in redelegations queue
	ValidatorQueueKey    = []byte{0x43} // prefix for the timestamps in validator queue

	HistoricalInfoKey = []byte{0x50} // prefix for the historical info
)

// GetValidatorKey creates the key for the validator with address
// VALUE: staking/Validator
func GetValidatorKey(operatorAddr sdk.ValAddress) []byte { _ = "STUB: not implemented"; return nil }

// GetValidatorByConsAddrKey creates the key for the validator with pubkey
// VALUE: validator operator address ([]byte)
func GetValidatorByConsAddrKey(addr sdk.ConsAddress) []byte { _ = "STUB: not implemented"; return nil }

// AddressFromValidatorsKey creates the validator operator address from ValidatorsKey
func AddressFromValidatorsKey(key []byte) []byte { _ = "STUB: not implemented"; return nil }

// remove prefix bytes and address length

// AddressFromLastValidatorPowerKey creates the validator operator address from LastValidatorPowerKey
func AddressFromLastValidatorPowerKey(key []byte) []byte { _ = "STUB: not implemented"; return nil }

// remove prefix bytes and address length

// GetValidatorsByPowerIndexKey creates the validator by power index.
// Power index is the key used in the power-store, and represents the relative
// power ranking of the validator.
// VALUE: validator operator address ([]byte)
func GetValidatorsByPowerIndexKey(validator Validator, powerReduction sdk.Int) []byte {
	_ = "STUB: not implemented"
	// NOTE the address doesn't need to be stored because counter bytes must always be different
	// NOTE the larger values are of higher value
	return nil
}

//nolint:gosec // bounds checked above

// 8

// key is of format prefix || powerbytes || addrLen (1byte) || addrBytes

//nolint:gosec // address length is always <= 255 bytes per address.MaxAddrLen

// GetLastValidatorPowerKey creates the bonded validator index key for an operator address
func GetLastValidatorPowerKey(operator sdk.ValAddress) []byte {
	_ = "STUB: not implemented"
	return nil
}

// ParseValidatorPowerRankKey parses the validators operator address from power rank key
func ParseValidatorPowerRankKey(key []byte) (operAddr []byte) {
	_ = "STUB: not implemented"

	// key is of format prefix (1 byte) || powerbytes || addrLen (1byte) || addrBytes
	return nil
}

// GetValidatorQueueKey returns the prefix key used for getting a set of unbonding
// validators whose unbonding completion occurs at the given time and height.
func GetValidatorQueueKey(timestamp time.Time, height int64) []byte {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // bounds checked above

// copy the prefix

// copy the encoded time bytes length
//nolint:gosec // len() is always non-negative

// copy the encoded time bytes

// copy the encoded height

// ParseValidatorQueueKey returns the encoded time and height from a key created
// from GetValidatorQueueKey.
func ParseValidatorQueueKey(bz []byte) (time.Time, int64, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), 0, nil
}

//nolint:gosec // remaining is validated non-negative

//nolint:gosec // timeBzL bounds checked above

//nolint:gosec // timeBzL bounds checked above

//nolint:gosec // height from trusted store data

// GetDelegationKey creates the key for delegator bond with validator
// VALUE: staking/Delegation
func GetDelegationKey(delAddr sdk.AccAddress, valAddr sdk.ValAddress) []byte {
	_ = "STUB: not implemented"
	return nil
}

// GetDelegationsKey creates the prefix for a delegator for all validators
func GetDelegationsKey(delAddr sdk.AccAddress) []byte { _ = "STUB: not implemented"; return nil }

// GetUBDKey creates the key for an unbonding delegation by delegator and validator addr
// VALUE: staking/UnbondingDelegation
func GetUBDKey(delAddr sdk.AccAddress, valAddr sdk.ValAddress) []byte {
	_ = "STUB: not implemented"
	return nil
}

// GetUBDByValIndexKey creates the index-key for an unbonding delegation, stored by validator-index
// VALUE: none (key rearrangement used)
func GetUBDByValIndexKey(delAddr sdk.AccAddress, valAddr sdk.ValAddress) []byte {
	_ = "STUB: not implemented"
	return nil
}

// GetUBDKeyFromValIndexKey rearranges the ValIndexKey to get the UBDKey
func GetUBDKeyFromValIndexKey(indexKey []byte) []byte { _ = "STUB: not implemented"; return nil }

// remove prefix bytes

// GetUBDsKey creates the prefix for all unbonding delegations from a delegator
func GetUBDsKey(delAddr sdk.AccAddress) []byte { _ = "STUB: not implemented"; return nil }

// GetUBDsByValIndexKey creates the prefix keyspace for the indexes of unbonding delegations for a validator
func GetUBDsByValIndexKey(valAddr sdk.ValAddress) []byte { _ = "STUB: not implemented"; return nil }

// GetUnbondingDelegationTimeKey creates the prefix for all unbonding delegations from a delegator
func GetUnbondingDelegationTimeKey(timestamp time.Time) []byte {
	_ = "STUB: not implemented"
	return nil
}

// GetREDKey returns a key prefix for indexing a redelegation from a delegator
// and source validator to a destination validator.
func GetREDKey(delAddr sdk.AccAddress, valSrcAddr, valDstAddr sdk.ValAddress) []byte {
	_ = "STUB: not implemented"
	// key is of the form GetREDsKey || valSrcAddrLen (1 byte) || valSrcAddr || valDstAddrLen (1 byte) || valDstAddr
	return nil
}

// GetREDByValSrcIndexKey creates the index-key for a redelegation, stored by source-validator-index
// VALUE: none (key rearrangement used)
func GetREDByValSrcIndexKey(delAddr sdk.AccAddress, valSrcAddr, valDstAddr sdk.ValAddress) []byte {
	_ = "STUB: not implemented"
	return nil
}

// key is of the form REDSFromValsSrcKey || delAddrLen (1 byte) || delAddr || valDstAddrLen (1 byte) || valDstAddr

// GetREDByValDstIndexKey creates the index-key for a redelegation, stored by destination-validator-index
// VALUE: none (key rearrangement used)
func GetREDByValDstIndexKey(delAddr sdk.AccAddress, valSrcAddr, valDstAddr sdk.ValAddress) []byte {
	_ = "STUB: not implemented"
	return nil
}

// key is of the form REDSToValsDstKey || delAddrLen (1 byte) || delAddr || valSrcAddrLen (1 byte) || valSrcAddr

// GetREDKeyFromValSrcIndexKey rearranges the ValSrcIndexKey to get the REDKey
func GetREDKeyFromValSrcIndexKey(indexKey []byte) []byte {
	_ = "STUB: not implemented"
	// note that first byte is prefix byte, which we remove
	return nil
}

// GetREDKeyFromValDstIndexKey rearranges the ValDstIndexKey to get the REDKey
func GetREDKeyFromValDstIndexKey(indexKey []byte) []byte {
	_ = "STUB: not implemented"
	// note that first byte is prefix byte, which we remove
	return nil
}

// GetRedelegationTimeKey returns a key prefix for indexing an unbonding
// redelegation based on a completion time.
func GetRedelegationTimeKey(timestamp time.Time) []byte { _ = "STUB: not implemented"; return nil }

// GetREDsKey returns a key prefix for indexing a redelegation from a delegator
// address.
func GetREDsKey(delAddr sdk.AccAddress) []byte { _ = "STUB: not implemented"; return nil }

// GetREDsFromValSrcIndexKey returns a key prefix for indexing a redelegation to
// a source validator.
func GetREDsFromValSrcIndexKey(valSrcAddr sdk.ValAddress) []byte {
	_ = "STUB: not implemented"
	return nil
}

// GetREDsToValDstIndexKey returns a key prefix for indexing a redelegation to a
// destination (target) validator.
func GetREDsToValDstIndexKey(valDstAddr sdk.ValAddress) []byte {
	_ = "STUB: not implemented"
	return nil
}

// GetREDsByDelToValDstIndexKey returns a key prefix for indexing a redelegation
// from an address to a source validator.
func GetREDsByDelToValDstIndexKey(delAddr sdk.AccAddress, valDstAddr sdk.ValAddress) []byte {
	_ = "STUB: not implemented"
	return nil
}

// GetHistoricalInfoKey returns a key prefix for indexing HistoricalInfo objects.
func GetHistoricalInfoKey(height int64) []byte { _ = "STUB: not implemented"; return nil }
