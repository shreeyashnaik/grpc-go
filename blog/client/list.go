package main

import (
	"context"
	"io"
	"log"

	pb "github.com/shreeyashnaik/grpc-go/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlog(c pb.BlogServiceClient) {
	log.Println("listBlog was invoked")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error while calling listBlogs: %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Failed to connect: %v\n", err)
		}

		log.Println(res)
	}

}
