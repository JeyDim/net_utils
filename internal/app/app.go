package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

type Executable interface {
	Start(context.Context, chan<- error)
	Stop(context.Context) error
}
type App struct {
	executables []Executable
}

func NewApp(executables []Executable) *App {
	return &App{
		executables: executables,
	}
}

func (app *App) Run(ctx context.Context) {
	errorChan := make(chan error, 1)

	for _, executable := range app.executables {
		go executable.Start(ctx, errorChan)
	}

	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, os.Interrupt, syscall.SIGTERM)

	select {
	case <-stopSignal:
	case <-errorChan:
	}

	app.stop(ctx)
}

func (app *App) stop(ctx context.Context) {
	for _, executable := range app.executables {
		executable.Stop(ctx)
	}
}
