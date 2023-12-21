package db

import (
	"HomeWork7/rpc/userSection/dao/cathe"
	"context"
	"errors"
	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model
	UserId     string
	Username   string
	Password   string
	StarStatus string
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
	//登陆成功，将用户数据存入缓存
	cathe.Client.SAdd(ctx, realUser.Username, realUser.UserId, realUser.Username, realUser.Password)
	//先看看缓存中有没有
	//用户点赞的数据
	//
	num, _ := cathe.Client.Get(ctx, realUser.Username+"starInfo").Int()

	if num == 0 {
		//缓存中没查到，就把数据库中的数据同步到缓存中
		//cathe.Client.Set(ctx, realUser.Username+"starInfo", "false", 0)
		if realUser.StarStatus == "" {
			cathe.Client.Set(ctx, realUser.Username+"starInfo", "false", 0)
		} else if realUser.StarStatus == "true" {
			cathe.Client.Set(ctx, realUser.Username+"starInfo", "true", 0)
		}
	}

	return true, nil
}
