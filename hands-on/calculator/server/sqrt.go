package main

import (
	"context"
	"fmt"
	"math"

	pb "github.com/ingoaf/learning-go-grpc/hands-on/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s Server) Sqrt(ctx context.Context, req *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	if req.Number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("received negative number %d", req.Number),
		)
	}

	return &pb.SqrtResponse{
		Number: math.Sqrt(float64(req.Number)),
	}, nil
}
