package grpc

import (
	"reflect"
)

type AbsMethod interface {
	GetName() string
	SetName(name string)
	GetMethod() reflect.Value
	SetMethod(method reflect.Value)
}

type BasicMethod struct {
	name     string
	Instance reflect.Value
	Service  AbsService
}

func (m *BasicMethod) GetName() string {
	return m.name
}

func (m *BasicMethod) SetName(name string) {
	m.name = name
}
func (m *BasicMethod) GetMethod() reflect.Value {
	return m.Instance
}
func (m *BasicMethod) SetMethod(method reflect.Value) {
	m.Instance = method
}

func NewBasicMethod(name string) AbsMethod {
	return &BasicMethod{
		name:     name,
		Instance: reflect.Value{},
		Service:  nil,
	}
}

type AbsService interface {
	GetName() string
	GetMethod(methodName string) (AbsMethod, bool)
	GetMethodMap() map[string]AbsMethod
	SetName(name string)
	SetMethod(name string, method AbsMethod)
}

type BasicService struct {
	name      string
	MethodMap map[string]AbsMethod
}

func (s *BasicService) GetName() string {
	return s.name
}
func (s *BasicService) GetMethod(methodName string) (AbsMethod, bool) {
	if v, ok := s.MethodMap[methodName]; ok {
		return v, true
	}
	return nil, false
}
func (s *BasicService) GetMethodMap() map[string]AbsMethod {
	return s.MethodMap
}

func (s *BasicService) SetName(name string) {
	s.name = name
}
func (s *BasicService) SetMethod(name string, method AbsMethod) {
	s.MethodMap[name] = method
}

func NewBasicService(name string) AbsService {
	return &BasicService{
		name:      name,
		MethodMap: make(map[string]AbsMethod),
	}
}

type AbsRegistry interface {
	GetName() string
	GetService(serviceName string) (AbsService, bool)
	GetMethod(serviceName string, methodName string) (AbsMethod, bool)
	GetServiceMap() map[string]AbsService
	SetName(name string)
	SetService(serviceName string, service AbsService)
}

type BasicRegistry struct {
	Name       string
	ServiceMap map[string]AbsService
}

func (r *BasicRegistry) GetName() string {
	return r.Name
}
func (r *BasicRegistry) GetService(serviceName string) (AbsService, bool) {
	if v, ok := r.ServiceMap[serviceName]; ok {
		return v, true
	}
	return nil, false
}
func (r *BasicRegistry) GetMethod(serviceName string, methodName string) (AbsMethod, bool) {
	if v, ok := r.ServiceMap[serviceName].GetMethod(methodName); ok {
		return v, true
	}
	return nil, false
}
func (r *BasicRegistry) GetServiceMap() map[string]AbsService {
	return r.ServiceMap
}

func (r *BasicRegistry) SetName(name string) {
	r.Name = name
}
func (r *BasicRegistry) SetService(serviceName string, service AbsService) {
	r.ServiceMap[serviceName] = service
}
