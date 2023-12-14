package db

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strings"
)

var DB *gorm.DB

func Init() {
	//设置配置文件路径
	var err error
	viper.SetConfigFile("./config/config.yaml") //虽然说这一行就可以读取到配置文件了
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err = viper.ReadInConfig() //读取配置文件
	if err != nil {
		log.Fatalln(err)
	}
	user := viper.GetString("mysql.user")
	pass := viper.GetString("mysql.pass")
	ip := viper.GetString("mysql.ip")
	port := viper.GetString("mysql.port")
	dbname := viper.GetString("mysql.dbname")
	fmt.Println(user)
	dsn := strings.Join([]string{user, ":", pass, "@tcp(", ip, ":", port, ")/", dbname, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")

	log.Println(dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connect failed......")
	}
	log.Println("connect success!")
	Migrate() // 表迁移
	log.Println("migrate success")
}
