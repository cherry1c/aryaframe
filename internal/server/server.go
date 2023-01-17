package server

import (
	"aryaframe/internal/conf"
	svrConf "aryaframe/internal/server/conf"
	grpc2 "aryaframe/internal/server/grpc"
	"context"
	"fmt"
)

type Server interface {
	Serve(ctx context.Context) error
	Stop() error
	GracefulStop(ctx context.Context) error
}

const serverConfigPrefix = "aryaframe.server"

var Servers []Server

func Init() error {
	if err := initServer(); err != nil {
		return err
	}

	if err := initService(); err != nil {
		return err
	}

	return nil
}

func initServer() error {
	// load server config
	sConfig := svrConf.NewDefaultConfig()
	if err := conf.UnmarshalKey(serverConfigPrefix, sConfig); err != nil {
		fmt.Printf("unmarshal server config failed err: %s\n", err.Error())
		return err
	}
	var grpcOption []grpc2.Option
	grpcOption = append(grpcOption, grpc2.WithAddress(sConfig.Grpc.Address))
	grpcOption = append(grpcOption, grpc2.WithReadTimeout(sConfig.ReadTimeout))
	grpcOption = append(grpcOption, grpc2.WithWriteTimeout(sConfig.WriteTimeout))

	// init rpc server
	srv := grpc2.NewServer(grpcOption...)

	// init http server

	Servers = append(Servers, srv)
	return nil
}

func initService() error {
	// init trace and flow control
	return nil
}
