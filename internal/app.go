package internal

import (
	"aryaframe/internal/client"
	"aryaframe/internal/log"
	"aryaframe/internal/server"
	"aryaframe/internal/util/initdefermanager"
	"aryaframe/internal/util/xcycle"
	"context"
	"fmt"
	"go.uber.org/multierr"
	"golang.org/x/sync/errgroup"
	"sync"
)

type App struct {
	cycle    *xcycle.Cycle
	initOnce sync.Once
	servers  []server.Server
}

const module = "aryaframe"

func (app *App) Init() (err error) {
	app.initOnce.Do(func() {
		app.servers = make([]server.Server, 2)
		initdefermanager.RegisterFrameWork(log.Init)
		initdefermanager.RegisterFrameWork(server.Init)
		initdefermanager.RegisterFrameWork(client.Init)

		if err = initdefermanager.Init(); err != nil {
			fmt.Printf("init app failed err: %s\n", err.Error())
		}
	})
	return
}

func (app *App) Run() (err error) {
	defer func() {
		if deferErr := initdefermanager.Defer(); deferErr != nil {
			fmt.Printf("call pkg.defer failed err: %s\n", deferErr.Error())
			err = multierr.Append(err, deferErr)
		}
	}()
	app.servers = append(app.servers, server.Servers...)

	app.cycle.Run(app.startServers)

	if err = <-app.cycle.Wait(); err != nil {
		fmt.Printf("run app failed err: %s\n", err.Error())
	} else {
		fmt.Printf("run app successful\n")
	}
	return
}

func (app *App) startServers() error {
	eg, ctx := errgroup.WithContext(context.Background())
	for _, s := range app.servers {
		s := s
		eg.Go(func() (err error) {
			err = s.Serve(ctx)
			return
		})
	}
	return eg.Wait()
}
