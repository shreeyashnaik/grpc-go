package main

import (
	"context"
	"io"
	"log"
	"time"

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

	doSum(c)
	doPrimes(c)
	doAvg(c)
	doMax(c)
}

/* UNARY STREAMING */
func doSum(c pb.CalculatorClient) {
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		Num1: 10,
		Num2: 3,
	})
	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}
	log.Printf("Sum: %d\n", res.Ans)
}

/* SERVER STREAMING */
func doPrimes(c pb.CalculatorClient) {
	res, err := c.Primes(context.Background(), &pb.PrimesRequest{
		Num: 120,
	})
	if err != nil {
		log.Fatalf("Could not primes: %v\n", err)
	}
	for {
		ans, err := res.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not primes: %v\n", err)
		}

		log.Println(ans)
	}
}

/* CLIENT STREAMING */
func doAvg(c pb.CalculatorClient) {
	reqs := []*pb.AvgRequest{
		{Num: 1},
		{Num: 2},
		{Num: 4},
	}
	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("Could not avg: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v", req)
		stream.Send(req)
		time.Sleep(time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Could not avg: %v\n", err)
	}
	log.Println("Avg: ", res.Avg)
}

/* BI-DIRECTIONAL STREAMING */
func doMax(c pb.CalculatorClient) {
	reqs := []*pb.MaxRequest{
		{Num: 10},
		{Num: 2},
		{Num: 14},
		{Num: 12},
		{Num: 21},
	}

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while creating client stream: %v\n", err)
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

			log.Printf("Received new max: %v\n", res.Max)
		}

		close(waitc)
	}()

	<-waitc
}
