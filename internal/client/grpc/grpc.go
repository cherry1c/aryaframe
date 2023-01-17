package grpc

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"reflect"
	"time"
)

type Client struct {
	Name           string
	Address        string
	RequestTimeout time.Duration
	Service        AbsService
	RegisterFunc   reflect.Value
}

var client map[string]*Client

func init() {
	client = make(map[string]*Client)
}

func newClient(opts ...Option) *Client {
	c := &Client{
		RequestTimeout: 5 * time.Second,
	}
	for _, op := range opts {
		op(c)
	}
	return c
}

func SetClientOptions(serviceName string, opts ...Option) {
	if v, ok := client[serviceName]; !ok {
		return
	} else {
		for _, op := range opts {
			op(v)
		}
	}
}

func Start() error {
	for _, c := range client {
		con, err := grpc.Dial(c.Address, grpc.WithInsecure())
		if err != nil {
			fmt.Printf("dial %s failed err: %s\n", c.Address, err.Error())
			return err
		}
		service := c.RegisterFunc.Call([]reflect.Value{reflect.ValueOf(con)})[0]
		c.Service = NewBasicService(c.Name)
		for i := 0; i < service.NumMethod(); i++ {
			methodName := service.Method(i).Type().Name()
			method := NewBasicMethod(methodName)
			method.SetMethod(service.Method(i))
			c.Service.SetMethod(methodName, method)
		}
	}
	return nil
}

func RegisterService(serviceName string, f reflect.Value) {
	c := newClient(WithName(serviceName), WithServiceFunc(f))
	client[serviceName] = c
}

func Call(s string, m string, ctx context.Context, req interface{}) (interface{}, error) {
	method, ok := client[s+"Client"].Service.GetMethod(m)
	if !ok {
		return nil, errors.New("no method")
	}
	result := method.GetMethod().Call([]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(req)})
	return result[0], result[1].Interface().(error)
}
