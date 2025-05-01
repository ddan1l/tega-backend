package config

import (
	"os"
	"strconv"
	"sync"
)

type (
	Config struct {
		Server *Server
		Db     *Db
	}

	Server struct {
		Port int
	}

	Db struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SSLMode  string
		TimeZone string
	}
)

var (
	once           sync.Once
	configInstance *Config
)

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func GetConfig() *Config {
	once.Do(func() {
		configInstance = &Config{
			Server: &Server{
				Port: getEnvInt("BACKEND_PORT", 8080),
			},
			Db: &Db{
				Host:     getEnv("DB_HOST", "localhost"),
				Port:     getEnvInt("DB_PORT", 5432),
				User:     getEnv("DB_USER", ""),
				Password: getEnv("DB_PASSWORD", ""),
				DBName:   getEnv("DB_NAME", ""),
				SSLMode:  getEnv("DB_SSLMODE", "disable"),
				TimeZone: getEnv("DB_TIMEZONE", "UTC"),
			},
		}
	})

	return configInstance
}
