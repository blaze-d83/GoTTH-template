package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/blaze-d83/go-GoTTH/store"
)

func main() {
	store := store.InitDatabase()
	defer func ()  {
		sqlDB, err := store.DB.DB()
		if err != nil {
			log.Println("Failed to connect SQLite DB for closing: ", err)
			return 
		}
		if err = sqlDB.Close(); err != nil {
			log.Println("Failed to close the database connection: ", err)
		} else {
			log.Println("Database connection close successfully")
		}
	}()

	srv := NewServer()

	httpServer := &http.Server{
		Addr: ":8090",
		Handler: srv,
	}

	go func()  {
		log.Printf("Listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error listening and serving: %s\n", err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup
	wg.Add(1)

	go func()  {
		defer wg.Done()
		<-sigs

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
		defer cancel()

		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			log.Printf("Error shutdown http server: %s\n", err)
		} else {
			log.Println("HTTP server shutdown gracefully")
		}
	}()
	wg.Wait()

}
