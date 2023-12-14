package main

import (
	"HomeWork7/gin/dao/db"
	"HomeWork7/gin/route"
	"net"
	"net/rpc"
)

func init() {
	db.Init() //初始化数据库
}
func main() {
	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		panic("connect to rpc failed")
	}
	defer conn.Close()
	client := rpc.NewClient(conn)

	r := route.NewRouter()
	r.Run(":8081")
}
