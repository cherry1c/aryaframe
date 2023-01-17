package service

import (
	"aryaframe/generate/pb/bkgrpc"
	"context"
)

type StringServices interface {
	Concat(ctx context.Context, request *bkgrpc.StringRequest, response *bkgrpc.StringResponse) error
	Diff(ctx context.Context, request *bkgrpc.StringRequest, response *bkgrpc.StringResponse) error
	HealtStatus(ctx context.Context, request *bkgrpc.HealthRequest, response *bkgrpc.HealthResponse) error
}
