package models

type HTTPClientResponse struct {
	Timestamp      string
	RequestMethod  string
	RequestURL     string
	RequestHeaders map[string][]string
	RequestBody    string
	StatusCode     int
	Status         string
	Headers        map[string][]string
	Body           string
	BodyTruncated  bool
	Duration       string
	Error          string
}
