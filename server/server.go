package server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc-http-server/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements the gRPC service
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received request from: %s", req.GetName())
	return &pb.HelloReply{Message: fmt.Sprintf("Hello, %s!", req.GetName())}, nil
}

// StartServer starts the gRPC server
func StartServer(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	log.Printf("Server listening on port %d", port)
	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}
