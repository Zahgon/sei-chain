package tx

import (
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// MaxGasWanted defines the max gas allowed.
const MaxGasWanted = uint64((1 << 63) - 1)

// Interface implementation checks.
var (
	_, _, _, _ codectypes.UnpackInterfacesMessage = &Tx{}, &TxBody{}, &AuthInfo{}, &SignerInfo{}
	_          sdk.Tx                             = &Tx{}
)

// GetMsgs implements the GetMsgs method on sdk.Tx.
func (t *Tx) GetMsgs() []sdk.Msg { _ = "STUB: not implemented"; return nil }

// ValidateBasic implements the ValidateBasic method on sdk.Tx.
func (t *Tx) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (t *Tx) GetGasEstimate() uint64 { _ = "STUB: not implemented"; return 0 }

// GetSigners retrieves all the signers of a tx.
// This includes all unique signers of the messages (in order),
// as well as the FeePayer (if specified and not already included).
func (t *Tx) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// ensure any specified fee payer is included in the required signers (at the end)

func (t *Tx) GetGas() uint64 { _ = "STUB: not implemented"; return 0 }

func (t *Tx) GetFee() sdk.Coins { _ = "STUB: not implemented"; return *new(sdk.Coins) }

func (t *Tx) FeePayer() sdk.AccAddress { _ = "STUB: not implemented"; return *new(sdk.AccAddress) }

// use first signer as default if no payer specified

func (t *Tx) FeeGranter() sdk.AccAddress { _ = "STUB: not implemented"; return *new(sdk.AccAddress) }

// UnpackInterfaces implements the UnpackInterfaceMessages.UnpackInterfaces method
func (t *Tx) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements the UnpackInterfaceMessages.UnpackInterfaces method
func (m *TxBody) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements the UnpackInterfaceMessages.UnpackInterfaces method
func (m *AuthInfo) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements the UnpackInterfaceMessages.UnpackInterfaces method
func (m *SignerInfo) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterInterfaces registers the sdk.Tx interface.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) { _ = "STUB: not implemented"; return }
