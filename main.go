package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"grpc-http-server/client"
	"grpc-http-server/server"
)

func main() {
	// Start the server in a goroutine
	go func() {
		if err := server.StartServer(50051); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Create a client
	c, err := client.NewClient("localhost:50051")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer c.Close()

	// Make a request
	response, err := c.SayHello("World")
	if err != nil {
		log.Fatalf("Failed to say hello: %v", err)
	}
	log.Printf("Response: %s", response)

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}
