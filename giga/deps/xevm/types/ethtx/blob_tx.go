package ethtx

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func NewBlobTx(tx *ethtypes.Transaction) (*BlobTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// internally BlobTx uses uint256 which is guaranteed to not have overflow, so using NewIntFromBigInt directly here

func (tx *BlobTx) TxType() uint8 { _ = "STUB: not implemented"; return 0 }

func (tx *BlobTx) Copy() TxData { _ = "STUB: not implemented"; return *new(TxData) }

func (tx *BlobTx) GetChainID() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *BlobTx) GetAccessList() ethtypes.AccessList {
	_ = "STUB: not implemented"
	return *new(ethtypes.AccessList)
}

func (tx *BlobTx) GetData() []byte { _ = "STUB: not implemented"; return nil }

func (tx *BlobTx) GetGas() uint64 { _ = "STUB: not implemented"; return 0 }

func (tx *BlobTx) BlobGas() uint64 { _ = "STUB: not implemented"; return 0 }

func (tx *BlobTx) GetGasPrice() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *BlobTx) GetGasTipCap() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *BlobTx) GetGasFeeCap() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *BlobTx) GetValue() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *BlobTx) GetNonce() uint64 { _ = "STUB: not implemented"; return 0 }

func (tx *BlobTx) GetTo() *common.Address { _ = "STUB: not implemented"; return nil }

func (tx *BlobTx) GetBlobFeeCap() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *BlobTx) GetBlobHashes() []common.Hash { _ = "STUB: not implemented"; return nil }

func (tx *BlobTx) AsEthereumData() ethtypes.TxData {
	_ = "STUB: not implemented"
	return *new(ethtypes.TxData)
}

func (tx *BlobTx) GetRawSignatureValues() (v, r, s *big.Int) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (tx *BlobTx) SetSignatureValues(chainID, v, r, s *big.Int) { _ = "STUB: not implemented"; return }

func (tx BlobTx) Validate() error { _ = "STUB: not implemented"; return nil }

// Amount can be 0

func (tx BlobTx) Fee() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx BlobTx) blobFee() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx BlobTx) Cost() *big.Int { _ = "STUB: not implemented"; return nil }

func (tx *BlobTx) EffectiveGasPrice(baseFee *big.Int) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

func (tx BlobTx) EffectiveFee(baseFee *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

func (tx BlobTx) EffectiveCost(baseFee *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

func sidecarConverter(ethSidecar *ethtypes.BlobTxSidecar) *BlobTxSidecar {
	_ = "STUB: not implemented"
	return nil
}

func sidecarToEthSidecar(sidecar *BlobTxSidecar) *ethtypes.BlobTxSidecar {
	_ = "STUB: not implemented"
	return nil
}
