package main

import (
	"log"
	"net"

	pb "github.com/ingoaf/learning-go-grpc/hands-on/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "localhost:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	// create listener
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("cannot start listener %v", err)
	}

	log.Printf("Listening...")

	// Create our Server
	srv := Server{}

	// Create gRPC server
	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &srv)

	// serve listener
	if err := s.Serve(lis); err != nil {
		log.Fatalf("cannot serve listener %v", err)
	}
}
