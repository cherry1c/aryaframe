package grpc

import "time"

type Option func(server *Server)

func WithAddress(address string) Option {
	return func(s *Server) {
		s.address = address
	}
}

func WithReadTimeout(v time.Duration) Option {
	return func(s *Server) {
		s.readTimeout = v
	}
}

func WithWriteTimeout(v time.Duration) Option {
	return func(s *Server) {
		s.writeTimeout = v
	}
}
