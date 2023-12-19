package main

import (
	"fmt"
	"io"

	pb "github.com/ingoaf/learning-go-grpc/hands-on/greet/proto"
)

func (s Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return fmt.Errorf("cannot receive client stream: %w", err)
		}

		res := fmt.Sprintf("Hello %s !", req.FirstName)

		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})
		if err != nil {
			return fmt.Errorf("cannot send greeting response to client: %w", err)
		}
	}
}
