package main

import (
	"context"
	"fmt"
	personpb "github.com/victorcel/grpc-gateway-proto/pkg/v1/person"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	client := personpb.NewPersonServiceClient(conn)

	person := &personpb.PersonRequest{
		Name:  "Victor Barrera",
		Email: "juan.perez@example.com",
		Age:   30,
	}

	response, err := client.GetPerson(context.Background(), person)
	if err != nil {
		log.Fatalf("failed to register: %v", err)
	}
	fmt.Println(response.Status)
}
