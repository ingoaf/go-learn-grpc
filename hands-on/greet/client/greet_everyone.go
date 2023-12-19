package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/ingoaf/learning-go-grpc/hands-on/greet/proto"
)

func doGreetEveryone(ctx context.Context, c pb.GreetServiceClient, names []string) error {

	stream, err := c.GreetEveryone(ctx)
	if err != nil {
		return fmt.Errorf("cannot create client stream: %w", err)
	}

	namesRequest := requestFromNames(names)

	errChan := make(chan error)
	go func() {
		for _, req := range namesRequest {
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}

		stream.CloseSend()
	}()

	go func(errChan chan<- error) {
		for {

			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				errChan <- fmt.Errorf("cannot receive server response: %v", err)
				break

			}

			log.Printf("Res: %s", res.Result)
		}

		close(errChan)
	}(errChan)

	err = <-errChan

	return err
}

func requestFromNames(names []string) []*pb.GreetRequest {
	xReq := make([]*pb.GreetRequest, len(names))

	for i := range names {
		xReq[i] = &pb.GreetRequest{
			FirstName: names[i],
		}
	}

	return xReq
}
