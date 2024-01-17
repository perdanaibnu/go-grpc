package main

import (
	"log"
	"net"

	"go-grpc/cmd/config"
	"go-grpc/cmd/services"
	productPb "go-grpc/pb/product"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	netListen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen %v", err.Error())
	}

	db := config.ConnectDB()
	grpcServer := grpc.NewServer()
	productService := services.ProductService{DB: db}
	productPb.RegisterProductServiceServer(grpcServer, &productService)

	log.Printf("Server started at %v", netListen.Addr())
	if err := grpcServer.Serve(netListen); err != nil {
		log.Fatalf("failed to server %v", err.Error())
	}
}
