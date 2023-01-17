package main

import (
	"aryaframe/internal"
	"aryaframe/internal/conf"
	"aryaframe/internal/log"
	_ "aryaframe/server"
	"context"
	"flag"
	"fmt"
	"os"
	"runtime/debug"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "conf/conf.yaml", "config file path")
	flag.Parse()
}

func endProcessingImpl(ctx context.Context) {
	if p := recover(); p != nil {
		log.Error("panic", log.String("", fmt.Sprintf("err: %v", p)),
			log.String("panicMsg", string(debug.Stack())))
		debug.PrintStack()
	}
}

func main() {
	defer endProcessingImpl(context.Background())
	if err := conf.NewDefaultConfig().LoadFile(configPath); err != nil {
		os.Exit(3)
	}

	var app internal.App
	err := app.Init()
	if err != nil {
		os.Exit(4)
	}

	err = app.Run()
	if err != nil {
		os.Exit(5)
	}
}
