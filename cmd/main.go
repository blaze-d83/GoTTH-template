package main

import (
	"log"
	"net/http"

	"github.com/blaze-d83/go-GoTTH/internal/handlers"
	"github.com/blaze-d83/go-GoTTH/pkg/config"
	"github.com/blaze-d83/go-GoTTH/pkg/logger"
)

func main() {
	cfg := config.LoadConfig()

	db, err := config.InitDB(cfg.DBConfig)
	if err != nil {
		log.Fatalf("Failed to initialize db: %v", err)
	}
	defer db.Close()

	logger, err := logger.InitializeLogger(cfg.LoggerConfig)
	if err != nil {
		log.Fatalf("Failed to intialize logger: %v", err)
	}

	handler := handlers.NewHandler(db, logger)

	http.HandleFunc("/", handler.HomePage)
	http.HandleFunc("/counter", handler.GetCounter)
	http.HandleFunc("/increment", handler.IncrementCounter)
	http.HandleFunc("/decrement", handler.DecrementCounter)

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
