# Log Parser System (Checkpoint 1)

A modular log parsing tool written in Go that extracts structured data from raw log strings.

## Architecture
The project follows standard Go project layout and clean architecture principles:
- **Domain**: Contains core business models (`Task`, `Result`) and lifecycle constants.
- **Storage**: In-memory storage implementing the `Storage` interface.
- **Parser**: Handles raw string validation and data extraction.
- **Manager**: Orchestrates the task lifecycle and links storage with the parser.

## Task Lifecycle
The system strictly enforces the following state transitions:
1. **PENDING**: Task is initialized with a raw log string and saved to memory.
2. **PROCESSING**: The manager fetches the task and changes its state. Repeated processing attempts on the same task are blocked (`ErrInvalidState`).
3. **COMPLETED**: The parser successfully extracts the `Level` and `Message`, saving the result back to the task.

## How to run
Run the main application:
`go run main.go`

Run all unit tests:
`go test ./...`