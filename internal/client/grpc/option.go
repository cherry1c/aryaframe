package grpc

import (
	"reflect"
	"time"
)

type Option func(client *Client)

func WithAddress(address string) Option {
	return func(c *Client) {
		c.Address = address
	}
}

func WithRequestTimeout(v time.Duration) Option {
	return func(c *Client) {
		c.RequestTimeout = v
	}
}

func WithService(s AbsService) Option {
	return func(c *Client) {
		c.Service = s
	}
}

func WithName(s string) Option {
	return func(c *Client) {
		c.Name = s
	}
}
func WithServiceFunc(s reflect.Value) Option {
	return func(c *Client) {
		c.RegisterFunc = s
	}
}
