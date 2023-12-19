package main

import (
	"context"
	"fmt"
	"io"

	pb "github.com/ingoaf/learning-go-grpc/hands-on/calculator/proto"
)

func calculateCurrentMaximum(ctx context.Context, c pb.CalculatorServiceClient, numbers []float64) ([]float64, error) {
	var maximums []float64

	stream, err := c.Max(ctx)
	if err != nil {
		return maximums, fmt.Errorf("cannot create maximum request stream: %w", err)
	}

	errChan := make(chan error)

	go func(c chan<- error) {
		for i := range numbers {
			err = stream.Send(&pb.MaxRequest{
				Number: numbers[i],
			})
			if err != nil {
				c <- err
			}
		}

		err := stream.CloseSend()
		if err != nil {
			c <- err
		}

	}(errChan)

	go func(c chan<- error) {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				c <- err
				break
			}

			maximums = append(maximums, resp.Maximum)
		}

		close(c)
	}(errChan)

	err = <-errChan
	if err != nil {
		return []float64{}, err
	}

	return maximums, nil
}
