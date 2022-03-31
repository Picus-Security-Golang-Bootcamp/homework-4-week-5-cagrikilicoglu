package main

import (
	"bookApp/internal/api/router"
	"bookApp/internal/domain/author"
	"bookApp/internal/domain/book"
	postgres "bookApp/pkg/db"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	// Set environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	// Initialize database
	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot be initalized.")
	}

	// Close db connection
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal("Database connection cannot be closed.")
	}
	defer sqlDb.Close()

	log.Println("Postgress connected")

	// Repositories
	router.BookRepo = book.NewBookRepository(db)
	router.AuthorRepo = author.NewAuthorRepository(db)

	// Setup databases
	router.BookRepo.SetupDatabase("./pkg/docs/books.csv")
	router.AuthorRepo.SetupDatabase("./pkg/docs/authors.csv")

}
func main() {

	// Create mux router
	r := mux.NewRouter()
	router.Handle(r)

	// Initialize server
	srv := &http.Server{
		Addr:         "127.0.0.1:8090",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	GracefulShutdown(srv, 15*time.Second)

}

func GracefulShutdown(srv *http.Server, timeout time.Duration) {
	c := make(chan os.Signal, 1)

	// when there is a interrupt signal, relay it to the channel
	signal.Notify(c, os.Interrupt)

	// block until any signal is received by the channel
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// wait until the timeout deadline and shutdown the server if there is no connections. if there is no connection shutdown immediately
	srv.Shutdown(ctx)

	log.Println("shutting down the server")
	os.Exit(0)
}
