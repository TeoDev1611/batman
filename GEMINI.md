# GEMINI.md - Batman Logger

## Project Overview

`batman` is a fast and powerful logger for Go, designed for both terminal output and file logging. It features colored terminal output and stores logs in a persistent location (user cache directory).

- **Main Technologies:** Go (1.17+), `github.com/fatih/color` for terminal styling.
- **Architecture:** 
  - `log/`: Core logging logic, configuration, and file generation.
  - `directories/`: Utility to manage log storage locations using `os.UserCacheDir`.
  - `errors/`: Simple error handling utilities.

## Building and Running

Since this is a Go library, standard Go commands apply:

- **Build:** `go build ./...`
- **Test:** `go test ./...`
- **Get Library:** `go get -u github.com/TeoDev1611/batman/log`
- **Lint/Format:** The project uses a custom formatting task in the `Makefile`.
  - Run `make fmt` to format using `gofumpt`, `goimports`, and `dprint`.

## Development Conventions

- **Formatting:** Rigorously follow the `make fmt` standards. Go files use tabs for indentation as per standard Go practices, despite some `.editorconfig` settings.
- **Commits:** Follow [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) (e.g., `feat:`, `fix:`, `docs:`).
- **Logging Levels:** The library supports `Debug`, `Info`, `Warning`, `Error`, and `Fatal`.
- **Log Format:** Files are generated with a space-separated format: `YYYY-MM-DD HH:MM:SS LEVEL MESSAGE`.
- **Error Handling:** Use the `errors.CheckErrors` utility when appropriate to wrap errors with custom messages.
- **Persistent Storage:** Logs are typically stored in the OS-specific user cache directory under the `AppName` specified in `log.Config`.
- **Roadmap:** See `TODO.md` for v2 features (Structured logging, log rotation, performance).

## Usage Example

```go
package main

import "github.com/TeoDev1611/batman/log"

func main() {
    log.Config.AppName = "MyApp"
    log.Config.FileToLog = "app.log"
    if err := log.Init(); err != nil {
        panic(err)
    }
    
    log.Debug("Debugging info")
    log.Info("Service started")
    log.Error("Something went wrong")
}
```
