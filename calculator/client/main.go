package main

import (
	"context"
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

	// res, err := c.Sum(context.Background(), &pb.SumRequest{
	// 	Num1: 10,
	// 	Num2: 3,
	// })
	// if err != nil {
	// 	log.Fatalf("Could not sum: %v\n", err)
	// }
	// log.Printf("Sum: %d\n", res.Ans)

	// res, err := c.Primes(context.Background(), &pb.PrimesRequest{
	// 	Num: 120,
	// })
	// if err != nil {
	// 	log.Fatalf("Could not primes: %v\n", err)
	// }
	// for {
	// 	ans, err := res.Recv()

	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatalf("Could not primes: %v\n", err)
	// 	}

	// 	log.Println(ans)
	// }

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
