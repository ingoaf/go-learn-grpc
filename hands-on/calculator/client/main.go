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

	calculatorServiceClient := pb.NewCalculatorServiceClient(conn)

	// var numberToDecompose int64 = 120
	// decomposition, err := calculatePrimeFactorDecomposition(context.Background(), calculatorServiceClient, numberToDecompose)
	// if err != nil {
	// 	log.Fatalf("cannot compute prime decomposition of %d, %v", numberToDecompose, err)
	// }
	// log.Printf("The response is: %v", decomposition)

	// var numbersForAverageCalculation []int64 = []int64{1, 2, 3, 4}
	// average, err := calculateAverage(context.Background(), calculatorServiceClient, numbersForAverageCalculation)
	// if err != nil {
	// 	log.Fatalf("cannot calculate average of %v: %v", numbersForAverageCalculation, err)
	// }
	// log.Printf("The average is: %f", average)

	// numbersForMaximumCalculation := []float64{1, 5, 3, 6, 2, 20}
	// maximums, err := calculateCurrentMaximum(
	// 	context.Background(), calculatorServiceClient, numbersForMaximumCalculation,
	// )
	// if err != nil {
	// 	log.Fatalf("can not calculate maximums: %v", maximums)
	// }
	// log.Printf("\n Numbers: %v \n Maximums: %v", numbersForMaximumCalculation, maximums)

	var number int32 = -2
	sqrt, err := calculateSquareRoot(context.Background(), calculatorServiceClient, number)
	if err != nil {
		log.Fatalf("can not calculate square root: %v", err)
	}
	log.Printf("The square root of %d is %f", number, sqrt)

}
