package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"

	"github.com/gorilla/mux"
	"github.com/jpedrodelacerda/torus2/pkg/handlers"
	"github.com/jpedrodelacerda/torus2/pkg/storage/nodb"
)

func main() {
	l := log.New(os.Stdout, "[Torus] ", log.LstdFlags)

	repo := nodb.NewStorage(l)
	uh := handlers.NewService("localhost", 8080, repo)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/users", uh.ListUsers)
	getRouter.HandleFunc("/users/{id:[0-9]+}", uh.FetchUser)
	getRouter.Use(uh.MiddlewareWriteJSON)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/users", uh.AddUser)
	postRouter.Use(uh.MiddlewareWriteJSON, uh.MiddlewareValidateUser)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/users/{id:[0-9]+}", uh.UpdateUser)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/users/{id:[0-9]+}", uh.DeleteUser)

	// Documentation bits
	// handler for documentation
	docsRouter := sm.Methods(http.MethodGet).Subrouter()
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml", Title: "Torus API Documentation"}
	sh := middleware.Redoc(opts, nil)

	docsRouter.Handle("/docs", sh)
	docsRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	// opts := middleware.RedocOpts{RedocURL: "/swagger.yaml"}
	// docsHandle := middleware.Redoc(opts, nil)

	// getRouter.Handle("/docs", docsHandle)
	// getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	server := http.Server{
		Addr:         uh.Addr(),
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		l.Println("Starting server at", server.Addr)

		err := server.ListenAndServe()
		if err != nil {
			l.Printf("Failed to start server")
			os.Exit(1)
		}
	}()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	signal.Notify(sigchan, os.Kill)

	sig := <-sigchan
	l.Println("Got signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}
