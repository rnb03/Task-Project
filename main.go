package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/rnb03/task-project/api"
)

func main() {
	// Create a new JSONPlaceholder client
	client := api.NewClient()

	// Create a new handler
	handler := api.NewHandler(client)

	// Create a new router
	router := mux.NewRouter()

	// API routes
	apiRouter := router.PathPrefix("/api").Subrouter()

	// Posts routes
	apiRouter.HandleFunc("/posts", handler.GetPosts).Methods("GET")
	apiRouter.HandleFunc("/posts/{id:[0-9]+}", handler.GetPost).Methods("GET")
	apiRouter.HandleFunc("/posts/{id:[0-9]+}/comments", handler.GetCommentsForPost).Methods("GET")

	// Comments routes
	apiRouter.HandleFunc("/comments", handler.GetComments).Methods("GET")

	// Users routes
	apiRouter.HandleFunc("/users", handler.GetUsers).Methods("GET")
	apiRouter.HandleFunc("/users/{id:[0-9]+}", handler.GetUser).Methods("GET")

	// Albums routes
	apiRouter.HandleFunc("/albums", handler.GetAlbums).Methods("GET")

	// Photos routes
	apiRouter.HandleFunc("/photos", handler.GetPhotos).Methods("GET")

	// Todos routes
	apiRouter.HandleFunc("/todos", handler.GetTodos).Methods("GET")

	// Add middleware for logging
	router.Use(loggingMiddleware)

	// Create a new server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start the server in a goroutine
	go func() {
		log.Println("Starting server on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on :8080: %v\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}

// loggingMiddleware logs all requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
