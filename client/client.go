package client

import (
	"context"
	"fmt"
	"time"

	pb "grpc-http-server/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Client represents a gRPC client
type Client struct {
	conn   *grpc.ClientConn
	client pb.GreeterClient
}

// NewClient creates a new gRPC client
func NewClient(address string) (*Client, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %v", err)
	}

	client := pb.NewGreeterClient(conn)
	return &Client{
		conn:   conn,
		client: client,
	}, nil
}

// SayHello sends a greeting request to the server
func (c *Client) SayHello(name string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.client.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		return "", fmt.Errorf("could not greet: %v", err)
	}

	return resp.GetMessage(), nil
}

// Close closes the client connection
func (c *Client) Close() error {
	return c.conn.Close()
}
