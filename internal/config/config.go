package config

import (
	"os"
	"strconv"
)

type Config struct {
	DB  *DBConfig
	App *AppConfig
}

func Load() *Config {
	return &Config{
		DB:  LoadDBConfig(),
		App: LoadAppConfig(),
	}
}

// getEnv returns the value of the environment variable or a default value if not set
func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

// getEnvAsInt returns the value of the environment variable as an integer or a default value if not set
func getEnvAsInt(key string, defaultVal int) int {
	if val := os.Getenv(key); val != "" {
		if intVal, err := strconv.Atoi(val); err == nil {
			return intVal
		}
	}
	return defaultVal
}
