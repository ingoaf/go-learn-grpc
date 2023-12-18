package main

import (
	"fmt"

	pb "github.com/ingoaf/learning-go-grpc/hands-on/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	var err error

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("I greet %s for the %d time", in.FirstName, i)

		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})
		if err != nil {
			return err
		}
	}

	return err
}
