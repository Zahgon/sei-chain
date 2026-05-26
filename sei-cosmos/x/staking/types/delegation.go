package types

import (
	"time"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// Implements Delegation interface
var _ DelegationI = Delegation{}

// String implements the Stringer interface for a DVPair object.
func (dv DVPair) String() string { _ = "STUB: not implemented"; return "" }

// String implements the Stringer interface for a DVVTriplet object.
func (dvv DVVTriplet) String() string { _ = "STUB: not implemented"; return "" }

// NewDelegation creates a new delegation object
func NewDelegation(delegatorAddr sdk.AccAddress, validatorAddr sdk.ValAddress, shares sdk.Dec) Delegation {
	_ = "STUB: not implemented"
	return *new(Delegation)
}

// MustMarshalDelegation returns the delegation bytes. Panics if fails
func MustMarshalDelegation(cdc codec.BinaryCodec, delegation Delegation) []byte {
	_ = "STUB: not implemented"
	return nil
}

// MustUnmarshalDelegation return the unmarshaled delegation from bytes.
// Panics if fails.
func MustUnmarshalDelegation(cdc codec.BinaryCodec, value []byte) Delegation {
	_ = "STUB: not implemented"
	return *new(Delegation)
}

// return the delegation
func UnmarshalDelegation(cdc codec.BinaryCodec, value []byte) (delegation Delegation, err error) {
	_ = "STUB: not implemented"
	return *new(Delegation), nil
}

func (d Delegation) GetDelegatorAddr() sdk.AccAddress {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress)
}

func (d Delegation) GetValidatorAddr() sdk.ValAddress {
	_ = "STUB: not implemented"
	return *new(sdk.ValAddress)
}

func (d Delegation) GetShares() sdk.Dec {
	_ = "STUB: not implemented"

	// String returns a human readable string representation of a Delegation.
	return *new(sdk.Dec)
}

func (d Delegation) String() string { _ = "STUB: not implemented"; return "" }

// Delegations is a collection of delegations
type Delegations []Delegation

func (d Delegations) String() (out string) { _ = "STUB: not implemented"; return "" }

func NewUnbondingDelegationEntry(creationHeight int64, completionTime time.Time, balance sdk.Int) UnbondingDelegationEntry {
	_ = "STUB: not implemented"
	return *new(UnbondingDelegationEntry)
}

// String implements the stringer interface for a UnbondingDelegationEntry.
func (e UnbondingDelegationEntry) String() string { _ = "STUB: not implemented"; return "" }

