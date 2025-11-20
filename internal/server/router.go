package server

import (
	"net/http"

	"network-testing-util/internal/handlers"
	"network-testing-util/internal/httpclient"
	"network-testing-util/internal/security"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	validator := security.NewValidator()
	httpClient := httpclient.NewClient(validator)

	requestHandler := handlers.NewRequestHandler()
	clientHandler := handlers.NewClientHandler(validator, httpClient)

	mux.HandleFunc("/http/get", clientHandler.Handle)
	mux.HandleFunc("/http/post", clientHandler.Handle)
	mux.HandleFunc("/http/put", clientHandler.Handle)
	mux.HandleFunc("/http/delete", clientHandler.Handle)
	mux.HandleFunc("/http/patch", clientHandler.Handle)
	mux.HandleFunc("/http/head", clientHandler.Handle)
	mux.HandleFunc("/http/options", clientHandler.Handle)

	mux.HandleFunc("/", requestHandler.Handle)

	return mux
}
