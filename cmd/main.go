package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/lpernett/godotenv"
	"github.com/timur-danilchenko/project/internal/repository"
	"github.com/timur-danilchenko/project/internal/service"
	"github.com/timur-danilchenko/project/internal/transport"

	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can't load environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Can't export PORT value from environment")
	}

	dbuser := os.Getenv("DB_USER")
	if dbuser == "" {
		log.Fatal("No username for database connection")
	}
	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		log.Fatal("No database name passed")
	}

	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable", dbuser, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	conn, err := db.Conn(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	address := ":" + port
	router := http.NewServeMux()

	userRepository := &repository.UserRepository{Conn: conn}
	userService := &service.UserService{Repository: userRepository}
	userTransport := &transport.UserTransport{Service: userService}
	transport.SetUserTransport(router, userTransport)

	server := &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Server start on port:", port)
	server.ListenAndServe()
}
