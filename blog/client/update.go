package main

import (
	"context"
	"log"

	pb "github.com/shreeyashnaik/grpc-go/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("updateBlog was invoked")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Shreeyash, Hitler",
		Title:    "Mein Kampf 2",
		Content:  "I always struggle",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Error while updating: %v\n", err)
	}

	log.Println("Blog was updated!")
}
