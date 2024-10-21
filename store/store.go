package store

import (
	"fmt"
	"log"

	"github.com/blaze-d83/go-GoTTH/config"
	"github.com/blaze-d83/go-GoTTH/pkg/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Store struct {
	DB *gorm.DB
}

func InitDatabase() *Store  {
	db, err := NewSQLStorage()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	err = db.DB.AutoMigrate(types.Admin{})
	if err != nil {
		log.Fatalf("Failed to run migrate: %v", err)
	}
	return db

}

func NewSQLStorage(cfg config.Config) (*Store, error) {
	db, err := gorm.Open(sqlite.Open(cfg.Database.SQLite.DBPath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	return &Store{DB: db}, nil
}
