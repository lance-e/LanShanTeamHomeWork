package db

func Migrate() {
	err := DB.AutoMigrate(
		//&model.Friend{},
		//&model.FriendApplication{},
		&UserInfo{},
	)
	if err != nil {
		panic("migrate database failed ...")
	}
}
