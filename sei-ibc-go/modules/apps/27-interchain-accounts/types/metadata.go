package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const (
	// EncodingProtobuf defines the protocol buffers proto3 encoding format
	EncodingProtobuf = "proto3"

	// TxTypeSDKMultiMsg defines the multi message transaction type supported by the Cosmos SDK
	TxTypeSDKMultiMsg = "sdk_multi_msg"
)

// NewMetadata creates and returns a new ICS27 Metadata instance
func NewMetadata(version, controllerConnectionID, hostConnectionID, accAddress, encoding, txType string) Metadata {
	_ = "STUB: not implemented"
	return *new(Metadata)
}

// IsPreviousMetadataEqual compares a metadata to a previous version string set in a channel struct.
// It ensures all fields are equal except the Address string
func IsPreviousMetadataEqual(previousVersion string, metadata Metadata) bool {
	_ = "STUB: not implemented"
	return false
}

// ValidateControllerMetadata performs validation of the provided ICS27 controller metadata parameters
func ValidateControllerMetadata(ctx sdk.Context, channelKeeper ChannelKeeper, connectionHops []string, metadata Metadata) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateHostMetadata performs validation of the provided ICS27 host metadata parameters
func ValidateHostMetadata(ctx sdk.Context, channelKeeper ChannelKeeper, connectionHops []string, metadata Metadata) error {
	_ = "STUB: not implemented"
	return nil
}

// isSupportedEncoding returns true if the provided encoding is supported, otherwise false
func isSupportedEncoding(encoding string) bool { _ = "STUB: not implemented"; return false }

// getSupportedEncoding returns a string slice of supported encoding formats
func getSupportedEncoding() []string { _ = "STUB: not implemented"; return nil }

// isSupportedTxType returns true if the provided transaction type is supported, otherwise false
func isSupportedTxType(txType string) bool { _ = "STUB: not implemented"; return false }

// getSupportedTxTypes returns a string slice of supported transaction types
func getSupportedTxTypes() []string { _ = "STUB: not implemented"; return nil }

// validateConnectionParams compares the given the controller and host connection IDs to those set in the provided ICS27 Metadata
func validateConnectionParams(metadata Metadata, controllerConnectionID, hostConnectionID string) error {
	_ = "STUB: not implemented"
	return nil
}
