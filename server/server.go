package server

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

func Run() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	Register(s)
	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8080")

	log.Fatalln(s.Serve(lis))
}
