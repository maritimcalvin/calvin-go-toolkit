// main.go - Simple "Hello World" HTTP server in Go
package main

import (
    "fmt"      // For printing to console
    "net/http" // Built-in package for HTTP servers and clients
)

// handler function responds to every request to "/"
func handler(w http.ResponseWriter, r *http.Request) {
    // Send plain text response to the browser/client
    fmt.Fprintf(w, "Hello, Moringa AI Essentials! 🚀\nYou requested: %s", r.URL.Path)
}

func main() {
    // Register the handler for the root path "/"
    http.HandleFunc("/", handler)

    // Print startup message to console
    fmt.Println("Server is running on http://localhost:8080")

    // Start the server on port 8080 (blocks until stopped)
    // nil uses the default router
    http.ListenAndServe(":8080", nil)
}
