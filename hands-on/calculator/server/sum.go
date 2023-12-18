package main

import (
	"context"

	pb "github.com/ingoaf/learning-go-grpc/hands-on/calculator/proto"
)

func (s Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	res := in.FirstNumber + in.SecondNumber
	return &pb.SumResponse{
		Result: res,
	}, nil
}
