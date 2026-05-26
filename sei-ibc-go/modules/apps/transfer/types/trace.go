package types

import (
	"sort"

	tmbytes "github.com/sei-protocol/sei-chain/sei-tendermint/libs/bytes"
)

// ParseDenomTrace parses a string with the ibc prefix (denom trace) and the base denomination
// into a DenomTrace type.
//
// Examples:
//
// - "portidone/channel-0/uatom" => DenomTrace{Path: "portidone/channel-0", BaseDenom: "uatom"}
// - "portidone/channel-0/portidtwo/channel-1/uatom" => DenomTrace{Path: "portidone/channel-0/portidtwo/channel-1", BaseDenom: "uatom"}
// - "portidone/channel-0/gamm/pool/1" => DenomTrace{Path: "portidone/channel-0", BaseDenom: "gamm/pool/1"}
// - "gamm/pool/1" => DenomTrace{Path: "", BaseDenom: "gamm/pool/1"}
// - "uatom" => DenomTrace{Path: "", BaseDenom: "uatom"}
func ParseDenomTrace(rawDenom string) DenomTrace {
	_ = "STUB: not implemented"
	return *new(DenomTrace)
}

// Hash returns the hex bytes of the SHA256 hash of the DenomTrace fields using the following formula:
//
// hash = sha256(tracePath + "/" + baseDenom)
func (dt DenomTrace) Hash() tmbytes.HexBytes {
	_ = "STUB: not implemented"
	return *new(tmbytes.HexBytes)
}

// GetPrefix returns the receiving denomination prefix composed by the trace info and a separator.
func (dt DenomTrace) GetPrefix() string { _ = "STUB: not implemented"; return "" }

// IBCDenom a coin denomination for an ICS20 fungible token in the format
// 'ibc/{hash(tracePath + baseDenom)}'. If the trace is empty, it will return the base denomination.
func (dt DenomTrace) IBCDenom() string { _ = "STUB: not implemented"; return "" }

// GetFullDenomPath returns the full denomination according to the ICS20 specification:
// tracePath + "/" + baseDenom
// If there exists no trace then the base denomination is returned.
func (dt DenomTrace) GetFullDenomPath() string { _ = "STUB: not implemented"; return "" }

// extractPathAndBaseFromFullDenom returns the trace path and the base denom from
// the elements that constitute the complete denom.
func extractPathAndBaseFromFullDenom(fullDenomItems []string) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

// The IBC specification does not guarantee the expected format of the
// destination port or destination channel identifier. A short term solution
// to determine base denomination is to expect the channel identifier to be the
// one ibc-go specifies. A longer term solution is to separate the path and base
// denomination in the ICS20 packet. If an intermediate hop prefixes the full denom
// with a channel identifier format different from our own, the base denomination
// will be incorrectly parsed, but the token will continue to be treated correctly
// as an IBC denomination. The hash used to store the token internally on our chain
// will be the same value as the base denomination being correctly parsed.

func validateTraceIdentifiers(identifiers []string) error { _ = "STUB: not implemented"; return nil }

// validate correctness of port and channel identifiers

// Validate performs a basic validation of the DenomTrace fields.
func (dt DenomTrace) Validate() error {
	_ = "STUB: not implemented"
	// empty trace is accepted when token lives on the original chain
	return nil
}

// NOTE: no base denomination validation

// Traces defines a wrapper type for a slice of DenomTrace.
type Traces []DenomTrace

// Validate performs a basic validation of each denomination trace info.
func (t Traces) Validate() error { _ = "STUB: not implemented"; return nil }

var _ sort.Interface = Traces{}

// Len implements sort.Interface for Traces
func (t Traces) Len() int {
	_ = "STUB: not implemented"

	// Less implements sort.Interface for Traces
	return 0
}

func (t Traces) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Swap implements sort.Interface for Traces
func (t Traces) Swap(i, j int) { _ = "STUB: not implemented"; return }

// Sort is a helper function to sort the set of denomination traces in-place
func (t Traces) Sort() Traces {
	_ = "STUB: not implemented"
	return *

	// ValidatePrefixedDenom checks that the denomination for an IBC fungible token packet denom is correctly prefixed.
	// The function will return no error if the given string follows one of the two formats:
	//
	//   - Prefixed denomination: '{portIDN}/{channelIDN}/.../{portID0}/{channelID0}/baseDenom'
	//   - Unprefixed denomination: 'baseDenom'
	//
	// 'baseDenom' may or may not contain '/'s
	new(Traces)
}

func ValidatePrefixedDenom(denom string) error { _ = "STUB: not implemented"; return nil }

// NOTE: no base denomination validation

// NOTE: base denom contains slashes, so no base denomination validation

// ValidateIBCDenom validates that the given denomination is either:
//
//   - A valid base denomination (eg: 'uatom' or 'gamm/pool/1' as in https://github.com/cosmos/ibc-go/issues/894)
//   - A valid fungible token representation (i.e 'ibc/{hash}') per ADR 001 https://github.com/cosmos/ibc-go/blob/main/docs/architecture/adr-001-coin-source-tracing.md
func ValidateIBCDenom(denom string) error { _ = "STUB: not implemented"; return nil }

// ParseHexHash parses a hex hash in string format to bytes and validates its correctness.
func ParseHexHash(hexHash string) (tmbytes.HexBytes, error) {
	_ = "STUB: not implemented"
	return *new(tmbytes.HexBytes), nil
}
