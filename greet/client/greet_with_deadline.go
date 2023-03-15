package main

import (
	"context"
	"log"
	"time"

	pb "github.com/shreeyashnaik/grpc-go/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Shreeyash",
	}

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded!")
				return
			} else {
				log.Fatalf("Unexpected GRPC error: %v\n", err)
			}
		} else {
			log.Fatalf("NON-GRPC error: %v\n", err)
		}
	}

	log.Printf("Greet With Deadline: %s\n", res.Result)
}
