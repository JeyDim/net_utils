package handler

type response struct {
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
