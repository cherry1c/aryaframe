package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"time"
)

type Server struct {
	address      string
	network      string
	readTimeout  time.Duration
	writeTimeout time.Duration
}

func NewServer(opts ...Option) *Server {
	s := &Server{
		network: "tcp",
	}
	for _, op := range opts {
		op(s)
	}

	return s
}

func (s *Server) Serve(ctx context.Context) error {
	lis, err := net.Listen(s.network, s.address)
	if err != nil {
		fmt.Printf("grpc listen failed err: %s\n", err.Error())
		return err
	}
	fmt.Printf("start to server network %s port %s\n", s.network, s.address)
	svr := grpc.NewServer()
	registerServer(svr)
	if err = svr.Serve(lis); err != nil {
		fmt.Printf("start grpc failed err: %s\n", err.Error())
		return err
	}
	fmt.Printf("start to server successful\n")
	select {
	case <-ctx.Done():
		return ctx.Err()
	}
}
func (s *Server) Stop() error {
	return nil
}
func (s *Server) GracefulStop(ctx context.Context) error {
	return nil
}
