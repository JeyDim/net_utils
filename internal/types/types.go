package types

const MaxBodySize = 10 * 1024 * 1024

type RequestInfo struct {
	Timestamp     string              `json:"timestamp"`
	Method        string              `json:"method"`
	URL           string              `json:"url"`
	Path          string              `json:"path"`
	Query         map[string][]string `json:"query"`
	Headers       map[string][]string `json:"headers"`
	Host          string              `json:"host"`
	RemoteAddr    string              `json:"remote_addr"`
	ContentLength int64               `json:"content_length"`
	Body          string              `json:"body,omitempty"`
	BodyTruncated bool                `json:"body_truncated,omitempty"`
	Protocol      string              `json:"protocol"`
}

type HTTPClientResponse struct {
	Timestamp      string              `json:"timestamp"`
	RequestMethod  string              `json:"request_method"`
	RequestURL     string              `json:"request_url"`
	RequestHeaders map[string][]string `json:"request_headers,omitempty"`
	RequestBody    string              `json:"request_body,omitempty"`
	StatusCode     int                 `json:"status_code"`
	Status         string              `json:"status"`
	Headers        map[string][]string `json:"headers"`
	Body           string              `json:"body"`
	BodyTruncated  bool                `json:"body_truncated,omitempty"`
	Duration       string              `json:"duration"`
	Error          string              `json:"error,omitempty"`
}
