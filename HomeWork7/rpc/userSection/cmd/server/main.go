package main

import (
	"HomeWork7/rpc/userSection/controller"
	"HomeWork7/rpc/userSection/dao/db"
	user2 "HomeWork7/rpc/userSection/user"

	"google.golang.org/grpc"
	"log"
	"net"
)

func init() {
	db.InitDb()
}

func main() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Println("listen to net failed ")
		panic(err)
	}
	//创建服务
	// NewServer creates a gRPC server which has no service registered and has not
	// started to accept requests yet.
	rserver := grpc.NewServer()
	lserver := grpc.NewServer()

	//注册服务
	user2.RegisterRServerServer(rserver, &controller.UserRegister{})
	user2.RegisterLServerServer(lserver, &controller.UserLogin{})

	//启动服务
	//Serve方法中是有for循环的
	err1 := rserver.Serve(listener)
	err2 := lserver.Serve(listener)
	if err1 != nil || err2 != nil {
		log.Println("启动服务失败")
		log.Println("err1 : ", err1)
		log.Println("err2 : ", err2)
		return
	}
}
