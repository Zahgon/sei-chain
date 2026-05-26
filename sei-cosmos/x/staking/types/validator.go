package types

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	tmprotocrypto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/crypto"
)

const (
	// TODO: Why can't we just have one string description which can be JSON by convention
	MaxMonikerLength         = 70
	MaxIdentityLength        = 3000
	MaxWebsiteLength         = 140
	MaxSecurityContactLength = 140
	MaxDetailsLength         = 280
)

var (
	BondStatusUnspecified = BondStatus_name[int32(Unspecified)]
	BondStatusUnbonded    = BondStatus_name[int32(Unbonded)]
	BondStatusUnbonding   = BondStatus_name[int32(Unbonding)]
	BondStatusBonded      = BondStatus_name[int32(Bonded)]
)

var _ ValidatorI = Validator{}

// NewValidator constructs a new Validator
func NewValidator(operator sdk.ValAddress, pubKey cryptotypes.PubKey, description Description) (Validator, error) {
	_ = "STUB: not implemented"
	return *new(Validator), nil
}

// String implements the Stringer interface for a Validator object.
func (v Validator) String() string { _ = "STUB: not implemented"; return "" }

// Validators is a collection of Validator
type Validators []Validator

func (v Validators) String() (out string) { _ = "STUB: not implemented"; return "" }

// ToSDKValidators -  convenience function convert []Validator to []sdk.ValidatorI
func (v Validators) ToSDKValidators() (validators []ValidatorI) {
	_ = "STUB: not implemented"
	return nil
}

// Sort Validators sorts validator array in ascending operator address order
func (v Validators) Sort() {
	_ = "STUB: not implemented"

	// Implements sort interface
	return
}

func (v Validators) Len() int {
	_ = "STUB: not implemented"

	// Implements sort interface
	return 0
}

func (v Validators) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Implements sort interface
func (v Validators) Swap(i, j int) { _ = "STUB: not implemented"; return }

// ValidatorsByVotingPower implements sort.Interface for []Validator based on
// the VotingPower and Address fields.
// The validators are sorted first by their voting power (descending). Secondary index - Address (ascending).
// Copied from tendermint/types/validator_set.go
type ValidatorsByVotingPower []Validator

func (valz ValidatorsByVotingPower) Len() int { _ = "STUB: not implemented"; return 0 }

func (valz ValidatorsByVotingPower) Less(i, j int, r sdk.Int) bool {
	_ = "STUB: not implemented"
	return false
}

// If either returns error, then return false

