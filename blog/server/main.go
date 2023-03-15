package main

import (
	"context"
	"log"
	"net"

	pb "github.com/shreeyashnaik/grpc-go/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

var collection *mongo.Collection

type Server struct {
	pb.BlogServiceServer
}

func main() {
	dbClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	err = dbClient.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	collection = dbClient.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Unable to listen: %v\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
