package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/ingoaf/learning-go-grpc/hands-on/greet/proto"
)

var addr string = "localhost:50051"

type Service struct {
	pb.GreetServiceClient
}

func main() {
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()), // disabes SSL, which gRPC needs by default
	)
	if err != nil {
		log.Fatalf("cannot open connection to %s, %v", addr, err)
	}

	c := pb.NewGreetServiceClient(conn)
	// doGreet(c, "Ilija")
	doGreetManyTimes(c, "Ilija")

	defer conn.Close()
}
