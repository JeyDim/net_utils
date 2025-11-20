package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"network-testing-util/internal/httpclient"
	"network-testing-util/internal/security"
	"network-testing-util/internal/types"
)

type ClientHandler struct {
	validator  *security.Validator
	httpClient *httpclient.Client
}

func NewClientHandler(validator *security.Validator, httpClient *httpclient.Client) *ClientHandler {
	return &ClientHandler{
		validator:  validator,
		httpClient: httpClient,
	}
}

func (h *ClientHandler) Handle(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	targetURL := r.URL.Query().Get("url")
	if targetURL == "" {
		http.Error(w, "Missing 'url' query parameter", http.StatusBadRequest)
		return
	}

	if !strings.HasPrefix(targetURL, "http://") && !strings.HasPrefix(targetURL, "https://") {
		targetURL = "https://" + targetURL
	}

	if err := h.validator.ValidateURL(targetURL); err != nil {
		response := types.HTTPClientResponse{
			Timestamp:     time.Now().Format(time.RFC3339),
			RequestMethod: r.URL.Path[len("/http/"):],
			RequestURL:    targetURL,
			Error:         err.Error(),
			Duration:      time.Since(startTime).String(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	method := strings.ToUpper(r.URL.Path[len("/http/"):])

	r.Body = http.MaxBytesReader(w, r.Body, types.MaxBodySize)
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		if _, ok := err.(*http.MaxBytesError); ok {
			response := types.HTTPClientResponse{
				Timestamp:     time.Now().Format(time.RFC3339),
				RequestMethod: method,
				RequestURL:    targetURL,
				Error:         fmt.Sprintf("Request body too large (max: %d bytes)", types.MaxBodySize),
				Duration:      time.Since(startTime).String(),
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusRequestEntityTooLarge)
			json.NewEncoder(w).Encode(response)
			r.Body.Close()
			return
		}
	}
	defer r.Body.Close()

	response, err := h.httpClient.Do(method, targetURL, bodyBytes, r.Header)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
