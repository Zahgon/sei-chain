package errors

import (
	"sync"

	"github.com/coinbase/rosetta-sdk-go/types"
)

type errorRegistry struct {
	mu     *sync.RWMutex
	sealed bool
	errors map[int32]*types.Error
}

func (r *errorRegistry) add(err *Error) { _ = "STUB: not implemented"; return }

func (r errorRegistry) list() []*types.Error { _ = "STUB: not implemented"; return nil }

func (r *errorRegistry) seal() { _ = "STUB: not implemented"; return }

var registry = errorRegistry{
	mu:     new(sync.RWMutex),
	errors: make(map[int32]*types.Error),
}
