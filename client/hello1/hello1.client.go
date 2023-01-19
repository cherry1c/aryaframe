package hello1

import (
	"aryaframe/generate/pb/bkgrpc"
	"aryaframe/internal/client/grpc"
	"context"
	"reflect"
)

type StringServices struct {
}

func init() {
	grpc.RegisterService("stringservices", reflect.ValueOf(bkgrpc.NewStringServicesClient))
}

func (T StringServices) Concat(ctx context.Context, request *bkgrpc.StringRequest) (*bkgrpc.StringResponse, error) {
	rsp, err := grpc.Call("stringservices", "Concat", ctx, request)
	return rsp.(*bkgrpc.StringResponse), err
}

func (T StringServices) Diff(ctx context.Context, request *bkgrpc.StringRequest) (*bkgrpc.StringResponse, error) {
	rsp, err := grpc.Call("stringservices", "Diff", ctx, request)
	return rsp.(*bkgrpc.StringResponse), err
}

func (T StringServices) HealtStatus(ctx context.Context, request *bkgrpc.HealthRequest) (*bkgrpc.HealthResponse, error) {
	rsp, err := grpc.Call("stringservices", "HealtStatus", ctx, request)
	return rsp.(*bkgrpc.HealthResponse), err
}
