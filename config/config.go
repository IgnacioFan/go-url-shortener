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
	return &config, nil
}
