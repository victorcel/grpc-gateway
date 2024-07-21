package main

import (
	"context"
	personpb "github.com/victorcel/grpc-gateway-proto/pkg/v1/person"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	personpb.UnimplementedPersonServiceServer
}

func (s *server) GetPerson(_ context.Context, req *personpb.PersonRequest) (*personpb.PersonResponse, error) {
	//log.Printf("Received registration for: %s, %s, %d", req.Name, req.Email, req.Age)
	//fmt.Println("Datos personales almacenados correctamente.")

	if req.Age > 100 {
		return nil, status.Errorf(codes.Unavailable, "Edad no permitida.")
	}

	return &personpb.PersonResponse{
		Status: "Registro exitoso.",
	}, nil

}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()

	reflection.Register(s)

	// Attach the Greeter service to the server
	personpb.RegisterPersonServiceServer(s, &server{})
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:9090")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.NewClient(
		"0.0.0.0:8081",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()

	// Register Greeter
	err = personpb.RegisterPersonServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8091",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8081")
	log.Fatalln(gwServer.ListenAndServe())
}
