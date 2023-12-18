package main

import (
	"context"
	"fmt"

	pb "github.com/ingoaf/learning-go-grpc/hands-on/calculator/proto"
)

func calculateAverage(ctx context.Context, c pb.CalculatorServiceClient, numbers []int64) (float64, error) {
	stream, err := c.Average(ctx)
	if err != nil {
		return 0, fmt.Errorf("cannot stream requests: %v", err)
	}

	for i := range numbers {
		if err := stream.Send(&pb.AverageRequest{Number: numbers[i]}); err != nil {
			return 0, fmt.Errorf("can not send number %d: %v", numbers[i], err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return 0, fmt.Errorf("cannot receive response: %v", err)
	}

	return float64(res.Number), nil
}
