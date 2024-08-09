package main

import (
	handler "grpc-kitchen/services/orders/handler/orders"
	"grpc-kitchen/services/orders/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	ordersService := service.NewOrderServce()
	handler.NewGrpcOrdersService(grpcServer, ordersService)

	return grpcServer.Serve(lis)
}
