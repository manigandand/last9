package main

import (
	"fmt"
	"last9/api"
	"last9/config"
	appmiddleware "last9/middleware"
	"last9/store"
	"last9/task"
	"net/http"
	"os/signal"
	"syscall"

	"log"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

var (
	name    = "last9"
	version = "1.0.0"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Wrong length of arguments")
		return
	}

	config.Initialize(os.Args[1:]...)
	store.Init()
	api.InitAPI(name, version)
	task.Init()

	router := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "OPTIONS", "DELETE"},
		AllowedHeaders: []string{
			"Origin", "Authorization", "Access-Control-Allow-Origin",
			"Access-Control-Allow-Header", "Accept",
			"Content-Type", "X-CSRF-Token",
		},
		ExposedHeaders: []string{
			"Content-Length", "Access-Control-Allow-Origin", "Origin",
		},
		AllowCredentials: true,
		MaxAge:           300,
	})

	// cross & loger middleware
	router.Use(cors.Handler)
	router.Use(
		middleware.Logger,
		appmiddleware.Recoverer,
	)

	router.Route("/", api.Routes)

	interruptChan := make(chan os.Signal, 1)
	go func() {
		signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		// Block until we receive our signal.
		<-interruptChan

		log.Println("Shutting down db...")
		store.Store.Close()
		os.Exit(0)
	}()

	log.Println("Starting server on port:", config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.Port), router); err != nil {
		log.Fatal(err)
	}
}
