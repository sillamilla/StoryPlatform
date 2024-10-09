package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	GetConfig()
}

var (
	c *Config
)

type Config struct {
	Postgres Postgres
	HTTP     HTTP
}

type HTTP struct {
	Port string
}

type Postgres struct {
	URI      string
	User     string
	Password string
	DBName   string
	Port     string
}

func GetConfig() *Config {
	if c == nil {
		// Postgres
		uri := os.Getenv("DATABASE_URI")
		if uri == "" {
			log.Println("DATABASE_URI is not set")
		}

		user := os.Getenv("POSTGRES_USER")
		if user == "" {
			panic("POSTGRES_USER is not set")
		}

		password := os.Getenv("POSTGRES_PASSWORD")
		if password == "" {
			panic("POSTGRES_PASSWORD is not set")
		}

		dbName := os.Getenv("POSTGRES_DB")
		if dbName == "" {
			panic("POSTGRES_DB is not set")
		}

		port := os.Getenv("POSTGRES_PORT")
		if port == "" {
			panic("POSTGRES_PORT is not set")
		}

		// HTTP
		httpPort := os.Getenv("PORT")
		if httpPort == "" {
			panic("PORT is not set")
		}

		c = &Config{
			Postgres: Postgres{
				URI:      uri,
				User:     user,
				Password: password,
				DBName:   dbName,
				Port:     port,
			},
			HTTP: HTTP{
				Port: httpPort,
			},
		}
	}

	return c
}
