package db

func Migrate() {
	err := DB.AutoMigrate(
		&UserInfo{},
	)
	if err != nil {
		panic("migrate database failed ...")
	}
}
