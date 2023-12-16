package db

func Migrate() {
	err := DB.AutoMigrate(
		&Relationship{},
	)
	if err != nil {
		panic("migrate database failed ...")
	}
}
