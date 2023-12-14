package db

import "HomeWork7/gin/dao/model"

func Migrate() {
	err := DB.AutoMigrate(
		&model.Friend{},
		&model.FriendApplication{},
		&model.UserInfo{},
	)
	if err != nil {
		panic("migrate database failed ...")
	}
}
