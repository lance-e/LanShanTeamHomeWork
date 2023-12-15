package controller

import (
	"HomeWork7/rpc/userSection/dao/db"
	"HomeWork7/rpc/userSection/dao/model"
	user2 "HomeWork7/rpc/userSection/user"
	"context"
	"fmt"
)

type UserRegister struct {
	user2.UnimplementedRServerServer
}

func (u *UserRegister) Register(ctx context.Context, info *user2.UserInfo) (*user2.Response, error) {
	//注册业务
	userInfo := &model.UserInfo{

		UserId:   info.UserId,
		Username: info.Username,
		Password: info.Password,
	}
	//先进行查重
	n := db.DB.Find(&model.UserInfo{}).RowsAffected
	if n != 0 {
		return &user2.Response{
			Result: false,
		}, nil
	}

	//创建用户
	n = db.DB.Model(&model.UserInfo{}).Create(&userInfo).RowsAffected
	if n != 1 {
		return &user2.Response{
			Result: false,
		}, nil
	}

	fmt.Println("注册用户成功")
	return &user2.Response{
		Result: true,
	}, nil
}

type UserLogin struct {
	user2.UnimplementedLServerServer
}

func (u *UserLogin) Login(context.Context, *user2.UserInfo) (*user2.Response, error) {
	//登陆业务
	fmt.Println("登陆成功")
	return &user2.Response{
		Result: true,
	}, nil
}