func (valz ValidatorsByVotingPower) Swap(i, j int) { _ = "STUB: not implemented"; return }

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (v Validators) UnpackInterfaces(c codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// return the redelegation
func MustMarshalValidator(cdc codec.BinaryCodec, validator *Validator) []byte {
	_ = "STUB: not implemented"
	return nil
}

// unmarshal a redelegation from a store value
func MustUnmarshalValidator(cdc codec.BinaryCodec, value []byte) Validator {
	_ = "STUB: not implemented"
	return *new(Validator)
}

// unmarshal a redelegation from a store value
func UnmarshalValidator(cdc codec.BinaryCodec, value []byte) (v Validator, err error) {
	_ = "STUB: not implemented"
	return *new(Validator), nil
}

// IsBonded checks if the validator status equals Bonded
func (v Validator) IsBonded() bool { _ = "STUB: not implemented"; return false }

// IsUnbonded checks if the validator status equals Unbonded
func (v Validator) IsUnbonded() bool { _ = "STUB: not implemented"; return false }

// IsUnbonding checks if the validator status equals Unbonding
func (v Validator) IsUnbonding() bool { _ = "STUB: not implemented"; return false }

// constant used in flags to indicate that description field should not be updated
const DoNotModifyDesc = "[do-not-modify]"

func NewDescription(moniker, identity, website, securityContact, details string) Description {
	_ = "STUB: not implemented"
	return *new(Description)
}

// String implements the Stringer interface for a Description object.
func (d Description) String() string { _ = "STUB: not implemented"; return "" }

// UpdateDescription updates the fields of a given description. An error is
// returned if the resulting description contains an invalid length.
func (d Description) UpdateDescription(d2 Description) (Description, error) {
	_ = "STUB: not implemented"
	return *new(Description), nil
}

// EnsureLength ensures the length of a validator's description.
func (d Description) EnsureLength() (Description, error) {
	_ = "STUB: not implemented"
	return *new(Description), nil
}

// ABCIValidatorUpdate returns an abci.ValidatorUpdate from a staking validator type
// with the full validator power
func (v Validator) ABCIValidatorUpdate(r sdk.Int) abci.ValidatorUpdate {
	_ = "STUB: not implemented"
	return *new(abci.ValidatorUpdate)
}

// ABCIValidatorUpdateZero returns an abci.ValidatorUpdate from a staking validator type
// with zero power used for validator updates.
func (v Validator) ABCIValidatorUpdateZero() abci.ValidatorUpdate {
	_ = "STUB: not implemented"
	return *new(abci.ValidatorUpdate)
}

// SetInitialCommission attempts to set a validator's initial commission. An
// error is returned if the commission is invalid.
func (v Validator) SetInitialCommission(commission Commission) (Validator, error) {
	_ = "STUB: not implemented"
	return *new(Validator), nil
}

// In some situations, the exchange rate becomes invalid, e.g. if
// Validator loses all tokens due to slashing. In this case,
// make all future delegations invalid.
func (v Validator) InvalidExRate() bool { _ = "STUB: not implemented"; return false }

// calculate the token worth of provided shares
func (v Validator) TokensFromShares(shares sdk.Dec) sdk.Dec {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

// calculate the token worth of provided shares, truncated
func (v Validator) TokensFromSharesTruncated(shares sdk.Dec) sdk.Dec {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

// TokensFromSharesRoundUp returns the token worth of provided shares, rounded
// up.
func (v Validator) TokensFromSharesRoundUp(shares sdk.Dec) sdk.Dec {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

// SharesFromTokens returns the shares of a delegation given a bond amount. It
// returns an error if the validator has no tokens.
func (v Validator) SharesFromTokens(amt sdk.Int) (sdk.Dec, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Dec), nil
}

// SharesFromTokensTruncated returns the truncated shares of a delegation given
// a bond amount. It returns an error if the validator has no tokens.
func (v Validator) SharesFromTokensTruncated(amt sdk.Int) (sdk.Dec, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Dec), nil
}

// get the bonded tokens which the validator holds
func (v Validator) BondedTokens() sdk.Int { _ = "STUB: not implemented"; return *new(sdk.Int) }

// ConsensusPower gets the consensus-engine power. Aa reduction of 10^6 from
// validator tokens is applied
func (v Validator) ConsensusPower(r sdk.Int) int64 { _ = "STUB: not implemented"; return 0 }

// PotentialConsensusPower returns the potential consensus-engine power.
func (v Validator) PotentialConsensusPower(r sdk.Int) int64 { _ = "STUB: not implemented"; return 0 }

// UpdateStatus updates the location of the shares within a validator
// to reflect the new status
func (v Validator) UpdateStatus(newStatus BondStatus) Validator {
	_ = "STUB: not implemented"
	return *new(Validator)
}

// AddTokensFromDel adds tokens to a validator
func (v Validator) AddTokensFromDel(amount sdk.Int) (Validator, sdk.Dec) {
	_ = "STUB: not implemented"
	// calculate the shares to issue
	return *new(Validator), *new(sdk.Dec)
}

// the first delegation to a validator sets the exchange rate to one

// RemoveTokens removes tokens from a validator
func (v Validator) RemoveTokens(tokens sdk.Int) Validator {
	_ = "STUB: not implemented"
	return *new(Validator)
}

// RemoveDelShares removes delegator shares from a validator.
// NOTE: because token fractions are left in the valiadator,
//
//	the exchange rate of future shares of this validator can increase.
func (v Validator) RemoveDelShares(delShares sdk.Dec) (Validator, sdk.Int) {
	_ = "STUB: not implemented"
	return *new(Validator), *new(sdk.Int)
}

// last delegation share gets any trimmings

// leave excess tokens in the validator
// however fully use all the delegator shares

// MinEqual defines a more minimum set of equality conditions when comparing two
// validators.
func (v *Validator) MinEqual(other *Validator) bool { _ = "STUB: not implemented"; return false }

// Equal checks if the receiver equals the parameter
func (v *Validator) Equal(v2 *Validator) bool { _ = "STUB: not implemented"; return false }

func (v Validator) IsJailed() bool        { _ = "STUB: not implemented"; return false }
func (v Validator) GetMoniker() string    { _ = "STUB: not implemented"; return "" }
func (v Validator) GetStatus() BondStatus { _ = "STUB: not implemented"; return *new(BondStatus) }
func (v Validator) GetOperator() sdk.ValAddress {
	_ = "STUB: not implemented"
	return *new(sdk.ValAddress)
}

// ConsPubKey returns the validator PubKey as a cryptotypes.PubKey.
func (v Validator) ConsPubKey() (cryptotypes.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(cryptotypes.PubKey), nil
}

// TmConsPublicKey casts Validator.ConsensusPubkey to tmprotocrypto.PubKey.
func (v Validator) TmConsPublicKey() (tmprotocrypto.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(tmprotocrypto.PublicKey), nil
}

// GetConsAddr extracts Consensus key address
func (v Validator) GetConsAddr() (sdk.ConsAddress, error) {
	_ = "STUB: not implemented"
	return *new(sdk.ConsAddress), nil
}

func (v Validator) GetTokens() sdk.Int                { _ = "STUB: not implemented"; return *new(sdk.Int) }
func (v Validator) GetBondedTokens() sdk.Int          { _ = "STUB: not implemented"; return *new(sdk.Int) }
func (v Validator) GetConsensusPower(r sdk.Int) int64 { _ = "STUB: not implemented"; return 0 }

func (v Validator) GetCommission() sdk.Dec        { _ = "STUB: not implemented"; return *new(sdk.Dec) }
func (v Validator) GetMinSelfDelegation() sdk.Int { _ = "STUB: not implemented"; return *new(sdk.Int) }
func (v Validator) GetDelegatorShares() sdk.Dec {
	_ = "STUB: not implemented"
	return *

	// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
	new(sdk.Dec)
}

func (v Validator) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}
