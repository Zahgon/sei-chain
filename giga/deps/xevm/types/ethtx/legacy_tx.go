package ethtx

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func NewLegacyTx(tx *ethtypes.Transaction) (*LegacyTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tx *LegacyTx) TxType() uint8 { _ = "STUB: not implemented"; return 0 }

func (tx *LegacyTx) Copy() TxData { _ = "STUB: not implemented"; return *new(TxData) }

// copied from go-etherem/core/types:deriveChainId
func (tx *LegacyTx) GetChainID() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *LegacyTx) GetAccessList() ethtypes.AccessList {
	_ = "STUB: not implemented"
	return *new(ethtypes.AccessList)
}

func (tx *LegacyTx) GetData() []byte { _ = "STUB: not implemented"; return nil }

func (tx *LegacyTx) GetGas() uint64 { _ = "STUB: not implemented"; return 0 }

func (tx *LegacyTx) GetGasPrice() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *LegacyTx) GetGasTipCap() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *LegacyTx) GetGasFeeCap() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *LegacyTx) GetValue() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *LegacyTx) GetNonce() uint64 { _ = "STUB: not implemented"; return 0 }

func (tx *LegacyTx) GetTo() *common.Address { _ = "STUB: not implemented"; return nil }

func (tx *LegacyTx) AsEthereumData() ethtypes.TxData {
	_ = "STUB: not implemented"
	return *new(ethtypes.TxData)
}

func (tx *LegacyTx) GetRawSignatureValues() (v, r, s *big.Int) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (tx *LegacyTx) SetSignatureValues(_, v, r, s *big.Int) { _ = "STUB: not implemented"; return }

func (tx *LegacyTx) GetBlobFeeCap() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *LegacyTx) GetBlobHashes() []common.Hash { _ = "STUB: not implemented"; return nil }

func (tx *LegacyTx) Validate() error { _ = "STUB: not implemented"; return nil }

// Amount can be 0

func (tx LegacyTx) Fee() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx LegacyTx) Cost() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx LegacyTx) EffectiveGasPrice(_ *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

func (tx LegacyTx) EffectiveFee(_ *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

func (tx LegacyTx) EffectiveCost(_ *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }
