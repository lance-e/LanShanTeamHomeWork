package model

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	UserId   string
	Username string
	Password string
}

func (u *UserInfo) TableName() string {
	return "user_info"
}
