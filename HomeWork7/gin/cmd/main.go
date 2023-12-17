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

	defer func() {
		util.Conn1.Close()
		util.Conn2.Close()
	}()

}
