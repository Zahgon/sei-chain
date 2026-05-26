package types

import (
	"time"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const (
	// ModuleName is the name of the module
	ModuleName = "gov"

	// StoreKey is the store key string for gov
	StoreKey = ModuleName

	// RouterKey is the message route for gov
	RouterKey = ModuleName

	// QuerierRoute is the querier route for gov
	QuerierRoute = ModuleName
)

// Keys for governance store
// Items are stored with the following key: values
//
// - 0x00<proposalID_Bytes>: Proposal
//
// - 0x01<endTime_Bytes><proposalID_Bytes>: activeProposalID
//
// - 0x02<endTime_Bytes><proposalID_Bytes>: inactiveProposalID
//
// - 0x03: nextProposalID
//
// - 0x10<proposalID_Bytes><depositorAddrLen (1 Byte)><depositorAddr_Bytes>: Deposit
//
// - 0x20<proposalID_Bytes><voterAddrLen (1 Byte)><voterAddr_Bytes>: Voter
var (
	ProposalsKeyPrefix          = []byte{0x00}
	ActiveProposalQueuePrefix   = []byte{0x01}
	InactiveProposalQueuePrefix = []byte{0x02}
	ProposalIDKey               = []byte{0x03}

	DepositsKeyPrefix = []byte{0x10}

	VotesKeyPrefix = []byte{0x20}
)

var lenTime = len(sdk.FormatTimeBytes(time.Now()))

// GetProposalIDBytes returns the byte representation of the proposalID
func GetProposalIDBytes(proposalID uint64) (proposalIDBz []byte) {
	_ = "STUB: not implemented"
	return nil
}

// GetProposalIDFromBytes returns proposalID in uint64 format from a byte array
func GetProposalIDFromBytes(bz []byte) (proposalID uint64) { _ = "STUB: not implemented"; return 0 }

// ProposalKey gets a specific proposal from the store
func ProposalKey(proposalID uint64) []byte { _ = "STUB: not implemented"; return nil }

// ActiveProposalByTimeKey gets the active proposal queue key by endTime
func ActiveProposalByTimeKey(endTime time.Time) []byte { _ = "STUB: not implemented"; return nil }

// ActiveProposalQueueKey returns the key for a proposalID in the activeProposalQueue
func ActiveProposalQueueKey(proposalID uint64, endTime time.Time) []byte {
	_ = "STUB: not implemented"
	return nil
}

// InactiveProposalByTimeKey gets the inactive proposal queue key by endTime
func InactiveProposalByTimeKey(endTime time.Time) []byte { _ = "STUB: not implemented"; return nil }

// InactiveProposalQueueKey returns the key for a proposalID in the inactiveProposalQueue
func InactiveProposalQueueKey(proposalID uint64, endTime time.Time) []byte {
	_ = "STUB: not implemented"
	return nil
}

// DepositsKey gets the first part of the deposits key based on the proposalID
func DepositsKey(proposalID uint64) []byte { _ = "STUB: not implemented"; return nil }

// DepositKey key of a specific deposit from the store
func DepositKey(proposalID uint64, depositorAddr sdk.AccAddress) []byte {
	_ = "STUB: not implemented"
	return nil
}

// VotesKey gets the first part of the votes key based on the proposalID
func VotesKey(proposalID uint64) []byte { _ = "STUB: not implemented"; return nil }

// VoteKey key of a specific vote from the store
func VoteKey(proposalID uint64, voterAddr sdk.AccAddress) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Split keys function; used for iterators

// SplitProposalKey split the proposal key and returns the proposal id
func SplitProposalKey(key []byte) (proposalID uint64) { _ = "STUB: not implemented"; return 0 }

// SplitActiveProposalQueueKey split the active proposal key and returns the proposal id and endTime
func SplitActiveProposalQueueKey(key []byte) (proposalID uint64, endTime time.Time) {
	_ = "STUB: not implemented"
	return 0, *new(time.Time)
}

// SplitInactiveProposalQueueKey split the inactive proposal key and returns the proposal id and endTime
func SplitInactiveProposalQueueKey(key []byte) (proposalID uint64, endTime time.Time) {
	_ = "STUB: not implemented"
	return 0, *new(time.Time)
}

// SplitKeyDeposit split the deposits key and returns the proposal id and depositor address
func SplitKeyDeposit(key []byte) (proposalID uint64, depositorAddr sdk.AccAddress) {
	_ = "STUB: not implemented"
	return 0, *new(sdk.AccAddress)
}

// SplitKeyVote split the votes key and returns the proposal id and voter address
func SplitKeyVote(key []byte) (proposalID uint64, voterAddr sdk.AccAddress) {
	_ = "STUB: not implemented"
	return 0, *new(sdk.AccAddress)
}

// private functions

func splitKeyWithTime(key []byte) (proposalID uint64, endTime time.Time) {
	_ = "STUB: not implemented"
	return 0, *new(time.Time)
}

func splitKeyWithAddress(key []byte) (proposalID uint64, addr sdk.AccAddress) {
	_ = "STUB: not implemented"
	// Both Vote and Deposit store keys are of format:
	// <prefix (1 Byte)><proposalID (8 bytes)><addrLen (1 Byte)><addr_Bytes>
	return 0, *new(sdk.AccAddress)
}
