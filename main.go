package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"todo-app/handler"
	"todo-app/repository/postgres"
	"todo-app/service"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	repository, err := postgres.NewTodoRepository()
	if err != nil {
		log.Fatal("Error: ", err)
	}
	service := service.NewTodoService(repository)
	handler := handler.NewTodoHandler(service)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", handler.GetTodos)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", handler.AddTodo)

	ch := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"}),
	)

	s := http.Server{
		Addr:         ":" + os.Getenv("PORT"),               // configure the bind address
		Handler:      ch(sm),                                // set the default handler
		ErrorLog:     log.New(os.Stderr, "", log.LstdFlags), // set the logger for the client
		ReadTimeout:  5 * time.Second,                       // max time to read request from the client
		WriteTimeout: 10 * time.Second,                      // max time to write response to the client
		IdleTimeout:  120 * time.Second,                     // max time for connections using TCP Keep-Alive
	}

	go func() {
		log.Println("Starting server...")

		err := s.ListenAndServe()
		if err != nil {
			log.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	sig := <-c
	log.Println("Got signal:", sig)
}
