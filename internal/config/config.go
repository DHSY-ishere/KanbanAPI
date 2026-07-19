package config

import "os"

type Config struct {
	Addr      string
	DBUrl     string
	JWTSecret string
}

func Load() Config {
	return Config{
		Addr:      getEnv("ADDR", ":8080"),
		DBUrl:     getEnv("DB_URL", "postgres://ecom_user:ecom_pass@localhost:5432/ecom?sslmode=disable"),
		JWTSecret: getEnv("JWT_SECRET", "dev-secret-change-me"),
	}
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}
