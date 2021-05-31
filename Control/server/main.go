package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"strings"

	scannerPB "control/proto/scanner"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

const (
	userName = "root"
	password = "scanner"
	host     = "127.0.0.1"
	port     = "3306"
	dbName   = "scanner"
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
	path := strings.Join([]string{userName, ":", password, "@tcp(", host, ":", port, ")/", dbName, "?charset=utf8"}, "")
	fmt.Println(path)

	DB, _ := sql.Open("mysql", path)

	// 設定 database 最大連接數
	DB.SetConnMaxLifetime(100)

	//設定上 database 最大閒置連接數
	DB.SetMaxIdleConns(10)

	// 驗證是否連上 db
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail:", err)
		return
	}
	fmt.Println("connnect success")
	//gRPC Server
	fmt.Println("starting gPRC seerver...")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		//log.Fatal("failed to listen: %v \n", err)
	}

	grpcServer := grpc.NewServer()
	scannerPB.RegisterResourceRegisterServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}
}
