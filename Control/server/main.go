package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"control/pkg/db"
	scannerPB "control/proto/scanner"

	"google.golang.org/grpc"
)

type Server struct{}

func (*Server) Register(ctx context.Context, req *scannerPB.ResourceRegister) (*scannerPB.ResourceRegisterResult, error) {
	fmt.Printf("Got Data %v \n", req)

	port := req.GetPort()
	service := req.GetServiceType()

	res := &scannerPB.ResourceRegisterResult{
		Result: true,
	}

	fmt.Printf("Got Dataus %d %s  \n", port, service)

	return res, nil
}

func main() {
	//DB
	db.InitDB()
	db.RegisterService("192.168.100.50", 80, "https")
	//datatype := []db.Service{}
	datatype, _ := db.LoadService("192.168.100.50")
	for i := 0; i < len(datatype); i++ {
		fmt.Printf("port: %d", datatype[i].Port)
		fmt.Println("service: ", datatype[i].Servicetype)
	}

	//gRPC Server
	fmt.Println("starting gPRC seerver...")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("failed to listen: %v \n", err)
	}

	grpcServer := grpc.NewServer()
	scannerPB.RegisterResourceRegisterServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}
}
