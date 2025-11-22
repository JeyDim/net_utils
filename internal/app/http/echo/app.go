package echo

import (
	"context"
	"fmt"
	"net/http"
	"network-testing-util/internal/app/http/echo/handler"
	"os"
	"strings"
	"time"
)

type App struct {
	httpServer *http.Server
}

func NewApp() *App {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	return &App{
		httpServer: &http.Server{
			Addr:              fmt.Sprintf(":%v", port),
			ReadHeaderTimeout: 5 * time.Second,
			WriteTimeout:      5 * time.Second,
			IdleTimeout:       5 * time.Second,
			MaxHeaderBytes:    1 << 20,
		},
	}
}

func (a *App) PrintCommonInfo() {
	addr := a.httpServer.Addr
	if strings.HasPrefix(addr, ":") {
		addr = "0.0.0.0" + addr
	}

	fmt.Println("-------------------------------------")
	fmt.Println("=====================================")
	fmt.Println("📣 ECHO Server")
	fmt.Println("=====================================")
	fmt.Printf("📡 Listening on: %s\n", addr)
	fmt.Println("  Shows request details")
	fmt.Println("=====================================")
}

func (a *App) Start(_ context.Context, errors chan<- error) {
	mux := newServeMux()
	a.httpServer.Handler = mux

	go func() {
		errors <- a.httpServer.ListenAndServe()
	}()
}

func (a *App) Stop(ctx context.Context) error {
	return a.httpServer.Shutdown(ctx)
}

func newServeMux() *http.ServeMux {
	mux := http.NewServeMux()

	echoHandler := handler.NewEchoHandler()

	mux.HandleFunc("/", echoHandler.Handle)

	return mux
}
