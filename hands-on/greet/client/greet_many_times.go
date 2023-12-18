package main

import (
	"context"
	"io"
	"log"

	pb "github.com/ingoaf/learning-go-grpc/hands-on/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient, firstName string) {
	greetRequest := &pb.GreetRequest{
		FirstName: firstName,
	}

	stream, err := c.GreetManyTimes(context.TODO(), greetRequest)
	if err != nil {
		log.Fatalf("Cannot start greeting stream: %v", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		// not EOF, but another error
		if err != nil {
			log.Fatalf("Error while reading greeting stream: %v", err)
		}

		log.Printf(msg.Result)
	}

}
