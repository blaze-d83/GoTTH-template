package main

import (
	"log"
	"net/http"

	"github.com/blaze-d83/go-GoTTH/store"
)

func main() {
	connectDB := store.InitDatabase()
	defer func ()  {
		sqlDB, err := connectDB.DB.DB()
		if err != nil {
			log.Println("Failed to connect SQLite DB for closing ", err)
			return 
		}
		if err = sqlDB.Close(); err != nil {
			log.Println("Failed to close the database connection", err)
		} else {
			log.Println("Database connection close successfully", err)
		}
	}()


	srv := http.Server{
		Addr: ":8090",
	}
	log.Println("Starting server on :8090")
	srv.ListenAndServe()

}
