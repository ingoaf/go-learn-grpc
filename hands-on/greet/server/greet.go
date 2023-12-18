package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/ingoaf/learning-go-grpc/hands-on/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v", in)
	greeting := fmt.Sprintf("Hello %s", in.FirstName)

	return &pb.GreetResponse{
		Result: greeting,
	}, nil
}
