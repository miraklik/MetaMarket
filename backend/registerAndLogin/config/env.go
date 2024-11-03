package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	DbAddress  string
	DbUser     string
	DbPassword string
	DbName     string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost: GetEnv("PUBLIC_HOST", "http://localhost:3306"),
		Port:       GetEnv("PORT", "3306"),
		DbAddress:  fmt.Sprintf("%s:%s", GetEnv("DB_HOST", "localhost"), GetEnv("DB_PORT", "3306")),
		DbUser:     GetEnv("DB_USER", "root"),
		DbPassword: GetEnv("DB_PASSWORD", "55644"),
		DbName:     GetEnv("DB_NAME", "users"),
	}
}

func GetEnv(key, fallbacke string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallbacke
}
