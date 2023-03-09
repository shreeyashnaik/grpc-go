package main

import (
	"context"
	"log"

	pb "github.com/shreeyashnaik/grpc-go/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50052"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect on %s, %v\n", addr, err)
	}

	defer conn.Close()

	c := pb.NewCalculatorClient(conn)

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		Num1: 10,
		Num2: 3,
	})
	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Sum: %d\n", res.Ans)
}
