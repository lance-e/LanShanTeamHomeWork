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

	friend := engine.Group("/friendSection")
	friend.Use(middleware.JWT()) //jwt 鉴权
	{
		friend.POST("/requestAdd", controller.RequestAddFriend)
		friend.GET("/receiveAdd", controller.ReceiveAddFriend)
		friend.POST("/delete", controller.DeleteFriend)
		friend.POST("/groupFriend", controller.GroupFriend)
		friend.POST("/showAllFriend", controller.ShowAllFriend)
		friend.POST("/showGroupFriend", controller.ShowGroupFriend)
		friend.POST("/search", controller.SearchFriend)
	}

	return engine
}
