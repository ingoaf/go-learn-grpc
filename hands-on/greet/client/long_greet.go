package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/ingoaf/learning-go-grpc/hands-on/greet/proto"
)

func doLongGreet(ctx context.Context, c pb.GreetServiceClient, names []string) error {
	stream, err := c.LongGreet(ctx)
	if err != nil {
		log.Fatalf("cannot call rpc method LongGreet: %v", err)
	}

	for i := 0; i < 10; i++ {
		err := stream.Send(&pb.GreetRequest{
			FirstName: names[i],
		})
		if err != nil {
			return fmt.Errorf("cannot send name %s", names[i])
		}

		log.Printf("Sent request for %s", names[i])
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("cannot close stream, did not receive response: %v", err)
	}

	log.Printf("Received response from server: %s \n", res.Result)

	return nil
}
