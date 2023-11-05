package main

import (
	"log"
	"net"
	server "proto-docs-service/internal"

	"google.golang.org/grpc"
	pb "proto-docs-service/grpc/docs"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterDocumentServiceServer(s, server.GetNewServer())
	
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}