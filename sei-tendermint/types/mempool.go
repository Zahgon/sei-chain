package types

import (
	"fmt"

	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto"
	tmbytes "github.com/sei-protocol/sei-chain/sei-tendermint/libs/bytes"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
)

// TxHash is the fixed length array hash used as an index.
type TxHash crypto.Hash

func (txHash TxHash) Bytes() tmbytes.HexBytes {
	_ = "STUB: not implemented"
	return *new(tmbytes.HexBytes)
}

// ToProto converts Data to protobuf
func (txHash *TxHash) ToProto() *tmproto.TxKey { _ = "STUB: not implemented"; return nil }

func (txHash TxHash) String() string { _ = "STUB: not implemented"; return "" }

func (txHash TxHash) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

// TxHashFromProto takes a protobuf representation of TxHash &
// returns the native type.
func TxHashFromProto(dp *tmproto.TxKey) (TxHash, error) {
	_ = "STUB: not implemented"
	return *new(TxHash), nil
}

func TxHashesListFromProto(dps []*tmproto.TxKey) ([]TxHash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
