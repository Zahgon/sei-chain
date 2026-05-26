package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/transfer/types"
	clienttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/types"
	channeltypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/04-channel/types"
)

// SendTransfer handles transfer sending logic. There are 2 possible cases:
//
// 1. Sender chain is acting as the source zone. The coins are transferred
// to an escrow address (i.e locked) on the sender chain and then transferred
// to the receiving chain through IBC TAO logic. It is expected that the
// receiving chain will mint vouchers to the receiving address.
//
// 2. Sender chain is acting as the sink zone. The coins (vouchers) are burned
// on the sender chain and then transferred to the receiving chain though IBC
// TAO logic. It is expected that the receiving chain, which had previously
// sent the original denomination, will unescrow the fungible token and send
// it to the receiving address.
//
// Another way of thinking of source and sink zones is through the token's
// timeline. Each send to any chain other than the one it was previously
// received from is a movement forwards in the token's timeline. This causes
// trace to be added to the token's history and the destination port and
// destination channel to be prefixed to the denomination. In these instances
// the sender chain is acting as the source zone. When the token is sent back
// to the chain it previously received from, the prefix is removed. This is
// a backwards movement in the token's timeline and the sender chain
// is acting as the sink zone.
//
// Example:
// These steps of transfer occur: A -> B -> C -> A -> C -> B -> A
//
// 1. A -> B : sender chain is source zone. Denom upon receiving: 'B/denom'
// 2. B -> C : sender chain is source zone. Denom upon receiving: 'C/B/denom'
// 3. C -> A : sender chain is source zone. Denom upon receiving: 'A/C/B/denom'
// 4. A -> C : sender chain is sink zone. Denom upon receiving: 'C/B/denom'
// 5. C -> B : sender chain is sink zone. Denom upon receiving: 'B/denom'
// 6. B -> A : sender chain is sink zone. Denom upon receiving: 'denom'
//
// Note: An IBC Transfer must be initiated using a MsgTransfer via the Transfer rpc handler
func (k Keeper) SendTransfer(
	ctx sdk.Context,
	sourcePort,
	sourceChannel string,
	token sdk.Coin,
	sender sdk.AccAddress,
	receiver string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) error {
	_ = "STUB: not implemented"
	return nil
}

// sendTransfer handles transfer sending logic.
func (k Keeper) sendTransfer(
	ctx sdk.Context,
	sourcePort,
	sourceChannel string,
	token sdk.Coin,
	sender sdk.AccAddress,
	receiver string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
	memo string,
) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// get the next sequence

// begin createOutgoingPacket logic
// See spec for this logic: https://github.com/cosmos/ibc/tree/master/spec/app/ics-020-fungible-token-transfer#packet-relay

// NOTE: denomination and hex hash correctness checked during msg.ValidateBasic

// deconstruct the token denomination into the denomination trace info
// to determine if the sender is the source chain

// NOTE: SendTransfer simply sends the denomination as it exists on its own
// chain inside the packet data. The receiving chain will perform denom
// prefixing as necessary.

// create the escrow address for the tokens

// escrow source tokens. It fails if balance insufficient.

// transfer the coins to the module account and burn them

// NOTE: should not happen as the module account was
// retrieved on the step above and it has enough balace
// to burn.

// OnRecvPacket processes a cross chain fungible token transfer. If the
// sender chain is the source of minted tokens then vouchers will be minted
// and sent to the receiving address. Otherwise if the sender chain is sending
// back tokens this chain originally transferred to it, the tokens are
// unescrowed and sent to the receiving address.
func (k Keeper) OnRecvPacket(ctx sdk.Context, packet channeltypes.Packet, data types.FungibleTokenPacketData) error {
	_ = "STUB: not implemented"
	// validate packet data upon receiving
	return nil
}

// decode the receiver address

// parse the transfer amount

// This is the prefix that would have been prefixed to the denomination
// on sender chain IF and only if the token originally came from the
// receiving chain.
//
// NOTE: We use SourcePort and SourceChannel here, because the counterparty
// chain would have prefixed with DestPort and DestChannel when originally
// receiving this coin as seen in the "sender chain is the source" condition.

// sender chain is not the source, unescrow tokens

// remove prefix added by sender chain

// coin denomination used in sending from the escrow address

// The denomination used to send the coins is either the native denom or the hash of the path
// if the denomination is not native.

// unescrow tokens

// NOTE: this error is only expected to occur given an unexpected bug or a malicious
// counterparty module. The bug may occur in bank or any part of the code that allows
// the escrow address to be drained. A malicious counterparty module could drain the
// escrow address by allowing more tokens to be sent back then were escrowed.

// sender chain is the source, mint vouchers

// since SendPacket did not prefix the denomination, we must prefix denomination here

// NOTE: sourcePrefix contains the trailing "/"

// construct the denomination trace from the full raw denomination

// mint new tokens if the source of the transfer is the same chain

// send to receiver

// OnAcknowledgementPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain. If the acknowledgement
// was a success then nothing occurs. If the acknowledgement failed, then
// the sender is refunded their tokens using the refundPacketToken function.
func (k Keeper) OnAcknowledgementPacket(ctx sdk.Context, packet channeltypes.Packet, data types.FungibleTokenPacketData, ack channeltypes.Acknowledgement) error {
	_ = "STUB: not implemented"
	return nil
}

// the acknowledgement succeeded on the receiving chain so nothing
// needs to be executed and no error needs to be returned

// OnTimeoutPacket refunds the sender since the original packet sent was
// never received and has been timed out.
func (k Keeper) OnTimeoutPacket(ctx sdk.Context, packet channeltypes.Packet, data types.FungibleTokenPacketData) error {
	_ = "STUB: not implemented"
	return nil
}

// refundPacketToken will unescrow and send back the tokens back to sender
// if the sending chain was the source chain. Otherwise, the sent tokens
// were burnt in the original send so new tokens are minted and sent to
// the sending address.
func (k Keeper) refundPacketToken(ctx sdk.Context, packet channeltypes.Packet, data types.FungibleTokenPacketData) error {
	_ = "STUB: not implemented"
	// NOTE: packet data type already checked in handler.go
	return nil
}

// parse the denomination from the full denom path

// parse the transfer amount

// decode the sender address

// unescrow tokens back to sender

// NOTE: this error is only expected to occur given an unexpected bug or a malicious
// counterparty module. The bug may occur in bank or any part of the code that allows
// the escrow address to be drained. A malicious counterparty module could drain the
// escrow address by allowing more tokens to be sent back then were escrowed.

// mint vouchers back to sender

// DenomPathFromHash returns the full denomination path prefix from an ibc denom with a hash
// component.
func (k Keeper) DenomPathFromHash(ctx sdk.Context, denom string) (string, error) {
	_ = "STUB: not implemented"
	// trim the denomination prefix, by default "ibc/"
	return "", nil
}
