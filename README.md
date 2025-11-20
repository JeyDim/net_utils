# Network Testing Utility

## Overview

A lightweight Go HTTP server designed for testing network rules, debugging HTTP requests, and inspecting network
traffic. The server provides two main functions:

1. **Request Inspector**: Captures and displays detailed information about incoming HTTP requests
2. **HTTP Client**: Makes outbound HTTP requests to external URLs for testing connectivity and API endpoints

## Purpose

This utility helps developers and network engineers:

- Test and debug HTTP requests
- Make HTTP requests to external APIs and services
- Inspect request/response headers, query parameters, and body content
- Verify network routing and proxy configurations
- Analyze client IP addresses and connection details
- Debug API integrations and webhooks
- Test different HTTP methods (GET, POST, PUT, DELETE, etc.)

## Features

- **Complete Request Capture**: Captures all request details including method, URL, headers, body, and client
  information
- **HTTP Client**: Make outbound requests to external URLs using any HTTP method
- **SSRF Protection**: Validates URLs to prevent requests to private IP addresses
- **Request Size Limiting**: Protects against memory exhaustion with 10MB request/response limit
- **JSON Response**: Returns all captured data as formatted JSON for easy parsing
- **Console Logging**: Logs incoming requests and outbound requests to the console for real-time monitoring
- **Configurable Port**: Supports custom port configuration via PORT environment variable (defaults to 5000)
- **Zero Dependencies**: Built using only Go standard library

## Project Structure

```
.
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── internal/
│   ├── handlers/
│   │   ├── client.go            # HTTP client handler
│   │   └── request.go           # Request inspector handler
│   ├── httpclient/
│   │   └── client.go            # HTTP client with redirect validation
│   ├── security/
│   │   └── validator.go         # URL and IP validation
│   ├── server/
│   │   └── router.go            # Router configuration
│   └── types/
│       └── types.go             # Shared type definitions
├── go.mod                        # Go module definition
└── replit.md                     # Project documentation
```

The project follows Go best practices with:

- **cmd/server**: Main application entry point
- **internal/types**: Shared data structures and constants
- **internal/security**: URL validation and SSRF protection logic
- **internal/httpclient**: HTTP client wrapper with security features
- **internal/handlers**: HTTP request handlers for different endpoints
- **internal/server**: Server configuration and routing

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

## Recent Changes

- **2025-11-20**: Refactored to Go best practices structure
    - Reorganized code into standard Go project layout (cmd/, internal/)
    - Created separate packages for types, security, httpclient, handlers, and server
    - Improved code maintainability and testability through better separation of concerns
    - Moved main entry point to cmd/server/main.go
    - Made validator and client components reusable and independently testable
    - All functionality preserved and tested

- **2025-11-20**: Added HTTP client functionality
    - Added endpoints for making outbound HTTP requests (GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS)
    - Implemented SSRF protection to block requests to private IP addresses
    - Added URL validation and automatic https:// prefix
    - Added request/response body size limiting (10MB)
    - Added 30-second timeout for outbound requests
    - Enhanced logging to show outbound request details and duration
    - Fixed redirect validation to prevent SSRF bypasses

- **2025-11-20**: Initial project creation
    - Created Go HTTP server with request inspection
    - Added JSON formatting for responses
    - Configured workflow to run on port 5000
    - Added console logging for incoming requests
    - Implemented request body size limiting with http.MaxBytesReader
    - Added security measures against DoS attacks

## Technical Details

- Language: Go 1.24
- Server: Go net/http standard library
- Binding: 0.0.0.0:5000 (accessible from all network interfaces)
- Response Format: JSON with pretty printing
