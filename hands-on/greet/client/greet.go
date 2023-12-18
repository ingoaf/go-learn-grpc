package main

import (
	"context"
	"log"

	pb "github.com/ingoaf/learning-go-grpc/hands-on/greet/proto"
)

func doGreet(c pb.GreetServiceClient, firstName string) {
	req := pb.GreetRequest{
		FirstName: firstName,
	}

	res, err := c.Greet(context.TODO(), &req)
	if err != nil {
		log.Fatalf("cannot greet: %v", err)
	}

	log.Println(res.Result)
}
