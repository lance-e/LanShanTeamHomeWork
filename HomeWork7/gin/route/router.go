package route

import (
	"HomeWork7/gin/controller"
	"HomeWork7/gin/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.Cors()) //防止跨域问题

	engine.POST("/user/register", controller.Register)
	engine.POST("/user/login", controller.Login)

	g := engine.Group("/other")
	g.Use(middleware.JWT()) //jwt 鉴权
	{
		g.GET("/testToken", controller.TestToken)
	}

	return engine
}
