package server

import (
	"aryaframe/generate/pb/bkgrpc"
	grpc2 "aryaframe/internal/server/grpc"
	"aryaframe/service"
	"context"
	"github.com/go-kit/kit/transport/grpc"
	"reflect"
)

func init() {
	svr := service.NewStringServicesImpl()
	endpoints := NewStringServicesEndpoint(svr)
	grpc2.RegisterService(
		reflect.TypeOf(bkgrpc.RegisterStringServicesServer),
		reflect.ValueOf(bkgrpc.RegisterStringServicesServer),
		[]reflect.Value{reflect.ValueOf(NewRouter(endpoints))},
	)
}

func NewRouter(endpoints StringServicesEndpoint) service.StringServices {
	return &service.StringServicesA{
		Concat_:      grpc.NewServer(endpoints.ConcatEndpoint, DecodeStringRequest, EncodeStringResponse),
		Diff_:        grpc.NewServer(endpoints.DiffEndpoint, DecodeStringRequest, EncodeStringResponse),
		HealtStatus_: grpc.NewServer(endpoints.HealtStatusEndpoint, DecodeHealthRequest, EncodeHealthResponse),
	}
}

func DecodeStringRequest(ctx context.Context, req interface{}) (interface{}, error) {
	rsp := req.(*bkgrpc.StringRequest)
	return rsp, nil
}

func EncodeStringResponse(ctx context.Context, req interface{}) (interface{}, error) {
	rsp := req.(*bkgrpc.StringResponse)
	return rsp, nil
}

func DecodeHealthRequest(ctx context.Context, req interface{}) (interface{}, error) {
	rsp := req.(*bkgrpc.HealthRequest)
	return rsp, nil
}

func EncodeHealthResponse(ctx context.Context, req interface{}) (interface{}, error) {
	rsp := req.(*bkgrpc.HealthResponse)
	return rsp, nil
}
