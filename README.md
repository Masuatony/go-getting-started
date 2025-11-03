# Go Hello CLI
**A minimal, beginner-friendly command-line tool built in Go**  
*Greet the world (or anyone!) with today’s date — in a single, portable binary.*

---

![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go&logoColor=white)
![Platform](https://img.shields.io/badge/platform-Linux%20%7C%20macOS%20%7C%20Windows-blue)
![License](https://img.shields.io/badge/license-MIT-green)

---

## Overview

This project demonstrates the **core fundamentals of Go (Golang)** through a simple yet functional CLI tool:

- Parses a `-name` flag (with a default value)
- Fetches the current date using the standard library
- Prints a clean, formatted greeting
- Compiles to a **standalone binary** (no runtime needed)

Perfect for beginners learning Go, CLI development, or cross-compilation!

---

## Features

- Zero external dependencies (100% standard library)
- Single-file `main.go` — easy to read and extend
- Cross-platform: runs on Linux, macOS, and Windows
- Built with Go modules (`go.mod`)
- Includes error-safe flag parsing and date formatting

---

## Prerequisites

| Requirement       | Version     | Notes |
|-------------------|-------------|-------|
| Go                | `1.21+`     | [Download](https://go.dev/dl/) |
| Terminal / Shell  | Any         | PowerShell, Bash, Zsh, etc. |
| Code Editor (Optional) | VS Code + Go extension | Recommended for IntelliSense |

---

## Quick Start

```bash
# 1. Clone the repo
git clone https://github.com/Masuatony/go-getting-started.git
cd go-getting-started

# 2. Build the binary
go build -o hello

# 3. Run it!
./hello -name Alice