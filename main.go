package main

import (
	"context"
	"fmt"
	pb "go-design-performance-monitoring/proto/counter"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const grpcServerPort = "50055"

type CounterServer struct {
	pb.UnimplementedIncrementCounterServer
	value int32
}

func (s *CounterServer) Increment(ctx context.Context, req *pb.IncrementBy) (*pb.Status, error) {
	fmt.Println("Got request to increment counter")
	if req.Value < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Value cannot be negative")
	}
	s.value += req.Value
	msg := fmt.Sprintf("Incremented value is %d", s.value)
	return &pb.Status{Message: msg}, nil
}

func main() {
	fmt.Println("Inside main")
	lis, err := net.Listen("tcp", ":"+grpcServerPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterIncrementCounterServer(s, &CounterServer{})
	log.Println("Server listening on port ", grpcServerPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Printf("Exiting main")
}

