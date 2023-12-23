package main

import (
	"log"
	"net"

	pb "github.com/ingoaf/learning-go-grpc/hands-on/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	serverCertFile      = "./ssl/server.crt"
	serverPublicKeyFile = "./ssl/server.pem"
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

	// enable TLS
	grpcOpts := []grpc.ServerOption{}
	tls := true
	if tls {
		creds, err := credentials.NewServerTLSFromFile(serverCertFile, serverPublicKeyFile)
		if err != nil {
			log.Fatalf("cannot load server certificate (did you run the script in ssh folder?): %v", err)
		}

		grpcOpts = append(grpcOpts, grpc.Creds(creds))
	}

	// Create our Server
	srv := Server{}

	// Create gRPC server
	s := grpc.NewServer(grpcOpts...)
	pb.RegisterGreetServiceServer(s, &srv)

	// serve listener
	if err := s.Serve(lis); err != nil {
		log.Fatalf("cannot serve listener %v", err)
	}
}
