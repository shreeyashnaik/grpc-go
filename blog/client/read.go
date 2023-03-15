package main

import (
	"context"
	"log"

	pb "github.com/shreeyashnaik/grpc-go/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("readBlog was invoked")

	req := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Printf("Error while reading: %v\n", err)
	}

	log.Printf("Read Blog successfully: %v\n", res)
	return res
}
