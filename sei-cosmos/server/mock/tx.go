// nolint
package mock

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// An sdk.Tx which is its own sdk.Msg.
type kvstoreTx struct {
	key   []byte
	value []byte
	bytes []byte
}

// dummy implementation of proto.Message
func (msg kvstoreTx) Reset()         { _ = "STUB: not implemented"; return }
func (msg kvstoreTx) String() string { _ = "STUB: not implemented"; return "" }
func (msg kvstoreTx) ProtoMessage()  { _ = "STUB: not implemented"; return }

var _ sdk.Tx = kvstoreTx{}
var _ sdk.Msg = kvstoreTx{}

func NewTx(key, value string) kvstoreTx { _ = "STUB: not implemented"; return *new(kvstoreTx) }

func (tx kvstoreTx) Route() string { _ = "STUB: not implemented"; return "" }

func (tx kvstoreTx) Type() string { _ = "STUB: not implemented"; return "" }

func (tx kvstoreTx) GetMsgs() []sdk.Msg { _ = "STUB: not implemented"; return nil }

func (tx kvstoreTx) GetMemo() string { _ = "STUB: not implemented"; return "" }

func (tx kvstoreTx) GetSignBytes() []byte {
	_ = "STUB: not implemented"

	// Should the app be calling this? Or only handlers?
	return nil
}

func (tx kvstoreTx) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (tx kvstoreTx) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

func (tx kvstoreTx) GetGasEstimate() uint64 {
	_ = "STUB: not implemented"

	// takes raw transaction bytes and decodes them into an sdk.Tx. An sdk.Tx has
	// all the signatures and can be used to authenticate.
	return 0
}

func decodeTx(txBytes []byte) (sdk.Tx, error) { _ = "STUB: not implemented"; return *new(sdk.Tx), nil }
