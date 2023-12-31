package main

import "RedRockHomework6/route"

func main() {
	engine := route.NewRouter()
	engine.Run(":8080")

}
