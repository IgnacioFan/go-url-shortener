package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Postgres PostgresConfig `mapstructure:"postgres"`
}

type PostgresConfig struct {
	DSN string `mapstructure:"dsn"`
}

func New() (*Config, error) {
	viper.AddConfigPath("./deployment/config")
	viper.SetConfigName("default")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		return &config, err
	}
	fmt.Println(config)
	return &config, nil
}
