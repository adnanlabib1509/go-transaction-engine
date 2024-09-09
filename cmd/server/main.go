package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adnanlabib1509/go-transaction-engine/internal/api"
	"github.com/adnanlabib1509/go-transaction-engine/internal/store"
	"github.com/adnanlabib1509/go-transaction-engine/pkg/logger"
)

func main() {
	// Initialize logger
	l := logger.New()

	// Initialize in-memory store
	s := store.NewMemoryStore()

	// Initialize API handlers
	handler := api.NewHandler(s, l)

	// Configure server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	// Start server
	go func() {
		l.Info("Starting server on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.Fatal("Could not listen on :8080: %v\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	l.Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		l.Fatal("Server forced to shutdown: %v\n", err)
	}

	l.Info("Server exiting")
}