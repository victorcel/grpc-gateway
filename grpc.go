package main

import (
	"context"
	personpb "fire-storage/pkg/v1/person"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	personpb.UnimplementedPersonServiceServer
}

func (s *server) GetPerson(_ context.Context, req *personpb.PersonRequest) (*personpb.PersonResponse, error) {
	log.Printf("Received registration for: %s, %s, %d", req.Name, req.Email, req.Age)
	// Simular el almacenamiento de datos
	fmt.Println("Datos personales almacenados correctamente.")
	return &personpb.PersonResponse{
		Status: ".:OK:.",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer func(lis net.Listener) {
		err := lis.Close()
		if err != nil {
			log.Fatalf("failed to close listener: %v", err)
		}
	}(lis)

	grpcServer := grpc.NewServer()
	personpb.RegisterPersonServiceServer(grpcServer, &server{})
	log.Printf("server listening at %s", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
