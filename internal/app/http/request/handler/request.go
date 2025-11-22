package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"network-testing-util/internal/clients/http/common"
	"network-testing-util/internal/domain"
	"network-testing-util/internal/domain/models"
	"strings"
	"time"

	"network-testing-util/internal/security"
)

type RequestHandler struct {
	validator  *security.Validator
	httpClient *common.Client
}

func NewRequestHandler(validator *security.Validator, httpClient *common.Client) *RequestHandler {
	return &RequestHandler{
		validator:  validator,
		httpClient: httpClient,
	}
}

func (h *RequestHandler) Handle(w http.ResponseWriter, r *http.Request) {
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
		response := models.HTTPClientResponse{
			Timestamp:     time.Now().Format(time.RFC3339),
			RequestMethod: r.URL.Path[len("/http/"):],
			RequestURL:    targetURL,
			Error:         err.Error(),
			Duration:      time.Since(startTime).String(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(response)
		return
	}

	method := strings.ToUpper(r.URL.Path[len("/http/"):])

	r.Body = http.MaxBytesReader(w, r.Body, domain.MaxBodySize)
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		if _, ok := err.(*http.MaxBytesError); ok {
			response := models.HTTPClientResponse{
				Timestamp:     time.Now().Format(time.RFC3339),
				RequestMethod: method,
				RequestURL:    targetURL,
				Error:         fmt.Sprintf("Request body too large (max: %d bytes)", domain.MaxBodySize),
				Duration:      time.Since(startTime).String(),
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusRequestEntityTooLarge)
			_ = json.NewEncoder(w).Encode(response)
			_ = r.Body.Close()
			return
		}
	}
	defer r.Body.Close()

	httpClientResponse, err := h.httpClient.Do(method, targetURL, bodyBytes, r.Header)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(httpClientResponse); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
