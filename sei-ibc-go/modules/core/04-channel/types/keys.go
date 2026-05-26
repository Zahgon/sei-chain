package types

import (
	"regexp"
)

const (
	// SubModuleName defines the IBC channels name
	SubModuleName = "channel"

	// StoreKey is the store key string for IBC channels
	StoreKey = SubModuleName

	// RouterKey is the message route for IBC channels
	RouterKey = SubModuleName

	// QuerierRoute is the querier route for IBC channels
	QuerierRoute = SubModuleName

	// KeyNextChannelSequence is the key used to store the next channel sequence in
	// the keeper.
	KeyNextChannelSequence = "nextChannelSequence"

	// ChannelPrefix is the prefix used when creating a channel identifier
	ChannelPrefix = "channel-"
)

// FormatChannelIdentifier returns the channel identifier with the sequence appended.
// This is a SDK specific format not enforced by IBC protocol.
func FormatChannelIdentifier(sequence uint64) string { _ = "STUB: not implemented"; return "" }

// IsChannelIDFormat checks if a channelID is in the format required on the SDK for
// parsing channel identifiers. The channel identifier must be in the form: `channel-{N}
var IsChannelIDFormat = regexp.MustCompile(`^channel-[0-9]{1,20}$`).MatchString

// IsValidChannelID checks if a channelID is valid and can be parsed to the channel
// identifier format.
func IsValidChannelID(channelID string) bool { _ = "STUB: not implemented"; return false }

// ParseChannelSequence parses the channel sequence from the channel identifier.
func ParseChannelSequence(channelID string) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
