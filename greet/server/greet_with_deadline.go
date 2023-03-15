package main

import (
	"context"
	"log"
	"time"

	pb "github.com/shreeyashnaik/grpc-go/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline was invoked with %v\n", in)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Client cancelled the request")
			return nil, status.Error(codes.Canceled, "Client cancelled the request")
		}

		time.Sleep(time.Second)
	}

	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
