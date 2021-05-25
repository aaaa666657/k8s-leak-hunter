package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

)

type Server struct{}

func main() {
	fmt.Println("starting gPRC seerver...")
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal("failed to listen: %v \n", err)
	}

	grpcServer := grpc.NewServer()
	scannerPD.RegisterResourceRegisterServiceServer(grpcServer,$Server{})
	
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}
}
