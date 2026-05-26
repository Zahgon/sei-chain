package keeper

import (
	"context"
	"errors"

	"github.com/sei-protocol/sei-chain/x/evm/types"
)

var _ types.QueryServer = Querier{}

var ErrMustSpecifyPointer = errors.New("must specify a pointer")
var ErrMustSpecifyPointee = errors.New("must specify a pointee")

// Querier defines a wrapper around the x/mint keeper providing gRPC method
// handlers.
type Querier struct {
	*Keeper
}

func NewQuerier(k *Keeper) Querier { _ = "STUB: not implemented"; return *new(Querier) }

func (q Querier) SeiAddressByEVMAddress(c context.Context, req *types.QuerySeiAddressByEVMAddressRequest) (*types.QuerySeiAddressByEVMAddressResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q Querier) EVMAddressBySeiAddress(c context.Context, req *types.QueryEVMAddressBySeiAddressRequest) (*types.QueryEVMAddressBySeiAddressResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q Querier) StaticCall(c context.Context, req *types.QueryStaticCallRequest) (*types.QueryStaticCallResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q Querier) Pointer(c context.Context, req *types.QueryPointerRequest) (*types.QueryPointerResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q Querier) PointerVersion(c context.Context, req *types.QueryPointerVersionRequest) (*types.QueryPointerVersionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q Querier) Pointee(c context.Context, req *types.QueryPointeeRequest) (*types.QueryPointeeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
