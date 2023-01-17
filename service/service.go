package service

import (
	"aryaframe/generate/pb/bkgrpc"
	"context"
	"github.com/go-kit/kit/transport/grpc"
)

type StringServices interface {
	Concat(ctx context.Context, request *bkgrpc.StringRequest) (*bkgrpc.StringResponse, error)
	Diff(ctx context.Context, request *bkgrpc.StringRequest) (*bkgrpc.StringResponse, error)
	HealtStatus(ctx context.Context, request *bkgrpc.HealthRequest) (*bkgrpc.HealthResponse, error)
}

type StringServicesA struct {
	bkgrpc.UnimplementedStringServicesServer
	Concat_      grpc.Handler
	Diff_        grpc.Handler
	HealtStatus_ grpc.Handler
}

func (s StringServicesA) Concat(ctx context.Context, req *bkgrpc.StringRequest) (*bkgrpc.StringResponse, error) {
	_, rep, err := s.Concat_.ServeGRPC(ctx, req)
	return rep.(*bkgrpc.StringResponse), err
}

func (s StringServicesA) Diff(ctx context.Context, req *bkgrpc.StringRequest) (*bkgrpc.StringResponse, error) {
	_, rep, err := s.Diff_.ServeGRPC(ctx, req)
	return rep.(*bkgrpc.StringResponse), err
}

func (s StringServicesA) HealtStatus(ctx context.Context, req *bkgrpc.HealthRequest) (*bkgrpc.HealthResponse, error) {
	_, rep, err := s.HealtStatus_.ServeGRPC(ctx, req)
	return rep.(*bkgrpc.HealthResponse), err
}
