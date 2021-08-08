package main

import (
	"context"
	"fmt"
	"log"

	"github.com/diogofgomespinheiro/fc-grpc-examples/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:5051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to the grpc server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	AddUser(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Diogo Pinheiro",
		Email: "diogo@test.test",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make grpc request: %v", err)
	}

	fmt.Println(res)
}
