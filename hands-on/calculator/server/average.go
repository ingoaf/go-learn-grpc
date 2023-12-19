package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/ingoaf/learning-go-grpc/hands-on/calculator/proto"
)

func (s Server) Average(stream pb.CalculatorService_AverageServer) error {
	var totalSum, count float64

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			stream.SendAndClose(&pb.AverageResponse{
				Number: totalSum / count,
			})
			return nil
		}

		if err != nil {
			log.Fatalf("cannot receive value: %v", err)
		}

		fmt.Printf("Received from client: %d \n", req.Number)
		totalSum += float64(req.Number)
		count += 1
	}
}
