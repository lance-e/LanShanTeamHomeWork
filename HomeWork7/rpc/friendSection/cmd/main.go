package main

import (
	"HomeWork7/rpc/friendSection/controller"
	"HomeWork7/rpc/friendSection/friend"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:9092")
	if err != nil {
		log.Println("监听失败")
		panic(err)
	}
	friendServer := grpc.NewServer()
	friend.RegisterFriendServer(friendServer, &controller.FriendServer{})
	err = friendServer.Serve(listener)
	if err != nil {
		log.Println("启动服务失败")
		return
	}
	log.Println("=========好友服务启动！！！！===========")
}
