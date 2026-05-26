package types

import (
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

// message types for the IBC client
const (
	TypeMsgCreateClient       string = "create_client"
	TypeMsgUpdateClient       string = "update_client"
	TypeMsgUpgradeClient      string = "upgrade_client"
	TypeMsgSubmitMisbehaviour string = "submit_misbehaviour"
)

var (
	_ sdk.Msg = &MsgCreateClient{}
	_ sdk.Msg = &MsgUpdateClient{}
	_ sdk.Msg = &MsgSubmitMisbehaviour{}
	_ sdk.Msg = &MsgUpgradeClient{}

	_ codectypes.UnpackInterfacesMessage = MsgCreateClient{}
	_ codectypes.UnpackInterfacesMessage = MsgUpdateClient{}
	_ codectypes.UnpackInterfacesMessage = MsgSubmitMisbehaviour{}
	_ codectypes.UnpackInterfacesMessage = MsgUpgradeClient{}
)

// NewMsgCreateClient creates a new MsgCreateClient instance
func NewMsgCreateClient(
	clientState exported.ClientState, consensusState exported.ConsensusState, signer string,
) (*MsgCreateClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgCreateClient) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// GetSigners implements sdk.Msg
func (msg MsgCreateClient) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (msg MsgCreateClient) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// NewMsgUpdateClient creates a new MsgUpdateClient instance
func NewMsgUpdateClient(id string, header exported.Header, signer string) (*MsgUpdateClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgUpdateClient) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// GetSigners implements sdk.Msg
func (msg MsgUpdateClient) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (msg MsgUpdateClient) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// NewMsgUpgradeClient creates a new MsgUpgradeClient instance

func NewMsgUpgradeClient(clientID string, clientState exported.ClientState, consState exported.ConsensusState,
	proofUpgradeClient, proofUpgradeConsState []byte, signer string,
) (*MsgUpgradeClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgUpgradeClient) ValidateBasic() error {
	_ = "STUB: not implemented"
	// will not validate client state as committed client may not form a valid client state.
	// client implementations are responsible for ensuring final upgraded client is valid.
	return nil
}

// will not validate consensus state here since the trusted kernel may not form a valid consenus state.
// client implementations are responsible for ensuring client can submit new headers against this consensus state.

// GetSigners implements sdk.Msg
func (msg MsgUpgradeClient) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (msg MsgUpgradeClient) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// NewMsgSubmitMisbehaviour creates a new MsgSubmitMisbehaviour instance.
func NewMsgSubmitMisbehaviour(clientID string, misbehaviour exported.Misbehaviour, signer string) (*MsgSubmitMisbehaviour, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ValidateBasic performs basic (non-state-dependant) validation on a MsgSubmitMisbehaviour.
func (msg MsgSubmitMisbehaviour) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// GetSigners returns the single expected signer for a MsgSubmitMisbehaviour.
func (msg MsgSubmitMisbehaviour) GetSigners() []sdk.AccAddress {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (msg MsgSubmitMisbehaviour) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}
