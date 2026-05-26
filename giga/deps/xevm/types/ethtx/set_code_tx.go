package ethtx

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/holiman/uint256"
)

func NewSetCodeTx(tx *ethtypes.Transaction) (*SetCodeTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tx *SetCodeTx) TxType() uint8 { _ = "STUB: not implemented"; return 0 }

func (tx *SetCodeTx) Copy() TxData { _ = "STUB: not implemented"; return *new(TxData) }

func (tx *SetCodeTx) GetChainID() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *SetCodeTx) GetAccessList() ethtypes.AccessList {
	_ = "STUB: not implemented"
	return *new(ethtypes.AccessList)
}

func (tx *SetCodeTx) GetAuthList() []ethtypes.SetCodeAuthorization {
	_ = "STUB: not implemented"
	return nil
}

func (tx *SetCodeTx) GetData() []byte { _ = "STUB: not implemented"; return nil }

func (tx *SetCodeTx) GetGas() uint64 { _ = "STUB: not implemented"; return 0 }

func (tx *SetCodeTx) GetGasPrice() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *SetCodeTx) GetGasTipCap() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *SetCodeTx) GetGasFeeCap() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *SetCodeTx) GetValue() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *SetCodeTx) GetNonce() uint64 { _ = "STUB: not implemented"; return 0 }

func (tx *SetCodeTx) GetTo() *common.Address { _ = "STUB: not implemented"; return nil }

func (tx *SetCodeTx) AsEthereumData() ethtypes.TxData {
	_ = "STUB: not implemented"
	return *new(ethtypes.TxData)
}

func bigToUint256(b *big.Int) *uint256.Int { _ = "STUB: not implemented"; return nil }

func (tx *SetCodeTx) GetRawSignatureValues() (v, r, s *big.Int) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (tx *SetCodeTx) SetSignatureValues(chainID, v, r, s *big.Int) {
	_ = "STUB: not implemented"
	return
}

func (tx *SetCodeTx) GetBlobFeeCap() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *SetCodeTx) GetBlobHashes() []common.Hash { _ = "STUB: not implemented"; return nil }

func (tx SetCodeTx) Validate() error { _ = "STUB: not implemented"; return nil }

// Amount can be 0

func (tx SetCodeTx) Fee() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx SetCodeTx) Cost() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *SetCodeTx) EffectiveGasPrice(baseFee *big.Int) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

func (tx SetCodeTx) EffectiveFee(baseFee *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

func (tx SetCodeTx) EffectiveCost(baseFee *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }
