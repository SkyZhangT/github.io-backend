package config

import (
	"github.com/spf13/viper"
)


type Configuration struct {
	Environment string
	Token		string
	Img_dir		string
	Mongo       MongoConfiguration
	Limiter	LimiterConfiguration
}

type MongoConfiguration struct {
	Server     string
	Database   string
	Collection string
}

type LimiterConfiguration struct {
	Max int
    Methods []string
    TokenBucketTTL int
	Message string
    ContentType string
}

func GetConfig() Configuration {
	conf := Configuration{}

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	return conf
}
