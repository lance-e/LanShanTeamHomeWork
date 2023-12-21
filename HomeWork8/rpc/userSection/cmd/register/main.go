package main

import (
	"HomeWork7/rpc/userSection/controller"
	user2 "HomeWork7/rpc/userSection/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	//创建服务
	// NewServer creates a gRPC login which has no service registered and has not
	// started to accept requests yet.
	listener1, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Println("listen to net failed ")
		panic(err)
	}
	rserver := grpc.NewServer()
	user2.RegisterRServerServer(rserver, &controller.UserRegister{})

	log.Println("========注册服务启动！！！！！！============")

	err1 := rserver.Serve(listener1)
	if err1 != nil {
		log.Println("register serve failed")
		return
	}
	//注册服务

}
