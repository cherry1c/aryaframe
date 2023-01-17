package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"reflect"
)

var (
	typeOf    reflect.Type
	valueOf   reflect.Value
	paramList []reflect.Value
)

func RegisterService(tOf reflect.Type, vOf reflect.Value, pList []reflect.Value) {
	typeOf = tOf
	valueOf = vOf
	paramList = pList
}

func registerServer(s *grpc.Server) {
	valueOf.Call(paramList)
	fmt.Printf("%s register success.\n", typeOf.Name())
}
