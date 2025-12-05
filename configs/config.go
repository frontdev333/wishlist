package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Addr                string
	DataBaseCredentials string
}

func NewConfig() *Config {
	godotenv.Load()
	dbCredentials := fmt.Sprintf(
		"host=%s, port=%d, user=%s, password=%s, dbname=%s, sslmode=disable",
		"localhost", 5432, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"),
	)
	return &Config{
		Addr:                fmt.Sprintf(":%s", os.Getenv("Addr")),
		DataBaseCredentials: dbCredentials,
	}
}
