package controller

import (
	"HomeWork7/gin/controller/user"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var client user.RServerClient
	client.Register()
}

//.....
