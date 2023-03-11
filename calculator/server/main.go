package main

import (
	"context"
	"io"
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

func (s *Server) Primes(in *pb.PrimesRequest, stream pb.Calculator_PrimesServer) error {
	n := int(in.Num)
	factor := 2
	for n > 1 {
		if n%factor == 0 {
			n /= factor
			stream.Send(&pb.PrimesResponse{
				Factor: int32(factor),
			})
		} else {
			factor += 1
		}
	}
	return nil
}

func (s *Server) Avg(stream pb.Calculator_AvgServer) error {
	log.Println("Avg func invoked by Server")

	var avg float32 = 0.0
	i := 0
	for {
		req, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&pb.AvgResponse{
					Avg: avg / float32(i),
				})
			}
			log.Fatalf("Unable to Avg:%v\n", avg)
		}

		avg += float32(req.Num)
		i++
	}
}
