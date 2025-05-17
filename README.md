# gRPC HTTP Server

A simple gRPC server and client implementation in Go.

## Prerequisites

- Go 1.21 or later
- Protocol Buffers compiler (protoc)
- Go plugins for protoc:
  - `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
  - `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

## Setup

1. Install dependencies:
   ```bash
   go mod download
   ```

2. Generate protobuf code:
   ```bash
   make proto
   ```

## Building and Running

1. Build the project:
   ```bash
   make build
   ```

2. Run the server:
   ```bash
   make run
   ```

The server will start on port 50051, and the client will automatically make a test request.

## Project Structure

- `proto/` - Protocol Buffer definitions
- `server/` - gRPC server implementation
- `client/` - gRPC client implementation
- `main.go` - Main application entry point
- `Makefile` - Build and run commands
