package types

import (
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/gogo/protobuf/proto"
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/evm/types/ethtx"
)

const TypeMsgEVMTransaction = "evm_transaction"

var (
	_ sdk.Msg                            = &MsgEVMTransaction{}
	_ codectypes.UnpackInterfacesMessage = &MsgEVMTransaction{}
	_ sdk.ResultDecorator                = &MsgEVMTransactionResponse{}
)

func NewMsgEVMTransaction(txData proto.Message) (*MsgEVMTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (msg *MsgEVMTransaction) Route() string { _ = "STUB: not implemented"; return "" }

func (msg *MsgEVMTransaction) Type() string { _ = "STUB: not implemented"; return "" }

func (msg *MsgEVMTransaction) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

func (msg *MsgEVMTransaction) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

func (msg *MsgEVMTransaction) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (msg *MsgEVMTransaction) AsTransaction() (*ethtypes.Transaction, ethtx.TxData) {
	_ = "STUB: not implemented"
	return nil, *new(ethtx.TxData)
}

// UnpackInterfaces implements UnpackInterfacesMesssage.UnpackInterfaces
func (msg *MsgEVMTransaction) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

func (msg *MsgEVMTransaction) IsAssociateTx() bool { _ = "STUB: not implemented"; return false }

func (msg *MsgEVMTransaction) GetAssociateTx() (*ethtx.AssociateTx, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// should never happen

func MustGetEVMTransactionMessage(tx sdk.Tx) *MsgEVMTransaction {
	_ = "STUB: not implemented"
	return nil
}

func GetEVMTransactionMessage(tx sdk.Tx) *MsgEVMTransaction { _ = "STUB: not implemented"; return nil }

func (res *MsgEVMTransactionResponse) DecorateSdkResult(sdkRes *sdk.Result) {
	_ = "STUB: not implemented"
	return
}
