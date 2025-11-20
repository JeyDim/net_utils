# net_utils

![GitHub last commit](https://img.shields.io/github/last-commit/JeyDim/net_utils?style=flat-square)
![GitHub Stars](https://img.shields.io/github/stars/JeyDim/net_utils?style=flat-square)
![Docker Pulls](https://img.shields.io/badge/docker-ready-blue?logo=docker&style=flat-square)

A lightweight Go HTTP server designed for testing network rules, debugging HTTP requests, and inspecting network
traffic. The server provides two main functions:

1. **Request Inspector**: Captures and displays detailed information about incoming HTTP requests
2. **HTTP Client**: Makes outbound HTTP requests to external URLs for testing connectivity and API endpoints

---

## Purpose

This utility helps developers and network engineers:

- Test and debug HTTP requests
- Make HTTP requests to external APIs and services
- Inspect request/response headers, query parameters, and body content
- Verify network routing and proxy configurations
- Analyze client IP addresses and connection details
- Debug API integrations and webhooks
- Test different HTTP methods (GET, POST, PUT, DELETE, etc.)

## 🚀 Features

- **Complete Request Capture**: Captures all request details including method, URL, headers, body, and client
  information
- **HTTP Client**: Make outbound requests to external URLs using any HTTP method
- **SSRF Protection**: Validates URLs to prevent requests to private IP addresses
- **Request Size Limiting**: Protects against memory exhaustion with 10MB request/response limit
- **JSON Response**: Returns all captured data as formatted JSON for easy parsing
- **Console Logging**: Logs incoming requests and outbound requests to the console for real-time monitoring
- **Configurable Port**: Supports custom port configuration via PORT environment variable (defaults to 5000)
- **Zero Dependencies**: Built using only Go standard library

---

## 🐳 Docker Support

The project includes a Docker image for all major platforms!  
Automatically built and published to [GitHub Container Registry](https://github.com/JeyDim/net_utils/pkgs/container/net_utils).

### Build & Run with Docker

```sh
docker pull ghcr.io/jeydim/net_utils:latest
docker run --rm ghcr.io/jeydim/net_utils [command]
```

---

## 🛠️ Installation

### Clone

```sh
git clone https://github.com/JeyDim/net_utils.git
cd net_utils
```

### Build Locally

```sh
# Replace with your build instructions if needed
```

---

## 📘 Documentation

- [Usage Examples](USAGE.md)

---

## 🤝 Contributing

Pull requests are welcome!  
For major changes, please open an issue first to discuss what you’d like to change.

---

## 📄 License

This project is licensed under the MIT License.

---

## ✉️ Contact

Created and maintained by [JeyDim](https://github.com/JeyDim).  
For questions or suggestions, feel free to [open an issue](https://github.com/JeyDim/net_utils/issues).
