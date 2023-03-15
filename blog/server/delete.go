package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/shreeyashnaik/grpc-go/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*emptypb.Empty, error) {
	log.Println("DeleteBlog was invoked")

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	res, err := collection.DeleteOne(
		ctx,
		primitive.M{"_id": oid},
	)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Could not delete: %v", err),
		)
	}
	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find blog with ID",
		)
	}

	return &emptypb.Empty{}, nil
}
