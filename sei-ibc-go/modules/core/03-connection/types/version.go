package types

import (
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

var (
	// DefaultIBCVersion represents the latest supported version of IBC used
	// in connection version negotiation. The current version supports only
	// ORDERED and UNORDERED channels and requires at least one channel type
	// to be agreed upon.
	DefaultIBCVersion = NewVersion(DefaultIBCVersionIdentifier, []string{"ORDER_ORDERED", "ORDER_UNORDERED"})

	// DefaultIBCVersionIdentifier is the IBC v1.0.0 protocol version identifier
	DefaultIBCVersionIdentifier = "1"

	// AllowNilFeatureSet is a helper map to indicate if a specified version
	// identifier is allowed to have a nil feature set. Any versions supported,
	// but not included in the map default to not supporting nil feature sets.
	allowNilFeatureSet = map[string]bool{
		DefaultIBCVersionIdentifier: false,
	}
)

var _ exported.Version = &Version{}

// NewVersion returns a new instance of Version.
func NewVersion(identifier string, features []string) *Version {
	_ = "STUB: not implemented"
	return nil
}

// GetIdentifier implements the VersionI interface
func (version Version) GetIdentifier() string { _ = "STUB: not implemented"; return "" }

// GetFeatures implements the VersionI interface
func (version Version) GetFeatures() []string { _ = "STUB: not implemented"; return nil }

// ValidateVersion does basic validation of the version identifier and
// features. It unmarshals the version string into a Version object.
func ValidateVersion(version *Version) error { _ = "STUB: not implemented"; return nil }

// VerifyProposedVersion verifies that the entire feature set in the
// proposed version is supported by this chain. If the feature set is
// empty it verifies that this is allowed for the specified version
// identifier.
func (version Version) VerifyProposedVersion(proposedVersion exported.Version) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifySupportedFeature takes in a version and feature string and returns
// true if the feature is supported by the version and false otherwise.
func VerifySupportedFeature(version exported.Version, feature string) bool {
	_ = "STUB: not implemented"
	return false
}

// GetCompatibleVersions returns a descending ordered set of compatible IBC
// versions for the caller chain's connection end. The latest supported
// version should be first element and the set should descend to the oldest
// supported version.
func GetCompatibleVersions() []exported.Version { _ = "STUB: not implemented"; return nil }

// IsSupportedVersion returns true if the proposed version has a matching version
// identifier and its entire feature set is supported or the version identifier
// supports an empty feature set.
func IsSupportedVersion(proposedVersion *Version) bool { _ = "STUB: not implemented"; return false }

// FindSupportedVersion returns the version with a matching version identifier
// if it exists. The returned boolean is true if the version is found and
// false otherwise.
func FindSupportedVersion(version exported.Version, supportedVersions []exported.Version) (exported.Version, bool) {
	_ = "STUB: not implemented"
	return *new(exported.Version), false
}

// PickVersion iterates over the descending ordered set of compatible IBC
// versions and selects the first version with a version identifier that is
// supported by the counterparty. The returned version contains a feature
// set with the intersection of the features supported by the source and
// counterparty chains. If the feature set intersection is nil and this is
// not allowed for the chosen version identifier then the search for a
// compatible version continues. This function is called in the ConnOpenTry
// handshake procedure.
//
// CONTRACT: PickVersion must only provide a version that is in the
// intersection of the supported versions and the counterparty versions.
func PickVersion(supportedVersions, counterpartyVersions []exported.Version) (*Version, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check if the source version is supported by the counterparty

// GetFeatureSetIntersection returns the intersections of source feature set
// and the counterparty feature set. This is done by iterating over all the
// features in the source version and seeing if they exist in the feature
// set for the counterparty version.
func GetFeatureSetIntersection(sourceFeatureSet, counterpartyFeatureSet []string) (featureSet []string) {
	_ = "STUB: not implemented"
	return nil
}

// ExportedVersionsToProto casts a slice of the Version interface to a slice
// of the Version proto definition.
func ExportedVersionsToProto(exportedVersions []exported.Version) []*Version {
	_ = "STUB: not implemented"
	return nil
}

// ProtoVersionsToExported converts a slice of the Version proto definition to
// the Version interface.
func ProtoVersionsToExported(versions []*Version) []exported.Version {
	_ = "STUB: not implemented"
	return nil
}

// contains returns true if the provided string element exists within the
// string set.
func contains(elem string, set []string) bool { _ = "STUB: not implemented"; return false }
