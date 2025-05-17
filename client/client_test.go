package client

import (
	"context"
	"net"
	"testing"

	pb "grpc-http-server/proto"

	"google.golang.org/grpc"
)

type testServer struct {
	pb.UnimplementedGreeterServer
}

func (s *testServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello, " + req.Name + "!"}, nil
}

func startGRPCServer(t *testing.T) (addr string, stop func()) {
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &testServer{})
	go s.Serve(lis)
	return lis.Addr().String(), func() {
		s.Stop()
		lis.Close()
	}
}

func TestClient_SayHello(t *testing.T) {
	addr, stop := startGRPCServer(t)
	defer stop()

	client, err := NewClient(addr)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	resp, err := client.SayHello("TestClient")
	if err != nil {
		t.Fatalf("SayHello failed: %v", err)
	}
	want := "Hello, TestClient!"
	if resp != want {
		t.Errorf("unexpected message: got %q, want %q", resp, want)
	}
}
