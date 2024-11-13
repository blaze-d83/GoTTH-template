package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type DBConfig struct {
	DBPath string
}

func LoadConfig() DBConfig {
	return DBConfig{
			DBPath: getEnv("DB_PATH", "./internal/repository/db/"),
		}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func InitDB(dbPath DBConfig) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath.DBPath)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	return db, nil
}
