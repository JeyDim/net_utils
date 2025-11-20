package httpclient

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"network-testing-util/internal/security"
	"network-testing-util/internal/types"
)

type Client struct {
	validator  *security.Validator
	httpClient *http.Client
}

func NewClient(validator *security.Validator) *Client {
	return &Client{
		validator: validator,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) Do(method, targetURL string, bodyBytes []byte, headers http.Header) (*types.HTTPClientResponse, error) {
	startTime := time.Now()

	var req *http.Request
	var err error

	if len(bodyBytes) > 0 {
		req, err = http.NewRequest(method, targetURL, bytes.NewReader(bodyBytes))
	} else {
		req, err = http.NewRequest(method, targetURL, nil)
	}

	if err != nil {
		return &types.HTTPClientResponse{
			Timestamp:     time.Now().Format(time.RFC3339),
			RequestMethod: method,
			RequestURL:    targetURL,
			Error:         fmt.Sprintf("Failed to create request: %v", err),
			Duration:      time.Since(startTime).String(),
		}, err
	}

	requestHeaders := make(map[string][]string)
	hopByHopHeaders := map[string]bool{
		"Connection":          true,
		"Keep-Alive":          true,
		"Proxy-Authenticate":  true,
		"Proxy-Authorization": true,
		"Te":                  true,
		"Trailer":             true,
		"Transfer-Encoding":   true,
		"Upgrade":             true,
	}

	for key, values := range headers {
		if !hopByHopHeaders[key] && len(values) > 0 {
			req.Header.Del(key)
			for _, value := range values {
				req.Header.Add(key, value)
			}
			requestHeaders[key] = values
		}
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 10 {
				return fmt.Errorf("stopped after 10 redirects")
			}
			if err := c.validator.ValidateURL(req.URL.String()); err != nil {
				return fmt.Errorf("redirect blocked: %v", err)
			}
			return nil
		},
	}

	resp, err := client.Do(req)
	duration := time.Since(startTime)

	response := &types.HTTPClientResponse{
		Timestamp:      time.Now().Format(time.RFC3339),
		RequestMethod:  method,
		RequestURL:     targetURL,
		RequestHeaders: requestHeaders,
		Duration:       duration.String(),
	}

	if len(bodyBytes) > 0 {
		response.RequestBody = string(bodyBytes)
	}

	if err != nil {
		response.Error = fmt.Sprintf("Request failed: %v", err)
		log.Printf("🌐 %s %s -> Error: %v", method, targetURL, err)
		return response, err
	}
	defer resp.Body.Close()

	response.StatusCode = resp.StatusCode
	response.Status = resp.Status
	response.Headers = resp.Header

	if resp.Header.Get("Content-Encoding") == "gzip" {
		gr, _ := gzip.NewReader(resp.Body)
		defer gr.Close()
		resp.Body = gr
	}
	respBodyBytes, err := io.ReadAll(io.LimitReader(resp.Body, types.MaxBodySize))
	if err != nil {
		response.Body = ""
	} else {
		response.Body = string(respBodyBytes)
		if int64(len(respBodyBytes)) >= types.MaxBodySize {
			response.BodyTruncated = true
			response.Body = response.Body + "... (truncated)"
		}
	}

	log.Printf("🌐 %s %s -> %d %s (%s)", method, targetURL, resp.StatusCode, resp.Status, duration)

	return response, nil
}
