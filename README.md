# KAMI CLI Todo Manager

![CI](https://img.shields.io/github/actions/workflow/status/Debashich/goCLI/go.yaml?label=CI&style=for-the-badge&logo=github)
![Go](https://img.shields.io/badge/Go-1.22.6-blue?style=for-the-badge&logo=go)

A lightweight and interactive CLI Todo Manager built with Go, featuring JSON persistence, REPL support, and automated CI/CD pipelines using GitHub Actions.

---

## Preview

<img width="1920" height="1019" alt="Screenshot from 2026-05-24 19-23-41" src="https://github.com/user-attachments/assets/33b544ae-035b-473e-96ff-6af12c07d86f" />

---
## Features

- Interactive REPL mode with slash commands
- Full CRUD task management
- Persistent JSON-based storage
- Clean terminal table rendering
- Lightweight and minimal
- CI/CD integration with GitHub Actions
---
## Tech Stack

- **Go (Golang)** - Core application development
- **Go Modules** - Dependency management
- **`aquasecurity/table`** - Terminal table rendering
- **JSON (`todos.json`)** - Persistent local storage
- **GitHub Actions** - CI/CD automation
- **`golangci-lint`** - Linting and static analysis
---

## Installation

### 1. Clone the Repository

```bash
git clone https://github.com/Debashich/Kami.git
cd Kami
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Build the Binary

```bash
go build -o todo
```

---

## Quick Start

### 1. Run the Application

```bash
./todo
```

Or directly with Go:

```bash
go run .
```

---

# Interactive REPL Mode

### Available Commands

| Command | Description |
|---|---|
| `/add Finish project` | Add a new task |
| `/list` | Display all tasks |
| `/toggle 0` | Toggle task completion |
| `/edit 0 Updated text` | Edit a task |
| `/del 0` | Delete a task |
| `/help` | Show help menu |
| `/exit` | Exit the REPL |

---

# Flag Mode

Run commands directly using CLI flags.

### Add a Task

```bash
./todo -add "Finish Go project"
```

### List Tasks

```bash
./todo -list
```

### Toggle Completion

```bash
./todo -toggle 0
```

### Edit a Task

```bash
./todo -edit "0:Complete the Go CLI tutorial"
```

### Delete a Task

```bash
./todo -del 0
```

### Help

```bash
./todo -h
```

---

## CI/CD Pipeline

GitHub Actions automatically handles:

- Builds
- Linting
- Testing
- Continuous Integration checks

---

## Project Structure

```
.
├── .github/
│   └── workflows/
│       └── go.yaml
├── commands.go
├── main.go
├── storage.go
├── todo.go
├── todos.json
├── go.mod
├── go.sum
├── ui.go
└── README.md
```

---

## License

Licensed under the MIT License.
