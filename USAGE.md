## How to Use

### Request Inspector

Send any HTTP request to the server and it will echo back all request details:

```bash
curl -X POST http://localhost:5000/test?foo=bar -H "Content-Type: application/json" -H "X-Custom-Header: test-value" -d '{"key": "value"}'
```

The server returns JSON with:

- `timestamp`: Request time in RFC3339 format
- `method`: HTTP method (GET, POST, etc.)
- `url`: Full URL including query string
- `path`: Request path
- `query`: Query parameters as key-value map
- `headers`: All HTTP headers
- `host`: Host header value
- `remote_addr`: Client IP address and port
- `content_length`: Request body size
- `body`: Request body content
- `body_truncated`: Whether body was truncated due to size limit
- `protocol`: HTTP protocol version

### HTTP Client Endpoints

Make outbound HTTP requests to external URLs:

**Available Methods:**

- `/http/get?url=<target>` - Make GET request
- `/http/post?url=<target>` - Make POST request
- `/http/put?url=<target>` - Make PUT request
- `/http/delete?url=<target>` - Make DELETE request
- `/http/patch?url=<target>` - Make PATCH request
- `/http/head?url=<target>` - Make HEAD request
- `/http/options?url=<target>` - Make OPTIONS request

**Examples:**

```bash
# Simple GET request
curl "http://localhost:5000/http/get?url=https://api.github.com"

# GET request (URL without protocol defaults to https)
curl "http://localhost:5000/http/get?url=httpbin.org/get"

# POST request with body
curl -X POST "http://localhost:5000/http/post?url=https://httpbin.org/post" \
  -H "Content-Type: application/json" \
  -d '{"test": "data"}'
```

**Response Format:**

- `timestamp`: Request time
- `request_method`: HTTP method used
- `request_url`: Target URL
- `request_headers`: Headers sent with request
- `request_body`: Body sent with request
- `status_code`: HTTP response status code
- `status`: HTTP response status text
- `headers`: Response headers from target
- `body`: Response body from target
- `body_truncated`: Whether response body was truncated
- `duration`: Time taken for the request
- `error`: Error message if request failed
  
**Security:**

- **SSRF Protection**: Requests to private IP addresses (localhost, 127.0.0.1, 192.168.x.x, 10.x.x.x, etc.) are blocked
  to prevent SSRF attacks
- **Redirect Validation**: All HTTP redirects are validated - redirects to private IPs are blocked even if the initial
  URL is public
- **Request Size Limiting**: Inbound request bodies exceeding 10MB return HTTP 413 (Request Entity Too Large) with error
  details
- **Response Size Limiting**: Outbound response bodies are limited to 10MB and truncation is indicated in the response
- **Timeout Protection**: 30-second timeout on all outbound requests to prevent hanging connections
- **Redirect Limit**: Maximum of 10 redirects to prevent redirect loops