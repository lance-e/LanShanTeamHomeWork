package cathe

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"log"
)

var Client = &redis.Client{}

func init() {
	viper.SetConfigFile("./config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		return
	}
	address := viper.GetString("redis.address")
	password := viper.GetString("redis.password")
	db := viper.GetInt("redis.db")
	log.Println(address)
	log.Println(password)
	log.Println(db)

	option := &redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	}
	//op, _ := redis.ParseURL("redis://root:@localhost:6379/0")

	Client = redis.NewClient(option)
	//
	Client.Set(context.Background(), "default", 0, 0)
}
