package db

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strings"
)

// DB 全局对象
var DB *gorm.DB

func init() {
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
	user := viper.GetString("mysql2.user")
	pass := viper.GetString("mysql2.pass")
	ip := viper.GetString("mysql2.ip")
	port := viper.GetString("mysql2.port")
	dbname := viper.GetString("mysql2.dbname")
	log.Println(user)
	log.Println(pass)
	log.Println(ip)
	log.Println(port)
	log.Println(dbname)
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
