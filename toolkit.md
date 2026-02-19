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
3.**Verify Install**
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
```

5. Run (all OS):
textgo run main.go

Browser → http://localhost:8080
Expected:textHello, Moringa AI Essentials! 🚀
You requested: /
Stop: Ctrl+C in terminal.

6. AI Prompt Journal
Used AI (ai.moringaschool.com/ChatGPT) for fast learning.

Prompt 1: "Step-by-step install Go 1.26 on macOS/Windows/Linux including verify" → Guided downloads/PATH. Very helpful.
Prompt 2: "Go modules basics + project init command" → Explained go mod init. Clear.
Prompt 3: "Simplest net/http Hello World server with comments" → Base code; customized.
Prompt 4: "Fix 'stat main.go: no such file' error on macOS" → Checked dir/file existence. Solved my issue.

AI turned hours of confusion into quick wins.
7. Common Issues & Fixes

"stat main.go: no such file or directory" → Wrong folder/file missing. Fix: ls (macOS/Linux) or dir (Windows); ensure main.go in current dir.
"go: command not found" → PATH not set. Fix: Reopen terminal/Command Prompt; add Go bin to PATH.
Browser "connection refused" → Server not running/wrong URL. Fix: Run go run main.go; use http://localhost:8080 or 127.0.0.1:8080.
Installer PATH issues (Windows) → Restart Command Prompt or add manually.

8. References

1. Official Install Guide: https://go.dev/doc/install
2. Go 1.26 Release Notes: https://go.dev/doc/go1.26
3. A Tour of Go: https://go.dev/tour/welcome/1
4. Go by Example (HTTP): https://gobyexample.com/http-servers
5. GitHub Repo: https://github.com/maritimcalvin/calvin-go-toolkit
6. Tested on: macOS (MacBook Air), Nairobi, February 19, 2026

