package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type RedisConfig struct {
	Host string
	Port string
}

type Config struct {
	Redis    RedisConfig
	Services Services
}

type AuthService struct {
	Port int `mapstructure:"port"`
	Host string
}

type Services struct {
	AuthService AuthService `mapstructure:"auth_service"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	config := &Config{}

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	config.Redis.Host = os.Getenv("REDIS_HOST")
	config.Redis.Port = os.Getenv("REDIS_PORT")

	config.Services.AuthService.Host = os.Getenv("AUTH_HOST")
	return config, nil
}
