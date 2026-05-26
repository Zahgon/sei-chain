package ethtx

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func NewDynamicFeeTx(tx *ethtypes.Transaction) (*DynamicFeeTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tx *DynamicFeeTx) TxType() uint8 { _ = "STUB: not implemented"; return 0 }

func (tx *DynamicFeeTx) Copy() TxData { _ = "STUB: not implemented"; return *new(TxData) }

func (tx *DynamicFeeTx) GetChainID() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *DynamicFeeTx) GetAccessList() ethtypes.AccessList {
	_ = "STUB: not implemented"
	return *new(ethtypes.AccessList)
}

func (tx *DynamicFeeTx) GetData() []byte { _ = "STUB: not implemented"; return nil }

func (tx *DynamicFeeTx) GetGas() uint64 { _ = "STUB: not implemented"; return 0 }

func (tx *DynamicFeeTx) GetGasPrice() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *DynamicFeeTx) GetGasTipCap() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *DynamicFeeTx) GetGasFeeCap() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *DynamicFeeTx) GetValue() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *DynamicFeeTx) GetNonce() uint64 { _ = "STUB: not implemented"; return 0 }

func (tx *DynamicFeeTx) GetTo() *common.Address { _ = "STUB: not implemented"; return nil }

func (tx *DynamicFeeTx) AsEthereumData() ethtypes.TxData {
	_ = "STUB: not implemented"
	return *new(ethtypes.TxData)
}

func (tx *DynamicFeeTx) GetRawSignatureValues() (v, r, s *big.Int) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (tx *DynamicFeeTx) SetSignatureValues(chainID, v, r, s *big.Int) {
	_ = "STUB: not implemented"
	return
}

func (tx *DynamicFeeTx) GetBlobFeeCap() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *DynamicFeeTx) GetBlobHashes() []common.Hash { _ = "STUB: not implemented"; return nil }

func (tx DynamicFeeTx) Validate() error { _ = "STUB: not implemented"; return nil }

// Amount can be 0

func (tx DynamicFeeTx) Fee() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx DynamicFeeTx) Cost() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *DynamicFeeTx) EffectiveGasPrice(baseFee *big.Int) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

func (tx DynamicFeeTx) EffectiveFee(baseFee *big.Int) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

func (tx DynamicFeeTx) EffectiveCost(baseFee *big.Int) *big.Int {
	_ = "STUB: not implemented"
	return nil
}
