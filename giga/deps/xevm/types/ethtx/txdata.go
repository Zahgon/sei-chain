package ethtx

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/gogo/protobuf/proto"
)

var (
	_ TxData = &LegacyTx{}
	_ TxData = &AccessListTx{}
	_ TxData = &DynamicFeeTx{}
	_ TxData = &BlobTx{}
	_ TxData = &AssociateTx{}
	_ TxData = &SetCodeTx{}
)

// Unfortunately `TxData` interface in go-ethereum/core/types defines its functions
// as private, so we have to define our own here.
type TxData interface {
	proto.Message
	TxType() byte
	Copy() TxData
	GetChainID() *big.Int
	GetAccessList() ethtypes.AccessList
	GetData() []byte
	GetNonce() uint64
	GetGas() uint64
	GetGasPrice() *big.Int
	GetGasTipCap() *big.Int
	GetGasFeeCap() *big.Int
	GetValue() *big.Int
	GetTo() *common.Address

	GetRawSignatureValues() (v, r, s *big.Int)
	SetSignatureValues(chainID, v, r, s *big.Int)

	AsEthereumData() ethtypes.TxData
	Validate() error

	Fee() *big.Int
	Cost() *big.Int

	EffectiveGasPrice(baseFee *big.Int) *big.Int
	EffectiveFee(baseFee *big.Int) *big.Int
	EffectiveCost(baseFee *big.Int) *big.Int

	GetBlobHashes() []common.Hash
	GetBlobFeeCap() *big.Int
}

func NewTxDataFromTx(tx *ethtypes.Transaction) (TxData, error) {
	_ = "STUB: not implemented"
	return *new(TxData), nil
}

func rawSignatureValues(vBz, rBz, sBz []byte) (v, r, s *big.Int) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// fee = gas limit * gas price
func fee(gasPrice *big.Int, gas uint64) *big.Int { _ = "STUB: not implemented"; return nil }

// cost = fee + tokens to send
func cost(fee, value *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }
