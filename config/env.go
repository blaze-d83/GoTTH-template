package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName     string
	Environment string
	Server      ServerConfig
	Logger      LoggerConfig
	Database    DBConfig
}

type LoggerConfig struct {
	LogType   string
	LogFormat string
	LogLevel  string
	LogOutput string
}

type ServerConfig struct {
	Port string
	Mode string
}

type DBConfig struct {
	MySQL      MySQL
	PostgreSQL PostgreSQL
	SQLite     SQLite
}

type MySQL struct {
	User string
	Pass string
	Net  string
	Addr string
	Name string
}

type PostgreSQL struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

type SQLite struct {
	DBPath string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found,  using environment variables")
	}
	return &Config{
		AppName:     getEnv("APP", "my_app"),
		Environment: getEnv("ENVIRONMENT", "development"),
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
			Mode: getEnv("SERVER_MODE", "production"),
		},
		Logger: LoggerConfig{
			LogType:   getEnv("LOGGER_TYPE", "sync"),
			LogFormat: getEnv("LOG_FORMAT", "text"),
		},
		Database: DBConfig{
			MySQL:      MySQL{},
			PostgreSQL: PostgreSQL{},
			SQLite: SQLite{
				DBPath: getEnv("DB_PATH", "./db"),
			},
		},
	}

}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
