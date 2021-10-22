package main

import (
	"catalog/cmd/grpc/service"
	"catalog/proto"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Hey, I'm running")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := service.CatalogServer{}

	grpcServer := grpc.NewServer()

	proto.RegisterCatalogServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
