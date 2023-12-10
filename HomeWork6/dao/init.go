package dao

import (
	"RedRockHomework6/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strings"
)

var Db *gorm.DB

func init() {
	var databaseConfig config.Database
	err := config.Loading(&databaseConfig)
	if err != nil {
		log.Fatal(err)
	}
	dsn := strings.Join([]string{databaseConfig.Username, ":", databaseConfig.Password, "@tcp(", databaseConfig.Ip, ":", databaseConfig.Port, ")/", databaseConfig.Dbname, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	Db.AutoMigrate(&Userinfo{})
	Db.AutoMigrate(&Identify{})
	Db.AutoMigrate(&Message{})

}
