package main

import (
	"fmt"
	"io"

	pb "github.com/ingoaf/learning-go-grpc/hands-on/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	var greetingMessage string = "Hallo "

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			stream.SendAndClose(&pb.GreetResponse{
				Result: greetingMessage[0 : len(greetingMessage)-1],
			})
			return nil
		}

		if err != nil {
			return fmt.Errorf("cannot receive message: %v", err)
		}

		greetingMessage = fmt.Sprintf("%s %s,", greetingMessage, req.FirstName)
	}
}
