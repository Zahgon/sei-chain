package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// bank module event types
const (
	EventTypeTransfer    = "transfer"
	EventTypeWeiTransfer = "wei_transfer"

	AttributeKeyRecipient = "recipient"
	AttributeKeySender    = "sender"

	AttributeValueCategory = ModuleName

	// supply and balance tracking events name and attributes
	EventTypeCoinSpent    = "coin_spent"
	EventTypeCoinReceived = "coin_received"
	EventTypeWeiSpent     = "wei_spent"
	EventTypeWeiReceived  = "wei_received"
	EventTypeCoinMint     = "coinbase" // NOTE(fdymylja): using mint clashes with mint module event
	EventTypeCoinBurn     = "burn"

	AttributeKeySpender  = "spender"
	AttributeKeyReceiver = "receiver"
	AttributeKeyMinter   = "minter"
	AttributeKeyBurner   = "burner"
)

// NewCoinSpentEvent constructs a new coin spent sdk.Event

func NewCoinSpentEvent(spender sdk.AccAddress, amount sdk.Coins) sdk.Event {
	_ = "STUB: not implemented"
	return *new(sdk.Event)
}

// NewCoinReceivedEvent constructs a new coin received sdk.Event

func NewCoinReceivedEvent(receiver sdk.AccAddress, amount sdk.Coins) sdk.Event {
	_ = "STUB: not implemented"
	return *new(sdk.Event)
}

// NewWeiSpentEvent constructs a new wei spent sdk.Event

func NewWeiSpentEvent(spender sdk.AccAddress, amount sdk.Int) sdk.Event {
	_ = "STUB: not implemented"
	return *new(sdk.Event)
}

// NewWeiReceivedEvent constructs a new wei received sdk.Event

func NewWeiReceivedEvent(receiver sdk.AccAddress, amount sdk.Int) sdk.Event {
	_ = "STUB: not implemented"
	return *new(sdk.Event)
}

// NewCoinMintEvent construct a new coin minted sdk.Event

func NewCoinMintEvent(minter sdk.AccAddress, amount sdk.Coins) sdk.Event {
	_ = "STUB: not implemented"
	return *new(sdk.Event)
}

// NewCoinBurnEvent constructs a new coin burned sdk.Event

func NewCoinBurnEvent(burner sdk.AccAddress, amount sdk.Coins) sdk.Event {
	_ = "STUB: not implemented"
	return *new(sdk.Event)
}
