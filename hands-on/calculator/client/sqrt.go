package main

import (
	"context"
	"fmt"

	pb "github.com/ingoaf/learning-go-grpc/hands-on/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func calculateSquareRoot(ctx context.Context, c pb.CalculatorServiceClient, number int32) (float64, error) {
	res, err := c.Sqrt(ctx, &pb.SqrtRequest{Number: number})
	if err != nil {
		e, isGrpcError := status.FromError(err)

		if isGrpcError {
			if e.Code() == codes.InvalidArgument {
				return 0, fmt.Errorf("provide valid input for sqrt: %w", err)
			}
		}
	}

	return res.Number, err
}
