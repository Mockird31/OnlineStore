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
	Postgres PostgresConfig
	Port int `mapstructure:"port"`
}

type AuthService struct {
	Port int `mapstructure:"port"`
	Host string
}

type PostgresConfig struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	MaxOpenConns     int `mapstructure:"max_open_conns"`
	MaxIdleConns     int `mapstructure:"max_idle_conns"`
	MaxLifetime      int `mapstructure:"max_lifetime"`
}

type UserService struct {
	Port int `mapstructure:"port"`
	Host string
}

type Services struct {
	AuthService AuthService `mapstructure:"auth_service"`
	UserService UserService `mapstructure:"user_service"`
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

	config.Postgres.PostgresHost = os.Getenv("POSTGRES_HOST")
	config.Postgres.PostgresPort = os.Getenv("POSTGRES_PORT")
	config.Postgres.PostgresUser = os.Getenv("POSTGRES_USER")
	config.Postgres.PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	config.Postgres.PostgresDB = os.Getenv("POSTGRES_DB")

	config.Services.AuthService.Host = os.Getenv("AUTH_HOST")
	config.Services.UserService.Host = os.Getenv("USER_HOST")
	return config, nil
}
