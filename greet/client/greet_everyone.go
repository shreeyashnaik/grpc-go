package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/shreeyashnaik/grpc-go/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryOne was invoked!")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating client stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Shreeyash"},
		{FirstName: "Marie"},
		{FirstName: "Test"},
	}

	waitc := make(chan struct{})
	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			stream.Send(req)
			time.Sleep(time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err != nil {
				if err == io.EOF {
					break
				}
				log.Printf("Error while receiving: %v\n", err)
			}

			log.Printf("Received: %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
