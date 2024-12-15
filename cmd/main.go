package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/timur-danilchenko/project/internal/app/config"
	"github.com/timur-danilchenko/project/internal/repository"
	"github.com/timur-danilchenko/project/internal/service"
	"github.com/timur-danilchenko/project/internal/transport"

	_ "github.com/lib/pq"
)

func main() {
	conf, err := config.Setup()

	if err != nil {
		log.Fatal("Can't load environment variables")
	}

	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable", conf.DB_USER, conf.DB_NAME)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx := context.Background()
	conn, err := db.Conn(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	address := ":" + conf.PORT
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

	log.Println("Server start on port:", conf.PORT)
	go func() {
		server.ListenAndServe()
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan
	server.Shutdown(ctx)
}
