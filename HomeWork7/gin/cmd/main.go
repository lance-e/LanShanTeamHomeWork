package main

import (
	"HomeWork7/gin/route"
	"HomeWork7/gin/util"
)

func init() {
	util.ConnectToRpc()
}
func main() {

	r := route.NewRouter()
	r.Run(":8081")
	defer util.Conn1.Close()
	defer util.Conn2.Close()
}
