package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	LoggerConfig
	DBConfig
}

type LoggerConfig struct {
	LogFormat string
	LogType   string
	LogLevel  string
	LogOutput string
}

type DBConfig struct {
	DBPath string
}

func LoadConfig() Config {
	return Config{
		LoggerConfig: LoggerConfig{
			LogFormat: os.Getenv("LOG_FORMAT"),
			LogType:   os.Getenv("LOG_TYPE"),
			LogLevel:  os.Getenv("LOG_LEVEL"),
			LogOutput: os.Getenv("LOG_OUTPUT"),
		},
		DBConfig: DBConfig{
			DBPath: os.Getenv("DB_PATH"),
		},
	}
}

func InitDB(dbPath DBConfig) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath.DBPath)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	return db, nil
}

