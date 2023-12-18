package main

import (
	"context"
	"log"

	pb "github.com/ingoaf/learning-go-grpc/hands-on/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

type Service struct {
	pb.CalculatorServiceClient
}

func main() {
	securityOptions := grpc.WithTransportCredentials(insecure.NewCredentials())

	conn, err := grpc.Dial(addr, securityOptions)
	if err != nil {
		log.Fatalf("cannot dial connection to address %s: %v", addr, err)
	}

	// sumServiceClient := pb.NewCalculatorServiceClient(conn)
	// req := pb.SumRequest{
	// 	FirstNumber:  4,
	// 	SecondNumber: 5,
	// }

	// res, err := sumServiceClient.Sum(context.Background(), &req)
	// if err != nil {
	// 	log.Fatalf("cannot receive request from gRPC server: %v", err)
	// }

	var numberToDecompose int64 = 120
	calculatorServiceClient := pb.NewCalculatorServiceClient(conn)

	decomposition, err := calculatePrimeFactorDecomposition(context.Background(), calculatorServiceClient, numberToDecompose)
	if err != nil {
		log.Fatalf("cannot compute prime decomposition of %d, %v", numberToDecompose, err)
	}

	log.Printf("The response is: %v", decomposition)
}
