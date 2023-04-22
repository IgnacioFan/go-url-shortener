package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Http     HttpConfig     `mapstructure:"http"`
	Postgres PostgresConfig `mapstructure:"postgres"`
	Redis    RedisConfig    `mapstructure:"redis"`
}

type HttpConfig struct {
	Port int `mapstructure:"port"`
}

type PostgresConfig struct {
	DSN string `mapstructure:"dsn"`
}

type RedisConfig struct {
	Host         string        `mapstructure:"host"`
	Port         uint          `mapstructure:"port"`
	DB           int           `mapstructure:"db"`
	Password     string        `mapstructure:"password"`
	MinIdleConns int           `mapstructure:"min_idle_conns"`
	MaxPoolSize  int           `mapstructure:"max_pool_size"`
	DialTimeout  time.Duration `mapstructure:"dial_timeout"`
}

func New() (*Config, error) {
	// set the name of the config file (without extension)
	viper.SetConfigName("default")

	// Set the search paths for the config file
	viper.AddConfigPath("$HOME/go-url-shortener/deployment/config")
	viper.AddConfigPath("./deployment/config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		return &config, err
	}
	return &config, nil
}
