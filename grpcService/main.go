package main

import (
	grpc "google.golang.org/grpc"
	carService "grpcChassis/grpcService/api/ilya225/car-service"
	"log"
	"net"
)

const (
	port = ":5000"
)

type router struct {
	carService.UnimplementedCarServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	carService.RegisterCarServiceServer(s, &router{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
