package store

import (
	"fmt"
	"log"

	"github.com/blaze-d83/go-GoTTH/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Store struct {
	DB *gorm.DB
}

func InitDatabase() *Store  {
	dbPath := "./store/db/test.db"
	db, err := NewSQLStorage(dbPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	err = db.DB.AutoMigrate(types.Admin{})
	if err != nil {
		log.Fatalf("Failed to run migrate: %v", err)
	}
	return db

}

func NewSQLStorage(dbPath string) (*Store, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	return &Store{DB: db}, nil
}
