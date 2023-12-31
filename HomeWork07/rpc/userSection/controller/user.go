package controller

import (
	"HomeWork7/rpc/userSection/dao/db"
	user2 "HomeWork7/rpc/userSection/user"
	"context"
	"errors"
	"fmt"
	"log"
)

type UserRegister struct {
	user2.UnimplementedRServerServer
}

func (u *UserRegister) Register(ctx context.Context, info *user2.UserInfo) (*user2.Response, error) {
	userInfo := &db.UserInfo{

		UserId:   info.UserId,
		Username: info.Username,
		Password: info.Password,
	}

	//调用dao层中的方法
	ok, msg := userInfo.Register(ctx)
	if !ok {
		if msg == "the user already exist" {
			return &user2.Response{
				Result: false,
			}, nil
		} else {
			return &user2.Response{
				Result: false,
			}, errors.New(msg)
		}
	}
	log.Println("注册用户成功")
	return &user2.Response{
		Result: true,
	}, nil
}

type UserLogin struct {
	user2.UnimplementedLServerServer
}

func (u *UserLogin) Login(ctx context.Context, info *user2.UserInfo) (*user2.Response, error) {
	userInfo := &db.UserInfo{
		UserId:   info.UserId,
		Username: info.Username,
		Password: info.Password,
	}
	ok, err := userInfo.Login(ctx)
	if !ok || err != nil {
		return &user2.Response{
			Result: false,
		}, err
	}

	fmt.Println("登陆成功")
	return &user2.Response{
		Result: true,
	}, nil
}
