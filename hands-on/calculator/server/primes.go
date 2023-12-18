package main

import (
	"fmt"

	pb "github.com/ingoaf/learning-go-grpc/hands-on/calculator/proto"
)

func (s Server) Primes(in *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	n := in.Number
	k := int64(2)

	for {
		if n <= 1 {
			break
		}

		if n%k == 0 {
			n = n / k
			err := stream.Send(
				&pb.PrimesResponse{
					Number: k,
				},
			)
			if err != nil {
				return fmt.Errorf("cannot send factor: %v", err)
			}
		} else {
			k += 1
		}
	}

	return nil
}
