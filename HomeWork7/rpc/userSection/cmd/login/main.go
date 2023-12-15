package main

import (
	"HomeWork7/rpc/userSection/controller"
	user2 "HomeWork7/rpc/userSection/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	listener2, errs := net.Listen("tcp", "localhost:9091")
	if errs != nil {
		log.Println("listen to net failed ")
		panic(errs)
	}
	lserver := grpc.NewServer()
	user2.RegisterLServerServer(lserver, &controller.UserLogin{})

	log.Println("=========登陆服务启动！===========")

	err2 := lserver.Serve(listener2)
	if err2 != nil {
		log.Println("log serve failed")
		return
	}
}
