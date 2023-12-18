package main

import (
	"log"
	"net"

	pb "github.com/ingoaf/learning-go-grpc/hands-on/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "localhost:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("cannot create listener %v", err)
	}

	srv := Server{}
	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, srv)

	log.Println("Serving listener...")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("cannot serve listener %v", err)
	}

}
