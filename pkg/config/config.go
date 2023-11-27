package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

var once sync.Once

var config *Config

func InitConfig() *Config {
	if config == nil {
		once.Do(func() {
			err := setConfig()
			if err != nil {
				log.Fatal(err)
			}
		})
	} else {
		log.Fatal("Config already set.")
	}
	return config
}

func setConfig() error {
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return err
	}
	return nil
}
