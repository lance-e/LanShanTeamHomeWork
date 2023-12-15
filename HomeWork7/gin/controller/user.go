package controller

import (
	"HomeWork7/gin/controller/user"
	"HomeWork7/gin/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	}, nil)
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
			"message": "register failed",
		})
		return
	}
	c.JSON(util.RequestSuccess, gin.H{
		"status":  util.RequestSuccess,
		"message": "register success",
	})
	//
}
func Login(c *gin.Context) {

	//
}

//.....
