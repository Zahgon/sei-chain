package api

import (
	"github.com/sei-protocol/sei-chain/sei-wasmvm/types"
)

/***** Mock types.GoAPI ****/

func MockFailureCanonicalAddress(human string) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func MockFailureHumanAddress(canon []byte) (string, uint64, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func NewMockFailureAPI() *types.GoAPI { _ = "STUB: not implemented"; return nil }
