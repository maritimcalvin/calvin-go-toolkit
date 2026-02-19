# Prompt-Powered Kickstart: Building a Beginner’s Toolkit for Go (Golang)

## 1. Title & Objective

**Technology chosen:** Go (Golang)  
**Why I chose it:** Go is a modern, simple, and fast programming language created by Google. It has clean syntax, excellent performance, and is widely used for building scalable backend services, APIs, cloud tools (like Docker and Kubernetes), and microservices. It's beginner-friendly with few hidden complexities, compiles to a single binary, and is in high demand in 2025–2026 job markets, including Kenya's growing tech scene (fintech, cloud).  
**End goal:** Set up Go locally and run a minimal working HTTP web server that responds with "Hello, Moringa AI Essentials!" when visited in a browser — a foundational step toward building real APIs.

## 2. Quick Summary of the Technology

Go (often called Golang) is an open-source, statically typed, compiled programming language designed for simplicity, reliability, and efficiency.  
- **What is it?** A language focused on fast compilation, garbage collection, built-in concurrency (goroutines), and a powerful standard library — especially for networking and web servers.  
- **Where is it used?** Backend APIs, cloud infrastructure, DevOps tools, command-line utilities, and high-performance systems.  
- **One real-world example:** Docker (container platform), Kubernetes (orchestration), and many microservices at companies like Uber, Dropbox, and local Kenyan startups building scalable fintech solutions.

## 3. System Requirements

- **OS:** Windows 10/11, macOS (any recent version), or Linux (Ubuntu, etc.) — works on all major platforms.  
- **Tools/Editors required:** Any text editor (VS Code recommended — free, with excellent Go extension), terminal/command prompt.  
- **Any packages:** None required upfront — Go includes everything in its standard library for this project. Download ~150 MB installer/archive.

## 4. Installation & Setup Instructions

As of February 2026, the latest stable version is **Go 1.26** (released February 2026).

1. **Download Go:**  
   Go to the official site: https://go.dev/dl/  
   Choose the installer for your OS (e.g., `.msi` for Windows, `.pkg` for macOS, `.tar.gz` for Linux).

2. **Install:**
   - **Windows:** Run the `.msi` file → follow prompts (default path: `C:\Program Files\Go`). It auto-adds to PATH.
   - **macOS:** Open the `.pkg` file → follow installer prompts (installs to `/usr/local/go`).
   - **Linux:**  
     ```bash
     # Remove old version if exists
     sudo rm -rf /usr/local/go
     # Extract (replace with your downloaded file name, e.g., go1.26.linux-amd64.tar.gz)
     sudo tar -C /usr/local -xzf go1.26.linux-amd64.tar.gz
Open a **new** Terminal window and verify:
go version
textExpected: `go version go1.26.0 darwin/arm64` 
4. (If needed) Add to PATH: Edit `~/.zshrc` with `export PATH=$PATH:/usr/local/go/bin`, then `source ~/.zshrc`.

**Windows:**
1. Download `go1.26.0.windows-amd64.msi`.
2. Double-click the `.msi` file → follow installer prompts (default path: `C:\Program Files\Go` — it adds to PATH automatically).
3. Open a **new** Command Prompt or PowerShell and verify:
go version
textExpected: `go version go1.26.0 windows/amd64`.
4. (If PATH issue) Restart Command Prompt or add manually via System Properties → Environment Variables → Path → add `C:\Program Files\Go\bin`.

**Linux (e.g., Ubuntu/Debian):**
1. Download `go1.26.0.linux-amd64.tar.gz` (or arm64 if ARM-based).
2. Extract and install:
sudo rm -rf /usr/local/go   # Remove old version if exists
sudo tar -C /usr/local -xzf go1.26.0.linux-amd64.tar.gz
text3. Add to PATH (edit `~/.profile` or `~/.bashrc`):
export PATH=$PATH:/usr/local/go/bin
textApply: `source ~/.profile`
4. Verify in new terminal:
go version
textExpected: `go version go1.26.0 linux/amd64`.

**Project setup (all OS):**
mkdir calvin-go-toolkit
cd calvin-go-toolkit
go mod init calvin-go-toolkit

5. Minimal Working Example

**What the example does:** Basic HTTP server using `net/http`. Visit http://localhost:8080 to see the greeting.

**Code (`main.go`):**

```go
// main.go - Simple "Hello World" HTTP server in Go
package main

import (
    "fmt"      // For console output
    "net/http" // Built-in HTTP package
)

// handler responds to requests at "/"
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, Moringa AI Essentials! 🚀\nYou requested: %s", r.URL.Path)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}


