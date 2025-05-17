package server

import (
	"context"
	"testing"

	pb "grpc-http-server/proto"
)

func TestSayHello(t *testing.T) {
	s := &server{}
	resp, err := s.SayHello(context.Background(), &pb.HelloRequest{Name: "TestUser"})
	if err != nil {
		t.Fatalf("SayHello failed: %v", err)
	}
	want := "Hello, TestUser!"
	if resp.Message != want {
		t.Errorf("unexpected message: got %q, want %q", resp.Message, want)
	}
}
