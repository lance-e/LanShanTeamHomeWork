package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "serve from the docker container ",
			"code": 200,
		})
	})
	engine.Run(":8080")
}
