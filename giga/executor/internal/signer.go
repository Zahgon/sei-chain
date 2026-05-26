package internal

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var _ types.Signer = (*Signer)(nil)

type Signer struct {
	From common.Address
}

func (sig *Signer) Sender(_ *types.Transaction) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

func (sig *Signer) SignatureValues(_ *types.Transaction, _ []byte) (r, s, v *big.Int, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func (sig *Signer) ChainID() *big.Int { _ = "STUB: not implemented"; return nil }

func (sig *Signer) Hash(_ *types.Transaction) common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
}

func (sig *Signer) Equal(_ types.Signer) bool { _ = "STUB: not implemented"; return false }
