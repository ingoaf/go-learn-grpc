package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/ingoaf/learning-go-grpc/hands-on/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func greetWithDeadline(ctx context.Context, c pb.GreetServiceClient, name string) error {
	res, err := c.GreetWithDeadline(ctx, &pb.GreetRequest{
		FirstName: name,
	})

	if err != nil {
		e, grpcError := status.FromError(err)
		if grpcError && e.Code() == codes.DeadlineExceeded {
			return fmt.Errorf("server took too long to respond, client canceled: %w", err)
		}

		return err
	}

	log.Printf("Response: %v", res.Result)
	return nil
}
