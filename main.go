package main

import (
	"context"
	"flag"
	"fmt"
	"go-design-performance-monitoring/mutex"
	pb "go-design-performance-monitoring/proto/counter"
	"log"
	"net"

	inputQ "go-design-performance-monitoring/inputQ"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const grpcServerPort = "50055"
var program_mode string

type CounterServer struct {
	pb.UnimplementedIncrementCounterServer
}

func (s *CounterServer) Increment(ctx context.Context, req *pb.IncrementBy) (*pb.Status, error) {
	fmt.Println("Got request to increment counter")
	if req.Value < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Value cannot be negative")
	}
	var result int;
	if program_mode == "mutex" {
		result = mutex.Increment(int(req.Value))
	} else if program_mode == "inputQ" {
		result = inputQ.Increment(int(req.Value))
	}
	msg := fmt.Sprintf("Incremented value is %d", result)
	return &pb.Status{Message: msg}, nil
}

func main() {
	fmt.Println("Inside main")
	var mode *string = flag.String("mode", "", "Specify in which mode to run the server. Possible values are 'mutex' or 'inputQ'")
	flag.Parse()
	fmt.Println("mode: ", program_mode)
	program_mode = *mode
	fmt.Println(program_mode != "mutex")
	if program_mode == "" || (program_mode != "inputQ" && program_mode != "mutex") {
		log.Fatalf("Specify in which mode to run the server. Possible values are 'mutex' or 'inputQ'")
	}
	fmt.Println("Running the program in mode: ", program_mode)
	if program_mode == "inputQ" {
		go inputQ.ProcessInputQ()
	}
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

