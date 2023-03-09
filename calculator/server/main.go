package main

import (
	"context"
	"log"
	"net"

	pb "github.com/shreeyashnaik/grpc-go/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50052"

type Server struct {
	pb.CalculatorServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Unable to listen")
	}

	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	return &pb.SumResponse{
		Ans: (in.Num1 + in.Num2),
	}, nil
}
