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
	Conn1   *grpc.ClientConn
	Conn2   *grpc.ClientConn
	err     error
)

func ConnectToRpc() {
	Conn1, err = grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("connect to rpc failed")
	}
	Conn2, err = grpc.Dial("localhost:9091", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("connect to rpc failed")
	}
	//defer Conn.Close() // 不要这样写，会报错，defer是有作用域的，init之后，就直接关闭了，此时main函数都还没开始
	log.Println("正在监听rpc服务")

	Rclient = user.NewRServerClient(Conn1)
	Lcient = user.NewLServerClient(Conn2)

}
