package dao

import "gorm.io/gorm"

type Userinfo struct {
	gorm.Model
	Username string `json:"username,omitempty" form:"username" binding:"required" gorm:"column:username;unique"`
	Password string `json:"password,omitempty" form:"password" binding:"required" gorm:"column:password"`
}
type Identify struct {
	gorm.Model
	Username string `json:"username,omitempty" form:"username" gorm:"column:username;unique"`
	Question string `json:"question,omitempty" form:"question" gorm:"column:question"`
	Answer   string `json:"answer,omitempty" form:"answer" gorm:"column:answer"`
	Password string `json:"password,omitempty" form:"password" gorm:"column:password"`
}
type Message struct {
	gorm.Model
	SrcUser  string `json:"src_user,omitempty" form:"srcUser"  gorm:"column:src_user"`
	DestUser string `json:"dst_user,omitempty" form:"destUser" binding:"required" gorm:"column:dest_user"`
	Message  string `json:"message,omitempty" form:"message" binding:"required" gorm:"column:message"`
}
