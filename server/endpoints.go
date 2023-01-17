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

	var nConcatEndpoint endpoint.Endpoint
	nConcatEndpoint = genDiffEndpoint(svc)

	var nConcatEndpoint endpoint.Endpoint
	nConcatEndpoint = genHealtStatusEndpoint(svc)

	return StringServicesEndpoint{
		ConcatEndpoint: nConcatEndpoint,
	}
}

func (s StringServicesEndpoint) Concat(ctx context.Context, request *bkgrpc.StringRequest, response *bkgrpc.StringResponse) error {
	rsp, err := s.ConcatEndpoint(ctx, request)
	*response = *(rsp.(*bkgrpc.StringResponse))
	return err
}

func (s StringServicesEndpoint) Diff(ctx context.Context, request *bkgrpc.StringRequest, response *bkgrpc.StringResponse) error {
	rsp, err := s.DiffEndpoint(ctx, request)
	*response = *(rsp.(*bkgrpc.StringResponse))
	return err
}

func (s StringServicesEndpoint) HealtStatus(ctx context.Context, request *bkgrpc.HealthRequest, response *bkgrpc.HealthResponse) error {
	rsp, err := s.HealtStatusEndpoint(ctx, request)
	*response = *(rsp.(*bkgrpc.HealthResponse))
	return err
}

func genConcatEndpoint(service service.StringServices) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*bkgrpc.StringRequest)
		rsp := &bkgrpc.StringResponse{}
		err = service.Concat(ctx, req, rsp)
		return rsp, err
	}
}

func genDiffEndpoint(service service.StringServices) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*bkgrpc.StringRequest)
		rsp := &bkgrpc.StringResponse{}
		err = service.Diff(ctx, req, rsp)
		return rsp, err
	}
}

func genHealtStatusEndpoint(service service.StringServices) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*bkgrpc.HealthRequest)
		rsp := &bkgrpc.HealthResponse{}
		err = service.HealtStatus(ctx, req, rsp)
		return rsp, err
	}
}
