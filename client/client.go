package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	personpb "fire-storage/pkg/v1/person"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
