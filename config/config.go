package config

import (
	"os"
)

type Config struct {
	Port          string
	MetricsSvcURL string
	BusesSvcURL   string
}

func GetConfig() *Config {
	return &Config{
		Port:          getEnv("PORT", ":3000"),
		MetricsSvcURL: getEnv("METRICS_SVC_URL", "localhost:50051"),
		BusesSvcURL:   getEnv("BUSES_SVC_URL", "localhost:50052"),
	}
}

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = defaultValue
	}
	return value
}
