package main

import (
	"context"
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

	// uncomment the following line to enable unary request
	// doGreet(c, "Ilija")

	// uncomment the following line to enable server streaming request
	// doGreetManyTimes(c, "Ilija")

	// uncomment the following line to enable client streaming request
	// err = doLongGreet(
	// 	context.Background(),
	// 	c,
	// 	[]string{"Tim", "Anna", "Jon", "Jack", "Jill", "Flo", "Hun", "Alma", "Nat", "Em"},
	// )
	// if err != nil {
	// 	log.Fatalf("cannot do long greet: %v", err)
	// }

	// uncomment the following line to enable bi-directional streaming request
	err = doGreetEveryone(context.Background(), c, []string{"Tom", "Emma", "Paul"})
	if err != nil {
		log.Fatalf("cannot greet everyone: %v", err)
	}

	defer conn.Close()
}
