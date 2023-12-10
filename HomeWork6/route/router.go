package route

import (
	"RedRockHomework6/handle"
	"RedRockHomework6/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	engine := gin.Default()
	engine.POST("/register", handle.Register)
	engine.POST("/login", handle.Login)
	engine.POST("/question", handle.CheckQuestion)
	engine.POST("/findPassword", handle.FindPassword)
	auth := engine.Group("/message")
	auth.Use(middleware.JWT()) //使用jwt中间件
	{
		auth.GET("/list", handle.GetMessageList)
		auth.POST("/sent", handle.SendMessage)
	}

	return engine
}
