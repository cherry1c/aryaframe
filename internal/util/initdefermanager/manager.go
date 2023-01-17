package initdefermanager

import (
	"aryaframe/internal/util/xgo"
	"sync"
)

var (
	initFrameWorkFuncs  []func() error
	initUserFuncs       []func() error
	deferFrameWorkFuncs []func() error
	deferUserFuncs      []func() error
	mu                  sync.Mutex
)

func RegisterFrameWork(fn ...func() error) {
	mu.Lock()
	defer mu.Unlock()
	initFrameWorkFuncs = append(initFrameWorkFuncs, fn...)
}

func Register(fn ...func() error) {
	mu.Lock()
	defer mu.Unlock()
	initUserFuncs = append(initUserFuncs, fn...)
}

func RegisterFrameWorkDefer(fn ...func() error) {
	mu.Lock()
	defer mu.Unlock()
	deferFrameWorkFuncs = append(deferFrameWorkFuncs, fn...)
}

func RegisterDefer(fn ...func() error) {
	mu.Lock()
	defer mu.Unlock()
	deferUserFuncs = append(deferUserFuncs, fn...)
}

// 执行注册的初始化函数，顺序执行直到初始化函数返回error，并将error返回
func Init() error {
	mu.Lock()
	defer mu.Unlock()

	var fns []func() error
	fns = append(fns, initFrameWorkFuncs...)
	fns = append(fns, initUserFuncs...)

	// init and notify all timers that registered in init()/Init()
	//fns = append(fns, timer.Controller.InitAndNotify)

	return xgo.SerialUntilError(fns...)()
}

// 执行注册的defer函数，逆序执行所有defer函数，收集返回error聚合为errors返回
func Defer() error {
	mu.Lock()
	defer mu.Unlock()

	var fns []func() error
	fns = append(fns, deferFrameWorkFuncs...)
	fns = append(fns, deferUserFuncs...)

	// reverse fns in place
	for i, j := 0, len(fns)-1; i < j; i, j = i+1, j-1 {
		fns[i], fns[j] = fns[j], fns[i]
	}

	return xgo.SerialWithError(fns...)()
}
