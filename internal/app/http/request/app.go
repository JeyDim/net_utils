package request

import (
	"context"
	"fmt"
	"net/http"
	"network-testing-util/internal/app/http/request/handler"
	"network-testing-util/internal/clients/http/common"
	"network-testing-util/internal/security"
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
		port = "6000"
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
	fmt.Println("📣 Request Server")
	fmt.Println("=====================================")
	fmt.Printf("📡 Listening on: %s\n", addr)
	fmt.Println("📋 This server captures and displays all incoming request details")
	fmt.Println()
	fmt.Println("HTTP Endpoints:")
	fmt.Println("  /http/get?url=<target>     - Make GET request")
	fmt.Println("  /http/post?url=<target>    - Make POST request")
	fmt.Println("  /http/put?url=<target>     - Make PUT request")
	fmt.Println("  /http/delete?url=<target>  - Make DELETE request")
	fmt.Println("  /http/patch?url=<target>   - Make PATCH request")
	fmt.Println("  /http/head?url=<target>    - Make HEAD request")
	fmt.Println("  /http/options?url=<target> - Make OPTIONS request")
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

	validator := security.NewValidator()
	httpClient := common.NewClient(validator)

	requestHandler := handler.NewRequestHandler(validator, httpClient)

	mux.HandleFunc("/http/get", requestHandler.Handle)
	mux.HandleFunc("/http/post", requestHandler.Handle)
	mux.HandleFunc("/http/put", requestHandler.Handle)
	mux.HandleFunc("/http/delete", requestHandler.Handle)
	mux.HandleFunc("/http/patch", requestHandler.Handle)
	mux.HandleFunc("/http/head", requestHandler.Handle)
	mux.HandleFunc("/http/options", requestHandler.Handle)

	return mux
}
