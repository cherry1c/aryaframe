package main

import (
	"aryaframe/internal"
	"aryaframe/internal/conf"
	"flag"
	"os"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "conf/conf.yaml", "config file path")
	flag.Parse()
}

func main() {
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
