package curriculo

import "os"

type Config struct {
	Port     string
	Env      string
	DataPath string
}

func NewConfig() *Config {
	return &Config{
		Port:     getEnv("PORT", "8080"),
		Env:      getEnv("APP_ENV", "development"),
		DataPath: getEnv("DATA_PATH", "internal/curriculo/data/resume.json"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
