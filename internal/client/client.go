package client

import (
	cliConf "aryaframe/internal/client/conf"
	"aryaframe/internal/client/grpc"
	"aryaframe/internal/conf"
	"fmt"
)

const clientConfigPrefix = "aryaframe.client"

func Init() error {
	if err := conf.UnmarshalKey(clientConfigPrefix, &cliConf.Configs); err != nil {
		fmt.Printf("unmarshal client config failed err: %s\n", err.Error())
		return err
	}

	for serviceName, cli := range cliConf.Configs {
		var opts []grpc.Option
		opts = append(opts, grpc.WithAddress(cli.Address))
		opts = append(opts, grpc.WithRequestTimeout(cli.RequestTimeout))

		if cli.Protocol == "grpc" {
			grpc.SetClientOptions(serviceName, opts...)
		} else {
			continue
		}
	}

	if err := grpc.Start(); err != nil {
		fmt.Printf("start grpc client failed err: %s\n", err.Error())
		return err
	}

	return nil
}
