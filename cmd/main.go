package main

import (
	"log"
	"net/http"

	"github.com/blaze-d83/go-GoTTH/internal/handlers"
	"github.com/blaze-d83/go-GoTTH/internal/middleware"
	"github.com/blaze-d83/go-GoTTH/pkg/config"
	"github.com/blaze-d83/go-GoTTH/pkg/logger"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env file")
	}

	dbConfig := config.LoadConfig()

	logger := logger.NewLogger()

	db, err := config.InitDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to initialize db: %v", err)
	}
	defer db.Close()

	handler := handlers.NewHandler(db, logger)

	fs := http.FileServer(http.Dir("./static"))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fs))
	mux.HandleFunc("/", handler.HomePage)
	mux.HandleFunc("/counter", handler.GetCounter)
	mux.HandleFunc("/increment", handler.IncrementCounter)
	mux.HandleFunc("/decrement", handler.DecrementCounter)

	loggedMux := middleware.LoggingMiddleware(logger, mux)

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
