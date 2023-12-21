package controller

import (
	"HomeWork7/rpc/userSection/dao/cathe"
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
func (u *UserLogin) GetInfo(ctx context.Context, info *user2.UserInfo) (*user2.Response, error) {
	//通过用户名去查缓存
	information := cathe.Client.SMembers(ctx, info.Username).Val()
	if information == nil {
		//缓存中没有，再去数据库查
		realUser := db.UserInfo{}
		n := db.DB.Model(&db.UserInfo{}).Where("username = ?", info.Username).Find(&realUser).RowsAffected
		if n == 0 {
			return nil, errors.New("登陆信息错误")
		}
		//加入到缓存中
		cathe.Client.SAdd(ctx, realUser.Username, realUser.UserId, realUser.Username, realUser.Password)
		return &user2.Response{
			Result:   true,
			UserId:   realUser.UserId,
			Username: realUser.Username,
			Password: realUser.Password,
		}, nil
	}
	return &user2.Response{
		Result:   true,
		UserId:   information[0],
		Username: information[1],
		Password: information[2],
	}, nil
}
func (u *UserLogin) Star(ctx context.Context, info *user2.UserInfo) (*user2.Response, error) {
	//用户是否已经点过赞？先查缓存，没查到就查数据库
	isStared, err := cathe.Client.Get(ctx, info.Username+"starInfo").Result()
	if err != nil {
		//缓存中未查到，到数据去查
		realUser := db.UserInfo{}
		n := db.DB.Model(&db.UserInfo{}).First(&realUser).RowsAffected
		if n == 0 {
			return nil, errors.New("用户信息错误")
		}
		isStared = realUser.StarStatus

	}

	if isStared == "false" {
		n := cathe.Client.Incr(ctx, "default").Val()
		cathe.Client.Del(ctx, info.Username+"starInfo")
		cathe.Client.Set(ctx, info.Username+"starInfo", "true", 0)
		return &user2.Response{
			Result:    true,
			StarCount: n,
		}, nil
	}
	if isStared == "true" {
		n := cathe.Client.Decr(ctx, "default").Val()
		cathe.Client.Del(ctx, info.Username+"starInfo")
		cathe.Client.Set(ctx, info.Username+"starInfo", "false", 0)
		return &user2.Response{
			Result:    true,
			StarCount: n,
		}, nil
	}

	return &user2.Response{
		Result: false,
	}, errors.New("用户数据有误")
}
