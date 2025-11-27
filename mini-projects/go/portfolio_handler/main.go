package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	ExactHealth    = "/healthcheck"
	ExactPortfolio = "/portfolio"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(ExactHealth, HealthCheckHandler())
	mux.HandleFunc(ExactPortfolio, PortfolioListHandler(PortfolioData))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
		// ReadTimeout:  15 * time.Second,
		// WriteTimeout: 15 * time.Second,
		// IdleTimeout:  60 * time.Second,
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("starting the HTTP server on port :8080 ...")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start the server: %v\n", err)
		}
	}()

	<-shutdown
	log.Println("shutting down HTTP server ...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("HTTP server shutdown error: %v\n", err)
	} else {
		log.Println("HTTP server gracefully stop")
	}
}
