package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"network-testing-util/internal/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	router := server.NewRouter()
	addr := "0.0.0.0:" + port

	fmt.Println("=====================================")
	fmt.Println("🚀 Network Testing Utility Server")
	fmt.Println("=====================================")
	fmt.Printf("📡 Listening on: %s\n", addr)
	fmt.Println("📋 This server captures and displays all incoming request details")
	fmt.Println()
	fmt.Println("HTTP Client Endpoints:")
	fmt.Println("  /http/get?url=<target>     - Make GET request")
	fmt.Println("  /http/post?url=<target>    - Make POST request")
	fmt.Println("  /http/put?url=<target>     - Make PUT request")
	fmt.Println("  /http/delete?url=<target>  - Make DELETE request")
	fmt.Println("  /http/patch?url=<target>   - Make PATCH request")
	fmt.Println("  /http/head?url=<target>    - Make HEAD request")
	fmt.Println("  /http/options?url=<target> - Make OPTIONS request")
	fmt.Println()
	fmt.Println("Request Inspector:")
	fmt.Println("  Any other path             - Shows request details")
	fmt.Println("=====================================")
	fmt.Println()

	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
