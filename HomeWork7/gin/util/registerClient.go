package util

import (
	"HomeWork7/gin/controller/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var (
	Rclient user.RServerClient
	Lcient  user.LServerClient
)

func InitRPC() {
	conn, err := grpc.Dial("localhost:9090", grpc.DialOption(grpc.WithTransportCredentials(insecure.NewCredentials())))
	if err != nil {
		panic("connect to rpc failed")
	}
	defer conn.Close()
	log.Println("正在监听rpc服务")

	Rclient = user.NewRServerClient(conn)
	Lcient = user.NewLServerClient(conn)
}
