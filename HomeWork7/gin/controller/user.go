package controller

import (
	"HomeWork7/gin/controller/user"
	"HomeWork7/gin/middleware"
	"HomeWork7/gin/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	uuids, _ := uuid.NewUUID()
	userid := uuids.String()

	//远程调用
	response, err := util.Rclient.Register(c, &user.UserInfo{
		UserId:   userid,
		Username: username,
		Password: password,
	})
	log.Println("error :  ", err)
	log.Println("response :  ", response)
	if err != nil {
		c.JSON(util.RequestError, gin.H{
			"status": util.RequestError,
			"err":    err,
		})
		return
	}
	if !response.GetResult() {
		c.JSON(util.RequestFailed, gin.H{
			"status":  util.RequestFailed,
			"message": "register failed,this user is exist ,please change a username.",
		})
		return
	}
	c.JSON(util.RequestSuccess, gin.H{
		"status":  util.RequestSuccess,
		"message": "register success",
	})
	return
}
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	info := &user.UserInfo{
		Username: username,
		Password: password,
	}
	response, err := util.Lcient.Login(c, info)
	log.Println("error :  ", err)
	log.Println("response :  ", response)

	if !response.GetResult() || err != nil {
		c.JSON(util.RequestFailed, gin.H{
			"status":  util.RequestFailed,
			"message": err.Error(),
		})
		return
	}

	token, err := middleware.GetToken(info)
	if err != nil {
		c.JSON(util.RequestError, gin.H{
			"status": util.RequestError,
			"err":    err,
		})
		return
	}
	c.JSON(util.RequestSuccess, gin.H{
		"status":  util.RequestSuccess,
		"message": "login success",
		"token":   token,
	})
	return
}

//.....
