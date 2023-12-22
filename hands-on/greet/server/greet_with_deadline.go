package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/ingoaf/learning-go-grpc/hands-on/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	time.Sleep(3 * time.Second)
	if ctx.Err() == context.DeadlineExceeded {
		log.Printf("error: %v", status.Error(codes.Canceled, "client closed connection"))
		return nil, status.Error(codes.Canceled, "client closed connection")
	}

	return &pb.GreetResponse{
		Result: fmt.Sprintf("Hello %s", in.FirstName),
	}, nil
}
