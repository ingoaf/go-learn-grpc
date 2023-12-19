package main

import (
	"io"

	pb "github.com/ingoaf/learning-go-grpc/hands-on/calculator/proto"
)

func (s Server) Max(stream pb.CalculatorService_MaxServer) error {
	maximums := []float64{}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		if len(maximums) == 0 || (maximums[len(maximums)-1] <= req.Number) {
			maximums = append(maximums, req.Number)
		} else {
			maximums = append(maximums, maximums[len(maximums)-1])
		}

		stream.Send(&pb.MaxResponse{
			Maximum: maximums[len(maximums)-1],
		})
	}
}
