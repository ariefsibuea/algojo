package main

import (
	"net/http"
	"time"
)

type HealthCheckResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

func HealthCheckHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			response := HealthCheckResponse{
				Status:    "available",
				Timestamp: time.Now(),
			}

			WriteJSON(w, http.StatusOK, response)
			return
		}

		WriteJSON(w, http.StatusMethodNotAllowed, ErrorResponse{Error: "method not allowed"})
	}
}
