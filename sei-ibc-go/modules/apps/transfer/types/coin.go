package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// SenderChainIsSource returns false if the denomination originally came
// from the receiving chain and true otherwise.
func SenderChainIsSource(sourcePort, sourceChannel, denom string) bool {
	_ = "STUB: not implemented"
	// This is the prefix that would have been prefixed to the denomination
	// on sender chain IF and only if the token originally came from the
	// receiving chain.
	return false
}

// ReceiverChainIsSource returns true if the denomination originally came
// from the receiving chain and false otherwise.
func ReceiverChainIsSource(sourcePort, sourceChannel, denom string) bool {
	_ = "STUB: not implemented"
	// The prefix passed in should contain the SourcePort and SourceChannel.
	// If  the receiver chain originally sent the token to the sender chain
	// the denom will have the sender's SourcePort and SourceChannel as the
	// prefix.
	return false
}

// GetDenomPrefix returns the receiving denomination prefix
func GetDenomPrefix(portID, channelID string) string { _ = "STUB: not implemented"; return "" }

// GetPrefixedDenom returns the denomination with the portID and channelID prefixed
func GetPrefixedDenom(portID, channelID, baseDenom string) string {
	_ = "STUB: not implemented"
	return ""
}

// GetTransferCoin creates a transfer coin with the port ID and channel ID
// prefixed to the base denom.
func GetTransferCoin(portID, channelID, baseDenom string, amount sdk.Int) sdk.Coin {
	_ = "STUB: not implemented"
	return *new(sdk.Coin)
}
