package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/ingoaf/learning-go-grpc/hands-on/calculator/proto"
)

func (s Server) Average(stream pb.CalculatorService_AverageServer) error {
	totalSum := 0
	count := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			stream.SendAndClose(&pb.AverageResponse{
				Number: float32(totalSum) / float32(count),
			})
			return nil
		}

		if err != nil {
			log.Fatalf("cannot receive value: %v", err)
		}

		fmt.Printf("Received from client: %d \n", req.Number)
		totalSum += int(req.Number)
		count += 1
	}
}
