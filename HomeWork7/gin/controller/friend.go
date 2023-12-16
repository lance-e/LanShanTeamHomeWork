package controller

import (
	"HomeWork7/gin/util"
	"github.com/gin-gonic/gin"
)

func RequestAddFriend(c *gin.Context) {
	mode := util.AddFriend
	c.PostForm("srcId")
	c.PostForm("dstId")
}
func ReceiveAddFriend(c *gin.Context) {
	mode := util.DeleteFriend
}
func DeleteFriend(c *gin.Context) {
	mode := util.GroupFriend
}
func GroupFriend(c *gin.Context) {
	mode := util.ShowAllFriend
}
func ShowAllFriend(c *gin.Context) {
	mode := util.ShowGroupFriend
}
func ShowGroupFriend(c *gin.Context) {
	mode := util.SearchFriend
}
