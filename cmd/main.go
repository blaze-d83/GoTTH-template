package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/blaze-d83/go-GoTTH/internal"
	"github.com/blaze-d83/go-GoTTH/pkg/config"
	"github.com/blaze-d83/go-GoTTH/pkg/logger"
	"github.com/joho/godotenv"
)

func main() {

	logger := logger.NewLogger()

	err := godotenv.Load()
	if err != nil {
		logger.Println("WARNING: failed to load .env file, using system environment variables")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatalf("FATAL: failed to load config : %v", err)
	}

	db, err := config.NewSQLiteConnection(cfg.DBConfig)
	if err != nil {
		logger.Fatalf("Failed to initialize database: %v", err)
	}

	handler := internal.NewHandler(db, logger)

	mux := http.NewServeMux()

	router := internal.RegisterRoutes(mux, handler, logger)

	srv := &http.Server{
		Addr:         cfg.Port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}
	
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)
	go func()  {
		logger.Printf("Starting server on %s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("ListenAndServer error: %v", err)
		}
	}()

	<-done
	logger.Println("Shutting down server..")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("failed to shutdown server: %v", err)
	}

	logger.Println("Server shutdown gracefully")
}
