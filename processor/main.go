package main

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/viper"
	grpc "google.golang.org/grpc"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	port := "8082"
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	RegisterCVServiceServer(grpcServer, newServer())
	fmt.Printf("Listening on port %s...", port)
	grpcServer.Serve(lis)
}
