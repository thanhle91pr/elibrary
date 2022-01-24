package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	MongoUri string
}

func InitConfig() Config{
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf("read config file error:%s", err))
	}
	return Config{
		MongoUri: viper.GetString("MONGO_URI"),
	}
}
