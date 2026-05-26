package admin

import (
	"context"

	"github.com/sei-protocol/sei-chain/admin/types"
)

type service struct {
	types.UnimplementedAdminServiceServer
}

func (s *service) SetLogLevel(_ context.Context, req *types.SetLogLevelRequest) (*types.SetLogLevelResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec

func (s *service) GetLogLevel(_ context.Context, req *types.GetLogLevelRequest) (*types.GetLogLevelResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *service) ListLoggers(_ context.Context, req *types.ListLoggersRequest) (*types.ListLoggersResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
