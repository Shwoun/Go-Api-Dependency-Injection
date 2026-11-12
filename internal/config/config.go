package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	MongoURI string
	DBName   string
}

func NewConfig() *Config {

	err := godotenv.Load(".Env")
	if err != nil {
		return nil
	}
	return &Config{
		Port:     os.Getenv("Port"),
		MongoURI: os.Getenv("MongoDB_URI"),
		DBName:   os.Getenv("DBName"),
	}
}
