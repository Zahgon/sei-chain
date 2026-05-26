package types

import (
	"regexp"
)

const (
	// SubModuleName defines the IBC client name
	SubModuleName string = "client"

	// RouterKey is the message route for IBC client
	RouterKey string = SubModuleName

	// QuerierRoute is the querier route for IBC client
	QuerierRoute string = SubModuleName

	// KeyNextClientSequence is the key used to store the next client sequence in
	// the keeper.
	KeyNextClientSequence = "nextClientSequence"
)

// FormatClientIdentifier returns the client identifier with the sequence appended.
// This is a SDK specific format not enforced by IBC protocol.
func FormatClientIdentifier(clientType string, sequence uint64) string {
	_ = "STUB: not implemented"
	return ""
}

// IsClientIDFormat checks if a clientID is in the format required on the SDK for
// parsing client identifiers. The client identifier must be in the form: `{client-type}-{N}
var IsClientIDFormat = regexp.MustCompile(`^.*[^\n-]-[0-9]{1,20}$`).MatchString

// IsValidClientID checks if the clientID is valid and can be parsed into the client
// identifier format.
func IsValidClientID(clientID string) bool { _ = "STUB: not implemented"; return false }

// ParseClientIdentifier parses the client type and sequence from the client identifier.
func ParseClientIdentifier(clientID string) (string, uint64, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}
