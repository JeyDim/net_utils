package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"network-testing-util/internal/domain"
	"time"
)

type EchoHandler struct{}

func NewEchoHandler() *EchoHandler {
	return &EchoHandler{}
}

func (h *EchoHandler) Handle(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, domain.MaxBodySize)

	bodyBytes, err := io.ReadAll(r.Body)
	bodyTruncated := false
	statusCode := http.StatusOK

	if err != nil {
		if _, ok := err.(*http.MaxBytesError); ok {
			log.Printf("⚠️  Request body too large from %s (max: %d bytes)", r.RemoteAddr, domain.MaxBodySize)
			bodyTruncated = true
			bodyBytes = []byte("Body exceeded maximum size limit")
			statusCode = http.StatusRequestEntityTooLarge
		} else {
			log.Printf("Error reading body: %v", err)
			bodyBytes = []byte{}
		}
	}
	defer r.Body.Close()

	reqInfo := response{
		Timestamp:     time.Now().Format(time.RFC3339),
		Method:        r.Method,
		URL:           r.URL.String(),
		Path:          r.URL.Path,
		Query:         r.URL.Query(),
		Headers:       r.Header,
		Host:          r.Host,
		RemoteAddr:    r.RemoteAddr,
		ContentLength: r.ContentLength,
		Body:          string(bodyBytes),
		BodyTruncated: bodyTruncated,
		Protocol:      r.Proto,
	}

	log.Printf("📥 %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(reqInfo); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