// IsMature - is the current entry mature
func (e UnbondingDelegationEntry) IsMature(currentTime time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

// NewUnbondingDelegation - create a new unbonding delegation object
func NewUnbondingDelegation(
	delegatorAddr sdk.AccAddress, validatorAddr sdk.ValAddress,
	creationHeight int64, minTime time.Time, balance sdk.Int,
) UnbondingDelegation {
	_ = "STUB: not implemented"
	return *new(UnbondingDelegation)
}

// AddEntry - append entry to the unbonding delegation
func (ubd *UnbondingDelegation) AddEntry(creationHeight int64, minTime time.Time, balance sdk.Int) {
	_ = "STUB: not implemented"
	return
}

// RemoveEntry - remove entry at index i to the unbonding delegation
func (ubd *UnbondingDelegation) RemoveEntry(i int64) { _ = "STUB: not implemented"; return }

// return the unbonding delegation
func MustMarshalUBD(cdc codec.BinaryCodec, ubd UnbondingDelegation) []byte {
	_ = "STUB: not implemented"
	return nil
}

// unmarshal a unbonding delegation from a store value
func MustUnmarshalUBD(cdc codec.BinaryCodec, value []byte) UnbondingDelegation {
	_ = "STUB: not implemented"
	return *new(UnbondingDelegation)
}

// unmarshal a unbonding delegation from a store value
func UnmarshalUBD(cdc codec.BinaryCodec, value []byte) (ubd UnbondingDelegation, err error) {
	_ = "STUB: not implemented"
	return *new(UnbondingDelegation), nil
}

// String returns a human readable string representation of an UnbondingDelegation.
func (ubd UnbondingDelegation) String() string { _ = "STUB: not implemented"; return "" }

// UnbondingDelegations is a collection of UnbondingDelegation
type UnbondingDelegations []UnbondingDelegation

func (ubds UnbondingDelegations) String() (out string) { _ = "STUB: not implemented"; return "" }

func NewRedelegationEntry(creationHeight int64, completionTime time.Time, balance sdk.Int, sharesDst sdk.Dec) RedelegationEntry {
	_ = "STUB: not implemented"
	return *new(RedelegationEntry)
}

// String implements the Stringer interface for a RedelegationEntry object.
func (e RedelegationEntry) String() string { _ = "STUB: not implemented"; return "" }

// IsMature - is the current entry mature
func (e RedelegationEntry) IsMature(currentTime time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

func NewRedelegation(
	delegatorAddr sdk.AccAddress, validatorSrcAddr, validatorDstAddr sdk.ValAddress,
	creationHeight int64, minTime time.Time, balance sdk.Int, sharesDst sdk.Dec,
) Redelegation {
	_ = "STUB: not implemented"
	return *new(Redelegation)
}

// AddEntry - append entry to the unbonding delegation
func (red *Redelegation) AddEntry(creationHeight int64, minTime time.Time, balance sdk.Int, sharesDst sdk.Dec) {
	_ = "STUB: not implemented"
	return
}

// RemoveEntry - remove entry at index i to the unbonding delegation
func (red *Redelegation) RemoveEntry(i int64) { _ = "STUB: not implemented"; return }

// MustMarshalRED returns the Redelegation bytes. Panics if fails.
func MustMarshalRED(cdc codec.BinaryCodec, red Redelegation) []byte {
	_ = "STUB: not implemented"
	return nil
}

// MustUnmarshalRED unmarshals a redelegation from a store value. Panics if fails.
func MustUnmarshalRED(cdc codec.BinaryCodec, value []byte) Redelegation {
	_ = "STUB: not implemented"
	return *new(Redelegation)
}

// UnmarshalRED unmarshals a redelegation from a store value
func UnmarshalRED(cdc codec.BinaryCodec, value []byte) (red Redelegation, err error) {
	_ = "STUB: not implemented"
	return *new(Redelegation), nil
}

// String returns a human readable string representation of a Redelegation.
func (red Redelegation) String() string { _ = "STUB: not implemented"; return "" }

// Redelegations are a collection of Redelegation
type Redelegations []Redelegation

func (d Redelegations) String() (out string) { _ = "STUB: not implemented"; return "" }

// ----------------------------------------------------------------------------
// Client Types

// NewDelegationResp creates a new DelegationResponse instance
func NewDelegationResp(
	delegatorAddr sdk.AccAddress, validatorAddr sdk.ValAddress, shares sdk.Dec, balance sdk.Coin,
) DelegationResponse {
	_ = "STUB: not implemented"
	return *new(DelegationResponse)
}

// String implements the Stringer interface for DelegationResponse.
func (d DelegationResponse) String() string { _ = "STUB: not implemented"; return "" }

type delegationRespAlias DelegationResponse

// MarshalJSON implements the json.Marshaler interface. This is so we can
// achieve a flattened structure while embedding other types.
func (d DelegationResponse) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface. This is so we can
// achieve a flattened structure while embedding other types.
func (d *DelegationResponse) UnmarshalJSON(bz []byte) error { _ = "STUB: not implemented"; return nil }

// DelegationResponses is a collection of DelegationResp
type DelegationResponses []DelegationResponse

// String implements the Stringer interface for DelegationResponses.
func (d DelegationResponses) String() (out string) { _ = "STUB: not implemented"; return "" }

// NewRedelegationResponse crates a new RedelegationEntryResponse instance.
func NewRedelegationResponse(
	delegatorAddr sdk.AccAddress, validatorSrc, validatorDst sdk.ValAddress, entries []RedelegationEntryResponse,
) RedelegationResponse {
	_ = "STUB: not implemented"
	return *new(RedelegationResponse)
}

// NewRedelegationEntryResponse creates a new RedelegationEntryResponse instance.
func NewRedelegationEntryResponse(
	creationHeight int64, completionTime time.Time, sharesDst sdk.Dec, initialBalance, balance sdk.Int) RedelegationEntryResponse {
	_ = "STUB: not implemented"
	return *new(RedelegationEntryResponse)
}

type redelegationRespAlias RedelegationResponse

// MarshalJSON implements the json.Marshaler interface. This is so we can
// achieve a flattened structure while embedding other types.
func (r RedelegationResponse) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface. This is so we can
// achieve a flattened structure while embedding other types.
func (r *RedelegationResponse) UnmarshalJSON(bz []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// RedelegationResponses are a collection of RedelegationResp
type RedelegationResponses []RedelegationResponse

func (r RedelegationResponses) String() (out string) { _ = "STUB: not implemented"; return "" }
