package ethtx

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func NewAssociateTx(tx *ethtypes.Transaction, customMessage string) (*AssociateTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tx *AssociateTx) TxType() byte         { _ = "STUB: not implemented"; return 0 }
func (tx *AssociateTx) Copy() TxData         { _ = "STUB: not implemented"; return *new(TxData) }
func (tx *AssociateTx) GetChainID() *big.Int { _ = "STUB: not implemented"; return nil }
func (tx *AssociateTx) GetAccessList() ethtypes.AccessList {
	_ = "STUB: not implemented"
	return *new(ethtypes.AccessList)
}
func (tx *AssociateTx) GetData() []byte        { _ = "STUB: not implemented"; return nil }
func (tx *AssociateTx) GetNonce() uint64       { _ = "STUB: not implemented"; return 0 }
func (tx *AssociateTx) GetGas() uint64         { _ = "STUB: not implemented"; return 0 }
func (tx *AssociateTx) GetGasPrice() *big.Int  { _ = "STUB: not implemented"; return nil }
func (tx *AssociateTx) GetGasTipCap() *big.Int { _ = "STUB: not implemented"; return nil }
func (tx *AssociateTx) GetGasFeeCap() *big.Int { _ = "STUB: not implemented"; return nil }
func (tx *AssociateTx) GetValue() *big.Int     { _ = "STUB: not implemented"; return nil }
func (tx *AssociateTx) GetTo() *common.Address { _ = "STUB: not implemented"; return nil }

func (tx *AssociateTx) GetRawSignatureValues() (v, r, s *big.Int) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (tx *AssociateTx) SetSignatureValues(_, _, _, _ *big.Int) { _ = "STUB: not implemented"; return }

func (tx *AssociateTx) AsEthereumData() ethtypes.TxData {
	_ = "STUB: not implemented"
	return *new(ethtypes.TxData)
}
func (tx *AssociateTx) Validate() error { _ = "STUB: not implemented"; return nil }

func (tx *AssociateTx) Fee() *big.Int  { _ = "STUB: not implemented"; return nil }
func (tx *AssociateTx) Cost() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *AssociateTx) EffectiveGasPrice(_ *big.Int) *big.Int {
	_ = "STUB: not implemented"
	return nil
}
func (tx *AssociateTx) EffectiveFee(_ *big.Int) *big.Int  { _ = "STUB: not implemented"; return nil }
func (tx *AssociateTx) EffectiveCost(_ *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *AssociateTx) GetBlobHashes() []common.Hash { _ = "STUB: not implemented"; return nil }
func (tx *AssociateTx) GetBlobFeeCap() *big.Int      { _ = "STUB: not implemented"; return nil }
