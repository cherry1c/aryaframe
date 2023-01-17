package server

import (
	"aryaframe/generate/pb/bkgrpc"
	"aryaframe/service"
	"context"
	"github.com/go-kit/kit/endpoint"
)

type StringServicesEndpoint struct {
	ConcatEndpoint      endpoint.Endpoint
	DiffEndpoint        endpoint.Endpoint
	HealtStatusEndpoint endpoint.Endpoint
}

func NewStringServicesEndpoint(svc service.StringServices) StringServicesEndpoint {
	var nConcatEndpoint endpoint.Endpoint
	nConcatEndpoint = genConcatEndpoint(svc)

	var nDiffEndpoint endpoint.Endpoint
	nDiffEndpoint = genDiffEndpoint(svc)

	var nHealtStatusEndpoint endpoint.Endpoint
	nHealtStatusEndpoint = genHealtStatusEndpoint(svc)

	return StringServicesEndpoint{
		ConcatEndpoint:      nConcatEndpoint,
		DiffEndpoint:        nDiffEndpoint,
		HealtStatusEndpoint: nHealtStatusEndpoint,
	}
}

func (s StringServicesEndpoint) Concat(ctx context.Context, request *bkgrpc.StringRequest) (*bkgrpc.StringResponse, error) {
	rsp, err := s.ConcatEndpoint(ctx, request)
	return rsp.(*bkgrpc.StringResponse), err
}

func (s StringServicesEndpoint) Diff(ctx context.Context, request *bkgrpc.StringRequest) (*bkgrpc.StringResponse, error) {
	rsp, err := s.DiffEndpoint(ctx, request)
	return rsp.(*bkgrpc.StringResponse), err
}

func (s StringServicesEndpoint) HealtStatus(ctx context.Context, request *bkgrpc.HealthRequest) (*bkgrpc.HealthResponse, error) {
	rsp, err := s.HealtStatusEndpoint(ctx, request)
	return rsp.(*bkgrpc.HealthResponse), err
}

func genConcatEndpoint(service service.StringServices) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*bkgrpc.StringRequest)
		return service.Concat(ctx, req)
	}
}

func genDiffEndpoint(service service.StringServices) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*bkgrpc.StringRequest)
		return service.Diff(ctx, req)
	}
}

func genHealtStatusEndpoint(service service.StringServices) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*bkgrpc.HealthRequest)
		return service.HealtStatus(ctx, req)
	}
}
