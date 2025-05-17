.PHONY: proto build run clean

# Generate protobuf code
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/service.proto

# Build the project
build: proto
	go build -o bin/server main.go

# Run the server
run: build
	./bin/server

# Clean build artifacts
clean:
	rm -rf bin/ 