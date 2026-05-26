package ethtx

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func NewAccessListTx(tx *ethtypes.Transaction) (*AccessListTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tx *AccessListTx) TxType() uint8 { _ = "STUB: not implemented"; return 0 }

func (tx *AccessListTx) Copy() TxData { _ = "STUB: not implemented"; return *new(TxData) }

func (tx *AccessListTx) GetChainID() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *AccessListTx) GetAccessList() ethtypes.AccessList {
	_ = "STUB: not implemented"
	return *new(ethtypes.AccessList)
}

func (tx *AccessListTx) GetData() []byte { _ = "STUB: not implemented"; return nil }

func (tx *AccessListTx) GetGas() uint64 { _ = "STUB: not implemented"; return 0 }

func (tx *AccessListTx) GetGasPrice() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *AccessListTx) GetGasTipCap() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *AccessListTx) GetGasFeeCap() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *AccessListTx) GetValue() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *AccessListTx) GetNonce() uint64 { _ = "STUB: not implemented"; return 0 }

func (tx *AccessListTx) GetTo() *common.Address { _ = "STUB: not implemented"; return nil }

func (tx *AccessListTx) AsEthereumData() ethtypes.TxData {
	_ = "STUB: not implemented"
	return *new(ethtypes.TxData)
}

func (tx *AccessListTx) GetRawSignatureValues() (v, r, s *big.Int) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (tx *AccessListTx) SetSignatureValues(chainID, v, r, s *big.Int) {
	_ = "STUB: not implemented"
	return
}

func (tx *AccessListTx) GetBlobFeeCap() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *AccessListTx) GetBlobHashes() []common.Hash { _ = "STUB: not implemented"; return nil }

func (tx AccessListTx) Validate() error { _ = "STUB: not implemented"; return nil }

// Amount can be 0

func (tx AccessListTx) Fee() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx AccessListTx) Cost() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx AccessListTx) EffectiveGasPrice(_ *big.Int) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

func (tx AccessListTx) EffectiveFee(_ *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

func (tx AccessListTx) EffectiveCost(_ *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }
