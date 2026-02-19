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

## 5. Minimal Working Example

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

## 6. AI Prompt Journal
Used AI (ai.moringaschool.com/ChatGPT) for fast learning.

Prompt 1: "Step-by-step install Go 1.26 on macOS/Windows/Linux including verify" → Guided downloads/PATH. Very helpful.

Prompt 2: "Go modules basics + project init command" → Explained go mod init. Clear.

Prompt 3: "Simplest net/http Hello World server with comments" → Base code; customized.

Prompt 4: "Fix 'stat main.go: no such file' error on macOS" → Checked dir/file existence. Solved my issue.

AI turned hours of confusion into quick wins.

## 7. Common Issues & Fixes

Here are the most frequent problems beginners face when following this guide, along with step-by-step fixes. These come from my own debugging + AI prompt experiments + common reports online.

- **"go: command not found" or "'go' is not recognized as an internal or external command"**  
  → Go binary not in PATH after installation.  
  **macOS/Linux fix:**  
  - Close and reopen Terminal.  
  - Or add manually: `export PATH=$PATH:/usr/local/go/bin` → add to `~/.zshrc` (macOS) or `~/.bashrc` → `source ~/.zshrc` / `source ~/.bashrc`.  
  **Windows fix:**  
  - Restart Command Prompt/PowerShell.  
  - Or manually add `C:\Program Files\Go\bin` to System Environment Variables → Path → OK → restart terminal.  
  Test: `go version`

- **"stat main.go: no such file or directory"**  
  → You're running `go run main.go` but the file isn't in the current directory (very common!).  
  **Fix:**  
  - Run `ls` (macOS/Linux) or `dir` (Windows) to list files — make sure `main.go` is listed.  
  - If not: Use `cd` to navigate to the correct folder (`cd calvin-go-toolkit`).  
  - Or recreate the file: `nano main.go` / open in VS Code → paste code → save.  
  - Tip: Always run `pwd` (macOS/Linux) or `cd` (Windows) to confirm your location.

- **Browser shows "This site can’t be reached" / "ERR_CONNECTION_REFUSED" / "connection refused"**  
  → Server isn't running, wrong port, or firewall.  
  **Fix:**  
  - Make sure `go run main.go` is still running in the terminal (don't close the window).  
  - Use correct URL: http://localhost:8080 or http://127.0.0.1:8080 (not https).  
  - Try a different browser.  
  - macOS/Linux: Check if port 8080 is blocked (`lsof -i :8080`).  
  - Windows: Temporarily disable Windows Defender Firewall or add exception for Go.

- **"listen tcp :8080: bind: address already in use"**  
  → Another program (or previous instance) is using port 8080.  
  **Fix:**  
  - Stop the old server (Ctrl+C in its terminal).  
  - Or change port in code: `:8081` instead of `:8080` → save → rerun.  
  - macOS/Linux: Find and kill process: `lsof -i :8080` → `kill -9 <PID>`.  
  - Windows: `netstat -ano | findstr :8080` → `taskkill /PID <number> /F`.

- **"module declares its path as: calvin-go-toolkit but was found in .../different-name"**  
  → Folder name doesn't match `go mod init` name.  
  **Fix:**  
  - Rename folder to match: `mv old-name calvin-go-toolkit`  
  - Or re-init: `go mod init calvin-go-toolkit` in the correct folder.

- **"go mod init: modules disabled by GO111MODULE=off" (rare in 2026)**  
  → Old environment variable.  
  **Fix:** Remove or set `export GO111MODULE=on` / `set GO111MODULE=on` (Windows) → restart terminal.

- **VS Code shows red squiggles / "cannot find package" warnings**  
  → Go extension not installed or needs restart.  
  **Fix:**  
  - Install "Go" extension by Go Team at Google in VS Code.  
  - Run `go install golang.org/x/tools/gopls@latest` (needs internet).  
  - Restart VS Code or run "Go: Install/Update Tools" command (Cmd+Shift+P → type "Go Install").

- **Permission denied when extracting tar.gz on Linux**  
  → Missing sudo.  
  **Fix:** Use `sudo tar -C /usr/local -xzf go1.26.0.linux-amd64.tar.gz`

These cover ~90% of first-time issues. If you hit something else, prompt AI: "Fix this Go error on [your OS]: [error message]".

## 8. References

1. Official Install Guide: https://go.dev/doc/install
2. Go 1.26 Release Notes: https://go.dev/doc/go1.26
3. A Tour of Go: https://go.dev/tour/welcome/1
4. Go by Example (HTTP): https://gobyexample.com/http-servers
5. GitHub Repo: https://github.com/maritimcalvin/calvin-go-toolkit
6. Tested on: macOS (MacBook Air), Nairobi, February 19, 2026

