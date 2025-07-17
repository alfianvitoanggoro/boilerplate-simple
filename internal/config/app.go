package config

type AppConfig struct {
	Name      string
	Env       string
	Host      string
	Port      string
	JWTSecret string
}

func LoadAppConfig() *AppConfig {
	return &AppConfig{
		Name:      getEnv("APP_NAME", "Boilerplate"),
		Env:       getEnv("ENV", "development"),
		Host:      getEnv("APP_HOST", "localhost"),
		Port:      getEnv("APP_PORT", "8080"),
		JWTSecret: getEnv("JWT_SECRET", "password"),
	}
}
