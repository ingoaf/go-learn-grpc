package main

import (
	"context"
	"io"
	"log"

	pb "github.com/ingoaf/learning-go-grpc/hands-on/calculator/proto"
)

func calculatePrimeFactorDecomposition(ctx context.Context, grpcClient pb.CalculatorServiceClient, numberToDecompose int64) ([]int64, error) {
	var decomposition []int64
	req := pb.PrimesRequest{
		Number: numberToDecompose,
	}

	c, err := grpcClient.Primes(ctx, &req)
	if err != nil {
		log.Fatalf("cannot calculate decomposition: %v", err)
	}

	for {
		resp, err := c.Recv()
		if err == io.EOF {
			return decomposition, nil
		}

		if err != nil {
			return nil, err
		}

		decomposition = append(decomposition, resp.Number)
	}
}
