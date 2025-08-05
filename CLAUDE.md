# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go project named "go-get-it" that appears to be in its initial stages. The repository currently contains minimal structure with only a basic README.md file.

## Development Commands

Since this is a new Go project, you'll likely need to initialize it first:

```bash
# Initialize Go module (if not already done)
go mod init github.com/username/go-get-it

# Build the project
go build

# Run the project
go run .

# Run tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Format code
go fmt ./...

# Vet code for potential issues
go vet ./...

# Download dependencies
go mod download

# Tidy up dependencies
go mod tidy
```

## Architecture Notes

The project structure is currently minimal. As the codebase grows, typical Go project patterns should be followed:

- `main.go` for the main application entry point
- Package organization following Go conventions
- Separate packages for different concerns (handlers, models, services, etc.)
- Use of Go modules for dependency management

## Next Steps for Development

When starting development on this project, you'll likely need to:

1. Initialize the Go module if not already done
2. Create the main application structure
3. Set up proper package organization
4. Add necessary dependencies via `go mod`