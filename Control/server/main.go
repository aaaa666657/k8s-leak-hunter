package main

import (
	"context"
	"fmt"

	"control/pkg/db"
	"control/pkg/scanner"
	scannerPB "control/proto/scanner"
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
	//db.InitDB()
	err := db.RegisterService(8, 443, "https")
	if err != nil {
		fmt.Printf("failed to RegisterService: %v \n", err)
	}
	err = db.RegisterHost("nas.local", "192.168.100.209")
	if err != nil {
		fmt.Printf("failed to RegisterHost: %v \n", err)
	}

	scannerTypeRes, scannerErrorRes := scanner.ScannerService(7)

	for i := 0; i < len(scannerErrorRes.ErrorDiffService); i++ {
		fmt.Printf("scaneer port: %d ", scannerErrorRes.ErrorDiffService[i].Port)
		fmt.Println(" service: ", scannerErrorRes.ErrorDiffService[i].Servicetype)
	}
	fmt.Printf("------errorDiffService----------------------------------\n")

	for i := 0; i < len(scannerErrorRes.ErrorPortWithoutExist); i++ {
		fmt.Printf("scaneer port: %d ", scannerErrorRes.ErrorPortWithoutExist[i].Port)
		fmt.Println(" service: ", scannerErrorRes.ErrorPortWithoutExist[i].Servicetype)
	}
	fmt.Printf("--------errorPortWithoutExist--------------------------------\n")

	fmt.Printf("scanner Result : %d \n", scannerTypeRes)
	//err := db.RegisterHost("192.168.100.50")
	//datatype := []db.Service{}
	/* datatype, _ := db.LoadService("192.168.100.50")
	for i := 0; i < len(datatype); i++ {
		fmt.Printf("port: %d ", datatype[i].Port)
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
	} */
}
