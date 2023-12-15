package db

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model
	UserId   string
	Username string
	Password string
}

func (u *UserInfo) TableName() string {
	return "user_info"
}
func (u *UserInfo) Register(ctx context.Context) (bool, string) {

	//查重
	n := DB.Model(&UserInfo{}).Where("username = ?", u.Username).Find(&UserInfo{}).RowsAffected
	if n != 0 {
		return false, "the user already exist"
	}

	//创建用户
	n = DB.Model(&UserInfo{}).Create(&u).RowsAffected
	if n != 1 {
		return false, "create failed"
	}
	return true, ""
}
func (u *UserInfo) Login(ctx context.Context) (bool, error) {
	realUser := UserInfo{}

	//先查用户名是否存在，再查密码是否正确

	n := DB.Model(&UserInfo{}).Where("username = ?", u.Username).Find(&realUser).RowsAffected
	if n == 0 {
		return false, errors.New("用户名不存在")
	}
	if realUser.Password != u.Password {
		return false, errors.New("密码不正确")
	}
	return true, nil
}
