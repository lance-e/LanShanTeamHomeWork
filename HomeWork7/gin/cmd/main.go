package main

import (
	"HomeWork7/gin/route"
	"HomeWork7/gin/util"
)

func init() {
	util.InitRPC()
}

func main() {

	r := route.NewRouter()
	r.Run(":8081")
}
