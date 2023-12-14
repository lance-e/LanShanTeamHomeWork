package main

import (
	"HomeWork7/rpc/user"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type UserRegister struct {
	user.UnimplementedRServerServer
}

func (u *UserRegister) Register(ctx context.Context, info *user.UserInfo) (*user.Response, error) {
	//注册业务
	fmt.Println("注册成功")
	return &user.Response{
		Result: true,
	}, nil
}

type UserLogin struct {
	user.UnimplementedLServerServer
}

func (u *UserLogin) Login(context.Context, *user.UserInfo) (*user.Response, error) {
	//登陆业务
	fmt.Println("登陆成功")
	return &user.Response{
		Result: true,
	}, nil
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
	user.RegisterRServerServer(rserver, &UserRegister{})
	user.RegisterLServerServer(lserver, &UserLogin{})

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
