package service

import (
	"aryaframe/generate/pb/bkgrpc"
	"context"
)

type stringServicesImpl struct {
}

func NewStringServicesImpl() StringServices {
	return stringServicesImpl{}
}

func (e stringServicesImpl) Concat(ctx context.Context, request *bkgrpc.StringRequest) (*bkgrpc.StringResponse, error) {
	panic("implement me")
}
func (e stringServicesImpl) Diff(ctx context.Context, request *bkgrpc.StringRequest) (*bkgrpc.StringResponse, error) {
	panic("implement me")
}
func (e stringServicesImpl) HealtStatus(ctx context.Context, request *bkgrpc.HealthRequest) (*bkgrpc.HealthResponse, error) {
	panic("implement me")
}
