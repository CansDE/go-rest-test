package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Jwt      JwtConfig
	Database DBConfig
}

type JwtConfig struct {
	Secret string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func (d DBConfig) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", d.Host, d.User, d.Password, d.DBName, d.Port)
}

func Load() Config {
	_ = godotenv.Load()
	cfg := Config{
		Jwt: JwtConfig{
			Secret: getEnv("JWT_SECRET"),
		},
		Database: DBConfig{
			Host:     getEnv("DB_HOST"),
			DBName:   getEnv("DB_NAME"),
			Password: getEnv("DB_PASSWORD"),
			Port:     getEnv("DB_PORT"),
			User:     getEnv("DB_USER"),
		},
	}
	return cfg
}

func getEnv(value string) string {
	return os.Getenv(value)
}
