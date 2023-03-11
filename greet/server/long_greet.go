package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/shreeyashnaik/grpc-go/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet func was invoked!")

	res := ""

	for {
		req, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&pb.GreetResponse{
					Result: res,
				})
			}
			log.Fatalf("Error while reading client stream")
		}

		log.Printf("Receiving: %v\n", req)
		res += fmt.Sprintf("Hello %s!", req.FirstName)
	}

}
